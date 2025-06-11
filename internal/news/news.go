package news

import (
	"encoding/xml"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type RSS struct {
	Channel struct {
		Items []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

const feedURL = "https://habr.com/ru/rss/all/all/"

// Fetch retrieves news items from Habr RSS feed
func Fetch() ([]Item, error) {
	resp, err := http.Get(feedURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rss RSS
	if err := xml.Unmarshal(data, &rss); err != nil {
		return nil, err
	}
	return rss.Channel.Items, nil
}

// CleanText removes unwanted characters from text
func CleanText(s string) string {
	// Replace Russian ёЁ with еЕ
	s = strings.ReplaceAll(s, "ё", "е")
	s = strings.ReplaceAll(s, "Ё", "Е")

	re := regexp.MustCompile("[^a-zA-Z0-9а-яА-Я`!@#$%^&*()\\-_=+\\[\\]{}:;\"'|\\\\<>,.?/ ]+")
	return re.ReplaceAllString(s, "")
}

// FetchArticleText downloads the article page and extracts its plain text
func FetchArticleText(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	text := doc.Find(".article-formatted-body").Text()
	if text == "" {
		text = doc.Find(".tm-article-body").Text()
	}
	text = strings.TrimSpace(text)
	return CleanText(text), nil
}
