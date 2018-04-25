package main

import (
	"net/url"
	"time"

	fmt "github.com/cathalgarvey/fmtless"

	"bytes"

	"math/rand"

	"github.com/cathalgarvey/fmtless/encoding/json"
	"github.com/gopherjs/gopherjs/js"
	"github.com/gopherjs/jquery"
	"honnef.co/go/js/dom"
	"honnef.co/go/js/xhr"
)

var (
	jQuery   = jquery.NewJQuery
	window   = dom.GetWindow()
	document = window.Document()
)

const (
	engineName = "Flowfly"
)

// Replace main functionality with start so that start is only run when page fully loads
func main() {
	titleAnimation()
	rand.Seed(time.Now().UnixNano())
	js.Global.Set("svue", map[string]interface{}{
		"start":  start,
		"search": search,
	})
}

func start() {
}

func titleAnimation() {
	go func() {
		var colors = []string{"#171738", "#2E1760", "#3423A6", "#7180B9", "#171738", "#2E1760", "#3423A6"}
		for i := 0; true; i++ {
			shuffle(colors)
			jQuery(".logText").Each(func(index int, element interface{}) {
				if i+index > len(colors) {
					i = 0
					if index > len(colors) {
						println("Not enough colors specified!")
					}
				}
				jQuery(element).SetCss("color", colors[index+i-1])
			})
			time.Sleep(time.Millisecond * 1000)
		}
	}()
}

//http://localhost:8090/yacysearch.json?query=foaf&maximumRecords=10
func search(query string) {
	go func() {
		res, err := xhr.Send("GET", fmt.Sprintf("http://skyrisbactera.com:8090/yacysearch.json?query=%s&maximumRecords=10&resource=global&verify=false", url.QueryEscape(query)), []byte(""))
		if err != nil {
			fmt.Println(err)
		}
		var respMar yacy
		reader := bytes.NewReader(res)
		json.NewDecoder(reader).Decode(&respMar)
		for _, v := range respMar.Channels {
			println(v.Title)
		}
	}()
}

type yacy struct {
	Channels []struct {
		Description string `json:"description"`
		Image       struct {
			Link  string `json:"link"`
			Title string `json:"title"`
			URL   string `json:"url"`
		} `json:"image"`
		Items []struct {
			Code        string `json:"code"`
			Description string `json:"description"`
			FaviconURL  string `json:"faviconUrl"`
			File        string `json:"file"`
			GUID        string `json:"guid"`
			Host        string `json:"host"`
			Link        string `json:"link"`
			Path        string `json:"path"`
			PubDate     string `json:"pubDate"`
			Ranking     string `json:"ranking"`
			Size        string `json:"size"`
			Sizename    string `json:"sizename"`
			Title       string `json:"title"`
			Urlhash     string `json:"urlhash"`
		} `json:"items"`
		ItemsPerPage string `json:"itemsPerPage"`
		Link         string `json:"link"`
		Navigation   []struct {
			Displayname string `json:"displayname"`
			Elements    []struct {
				Count    string `json:"count"`
				Modifier string `json:"modifier"`
				Name     string `json:"name"`
				URL      string `json:"url"`
			} `json:"elements"`
			Facetname string `json:"facetname"`
			Max       string `json:"max"`
			Mean      string `json:"mean"`
			Min       string `json:"min"`
			Type      string `json:"type"`
		} `json:"navigation"`
		SearchTerms  string `json:"searchTerms"`
		StartIndex   string `json:"startIndex"`
		Title        string `json:"title"`
		TotalResults string `json:"totalResults"`
	} `json:"channels"`
}

func shuffle(vals []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) > 0 {
		n := len(vals)
		randIndex := r.Intn(n)
		vals[n-1], vals[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}
