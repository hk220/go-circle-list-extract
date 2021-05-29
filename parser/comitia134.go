package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/hk220/go-circle-list-extract/circle"
)

func Comitia134Parser(doc *goquery.Document) (*circle.CircleList, error) {
	var cl circle.CircleList

	// Find items
	doc.Find("table[border=\"0\"][cellpadding=\"0\"][cellspacing=\"0\"][style=\"width:100%;\"]").Find("tr").Each(func(index int, selection *goquery.Selection) {
		tags := selection.Find("td")
		spaceElem := tags.First()
		nameElem := tags.First().Next()

		// Skip index row
		if _, isIndex := spaceElem.Attr("colspan"); isIndex {
			return
		}

		space := normalize(spaceElem.Text())
		name := normalize(nameElem.Text())
		// Skip null row
		if space == "" {
			return
		}

		if url, urlExists := nameElem.Find("a").Attr("href"); urlExists {
			trimedUrl := trimSpace(url)
			cl.Add(&circle.Circle{Space: space, Name: name, URL: trimedUrl})
		} else {
			cl.Add(&circle.Circle{Space: space, Name: name, URL: ""})
		}
	})

	return &cl, nil
}
