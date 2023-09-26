package html

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

// RemoveHtmlTag 去除代码中的html标签
func RemoveHtmlTag(rawHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		return "", err
	}
	htmlString := doc.Text()
	return htmlString, nil
}
