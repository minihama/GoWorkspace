package main

import (
	"bytes"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-resty/resty/v2"
	"golang.org/x/net/html/charset"
	"gopkg.in/go-playground/pool.v3"
	"io"
	"net/http"
	"strings"
)

type HtmlResult struct {
	Url     string
	Title   string
	Charset string
}

var urls = []string{
	"https://www.naver.com/",
	"https://www.daum.net/",
	"https://www.google.co.kr/",
	"https://www.youtube.com/",
	"https://stackoverflow.com/",
	"https://regex101.com/",
	"https://httpbin.org/",
	"https://casbin.org/ko/",
	"https://jsonplaceholder.typicode.com/",
}

func main() {
	p := pool.NewLimited(4)
	defer p.Close()
	batch := p.Batch()
	go func() {
		for i := 0; i < len(urls); i++ {
			batch.Queue(fetchHtmlResult(urls[i]))
		}
		batch.QueueComplete()
	}()
	for fetechJob := range batch.Results() {
		if err := fetechJob.Error(); err != nil {
			panic(err)
		}
		if res, ok := fetechJob.Value().(HtmlResult); ok {
			spew.Dump(res)
		}
	}
}

func fetchHtmlResult(addr string) pool.WorkFunc {
	return func(wu pool.WorkUnit) (interface{}, error) {
		if wu.IsCancelled() {
			return nil, nil
		}
		client := resty.New().SetDoNotParseResponse(true)
		resp, err := client.R().Get(addr)
		if err != nil {
			return nil, err
		}
		defer func() {
			_ = resp.RawResponse.Body.Close()
		}()
		if resp.StatusCode() != http.StatusOK {
			return nil, errors.New("http error: " + resp.Status())
		}
		var buf bytes.Buffer
		tee := io.TeeReader(resp.RawResponse.Body, &buf)
		_, charsetName, ok := charset.DetermineEncoding(buf.Bytes(), "")
		if !ok {
			charsetName = "utf-8"
		}
		doc, err := goquery.NewDocumentFromReader(tee)
		if err != nil {
			return nil, err
		}
		title := doc.Find("title").Text()
		return HtmlResult{
			Url:     addr,
			Title:   title,
			Charset: strings.ToLower(charsetName),
		}, nil
	}
}
