package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocarina/gocsv"
	"github.com/spf13/afero"
	"golang.org/x/text/unicode/norm"
)

const comitiaCircleListURL = "https://www.comitia.co.jp/history/134list_sp.html"

type Circle struct {
	Space  string `json:"space" csv:"space"`
	Circle string `json:"circle" csv:"circle"`
	URL    string `json:"url" csv:"url"`
}

type CircleList struct {
	circles []Circle
}

func (cl *CircleList) Add(c *Circle) {
	cl.circles = append(cl.circles, *c)
}

func (cl *CircleList) String() [][]string {
	var result [][]string
	for _, c := range cl.circles {
		result = append(result, []string{c.Space, c.Circle, c.URL})
	}
	return result
}

type Parser struct {
	doc        *goquery.Document
	CircleList *CircleList
}

func NewParser(doc *goquery.Document) *Parser {
	return &Parser{doc: doc, CircleList: new(CircleList)}
}

func (p *Parser) Parse() error {
	// Find items
	p.doc.Find("table[border=\"0\"][cellpadding=\"0\"][cellspacing=\"0\"][style=\"width:100%;\"]").Find("tr").Each(func(index int, selection *goquery.Selection) {
		tags := selection.Find("td")
		spaceElem := tags.First()
		nameElem := tags.First().Next()

		// Skip index row
		if _, isIndex := spaceElem.Attr("colspan"); isIndex {
			return
		}

		space := p.normalize(spaceElem.Text())
		name := p.normalize(nameElem.Text())
		// Skip null row
		if space == "" {
			return
		}

		if url, urlExists := nameElem.Find("a").Attr("href"); urlExists {
			trimedUrl := p.trimSpace(url)
			p.CircleList.Add(&Circle{space, name, trimedUrl})
		} else {
			p.CircleList.Add(&Circle{space, name, ""})
		}
	})

	return nil
}

func (p *Parser) trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func (p *Parser) normalize(s string) string {
	return strings.TrimSpace(norm.NFKC.String(s))
}

type Printer struct {
	cl     *CircleList
	writer io.Writer
}

func NewPrinter(cl *CircleList, writer io.Writer) *Printer {
	return &Printer{cl: cl, writer: writer}
}

func (p *Printer) CSV() error {
	err := gocsv.Marshal(p.cl.circles, p.writer)
	if err != nil {
		return err
	}
	return nil
}

func (p *Printer) JSON() error {
	s, err := json.MarshalIndent(p.cl.circles, "", "\t")
	if err != nil {
		return err
	}

	_, err = p.writer.Write(s)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	client := &http.Client{}

	// Request the HTML Page
	resp, err := client.Get(comitiaCircleListURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	parser := NewParser(doc)
	err = parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	cl := parser.CircleList

	var appFs = afero.NewOsFs()

	// Write CSV File
	csvFile, err := appFs.OpenFile("test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvPrinter := NewPrinter(cl, csvFile)
	err = csvPrinter.CSV()
	if err != nil {
		log.Fatal(err)
	}

	// Write JSON File
	jsonFile, err := appFs.OpenFile("test.json", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonPrinter := NewPrinter(cl, jsonFile)
	err = jsonPrinter.JSON()
	if err != nil {
		log.Fatal(err)
	}
}
