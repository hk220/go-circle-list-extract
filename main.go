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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/unicode/norm"
)

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

type Event struct {
	CircleListURL string `mapstructure:"circle_list_url"`
	Parser        string `mapstructure:"parser"`
}

func rootRun(cmd *cobra.Command, args []string) {
	validate()

	client := &http.Client{}

	// Parse events
	var events map[string]Event
	if err := viper.UnmarshalKey("event", &events); err != nil {
		log.Fatal(err)
	}

	// event key check
	event, ok := events[args[0]]
	if !ok {
		log.Fatalf("no event key: %s", args[0])
	}

	// Request the HTML Page
	resp, err := client.Get(event.CircleListURL)
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

	if viper.GetString("format") == "csv" {
		// Write CSV File
		csvFile, err := appFs.OpenFile(viper.GetString("output"), os.O_RDWR|os.O_CREATE, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
		defer csvFile.Close()

		csvPrinter := NewPrinter(cl, csvFile)
		err = csvPrinter.CSV()
		if err != nil {
			log.Fatal(err)
		}
	}

	if viper.GetString("format") == "json" {
		// Write JSON File
		jsonFile, err := appFs.OpenFile(viper.GetString("output"), os.O_RDWR|os.O_CREATE, os.ModePerm)
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
}

var (
	rootCmd = &cobra.Command{
		Use:   "go-circle-list-extract",
		Short: "go-circle-list-extract extracts a circle list of Comitia in JSON or CSV format.",
		Args:  cobra.ExactArgs(1),
		Run:   rootRun,
	}

	cfgFile   string
	eventName string
)

func validate() {
	if viper.GetString("format") != "json" && viper.GetString("format") != "csv" {
		log.Fatal("invalid format")
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&eventName, "event", "e", "comitia134", "event name")
	rootCmd.PersistentFlags().StringP("format", "f", "csv", "output format")
	rootCmd.PersistentFlags().StringP("output", "o", "circles.csv", "output file name")
	viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
