package main

import (
	"reflect"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

var (
	parserReader = strings.NewReader(`
<main class="clearfix">
	<h2 id="page_title">COMITIA115参加サークルリスト</h2>
	<table border="0" cellpadding="0" cellspacing="0" style="width:100%;">
		<tbody>
		<tr><td colspan="2"><h3 class="heading_1"><a name="a" class="jump_link"></a>Ａ</h3></td></tr>
		<tr><td width="45"></td><td></td></tr>
		<tr><td width="45">Ａ01a</td><td><a href="http://www.comitia.co.jp/" target="_blank">COMITIA</a></td></tr>
		<tr><td colspan="2"><h3 class="heading_1"><a name="aa" class="jump_link"></a>あ</h3></td></tr>
		<tr><td width="45">あ01a</td><td>comitia</td></tr>
		<tr><td colspan="2"><h3 class="heading_1"><a name="ten" class="jump_link"></a>展示</h3></td></tr>
		<tr><td width="45">展01</td><td><a href="http://www.comitia.co.jp/html/about.html" target="_blank">コミティア</a></td></tr>
		</tbody>
	</table>
</main>
	`)

	parserExpected = []Circle{
		{"A01a", "COMITIA", "http://www.comitia.co.jp/"},
		{"あ01a", "comitia", ""},
		{"展01", "コミティア", "http://www.comitia.co.jp/html/about.html"},
	}
)

func TestParser(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(parserReader)
	if err != nil {
		t.Fatal(err)
	}

	parser := NewParser(doc)
	parser.Parse()
	actual := parser.CircleList.circles

	if !reflect.DeepEqual(parserExpected, actual) {
		t.Errorf("Not match circles: %+v", actual)
	}
}
