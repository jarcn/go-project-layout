package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"testing"
	"time"

	es "github.com/elastic/go-elasticsearch/v7"
	"github.com/stretchr/testify/assert"
)

var (
	gclient *es.Client
)

func init() {
	var err error
	gclient, err = es.NewClient(es.Config{
		Addresses: []string{
			"http://192.168.50.117:9201",
			"http://192.168.50.117:9202",
			"http://192.168.50.117:9203",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

func TestNewESClient(t *testing.T) {
	fmt.Println(gclient.Info())
}

func TestCreateIndex(t *testing.T) {
	a := assert.New(t)
	response, err := gclient.Indices.Create("book_002", gclient.Indices.Create.WithBody(strings.NewReader(`
	{
		"aliases": {
			"book":{}
		},
		"settings": {
			"analysis": {
				"normalizer": {
					"lowercase": {
						"type": "custom",
						"char_filter": [],
						"filter": ["lowercase"]
					}
				}
			},
			"number_of_shards": 3,
			"number_of_replicas": 1,
			"refresh_interval": "10s"
		},
		"mappings": {
			"properties": {
				"name": {
					"type": "keyword",
					"normalizer": "lowercase"
				},
				"price": {
					"type": "double"
				},
				"summary": {
					"type": "text"
				},
				"author": {
					"type": "keyword"
				},
				"pubDate": {
					"type": "date"
				},
				"pages": {
					"type": "integer"
				}
			}
		}
	}
	`)))
	a.Nil(err)
	fmt.Println(response)
}

func TestGetIndex(t *testing.T) {
	a := assert.New(t)
	response, err := gclient.Indices.Get([]string{"book"})
	a.Nil(err)
	fmt.Println(response)
}

func TestDeleteIndex(t *testing.T) {
	a := assert.New(t)
	response, err := gclient.Indices.Delete([]string{"book_002"})
	a.Nil(err)
	fmt.Println(response)
}

type doc struct {
	Doc interface{} `json:"doc"`
}

type Book struct {
	ID      string     `json:"id,omitempty"`
	Author  string     `json:"author,omitempty"`
	Name    string     `json:"name,omitempty"`
	Pages   int        `json:"pages,omitempty"`
	Price   float64    `json:"price,omitempty"`
	PubDate *time.Time `json:"pubDate,omitempty"`
	Summary string     `json:"summary,omitempty"`
}

func TestCreateDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	pubDate := time.Now()
	err := json.NewEncoder(body).Encode(&Book{
		Author:  "金庸",
		Price:   96.0,
		Name:    "天龙八部",
		Pages:   1978,
		PubDate: &pubDate,
		Summary: "...",
	})
	a.Nil(err)
	response, err := gclient.Create("book", "10001", body)
	a.Nil(err)
	fmt.Println(response)
}

func TestIndexDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	pubDate := time.Now()
	err := json.NewEncoder(body).Encode(&Book{
		Author:  "金庸",
		Price:   96.0,
		Name:    "天龙八部",
		Pages:   1978,
		PubDate: &pubDate,
		Summary: "...",
	})
	a.Nil(err)
	response, err := gclient.Index("book", body, gclient.Index.WithDocumentID("10001"))
	a.Nil(err)
	t.Log(response)
}

func TestPartialUpdateDocument(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(&doc{
		Doc: &Book{
			Name: "天龙八部",
		},
	})
	a.Nil(err)
	response, err := gclient.Update("book", "10001", body)
	a.Nil(err)
	t.Log(response)
}

func TestGetDocument(t *testing.T) {
	a := assert.New(t)
	response, err := gclient.Get("book", "10001")
	a.Nil(err)
	t.Log(response)
}

func TestBulk(t *testing.T) {
	createBooks := []*Book{
		{
			ID:     "10001",
			Name:   "笑傲江湖",
			Author: "金庸",
		},
	}
	deleteBookIds := []string{}

	a := assert.New(t)
	body := &bytes.Buffer{}
	for _, book := range createBooks {
		meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, book.ID, "\n"))
		data, err := json.Marshal(book)
		a.Nil(err)
		data = append(data, "\n"...)
		body.Grow(len(meta) + len(data))
		body.Write(meta)
		body.Write(data)
	}
	for _, id := range deleteBookIds {
		meta := []byte(fmt.Sprintf(`{ "delete" : { "_id" : "%s" } }%s`, id, "\n"))
		body.Grow(len(meta))
		body.Write(meta)
	}
	t.Log(body.String())

	response, err := gclient.Bulk(body, gclient.Bulk.WithIndex("book"))
	a.Nil(err)
	t.Log(response)
}

func TestSearch(t *testing.T) {
	a := assert.New(t)
	body := &bytes.Buffer{}
	body.WriteString(`
	{
		"_source":{
		  "excludes": ["name"]
		}, 
		"query": {
		  "match_phrase": {
			"name": "神雕侠侣"
		  }
		},
		"sort": [
		  {
			"pages": {
			  "order": "desc"
			}
		  }
		], 
		"from": 0,
		"size": 5
	}
	`)
	response, err := gclient.Search(gclient.Search.WithIndex("book"), gclient.Search.WithBody(body))
	a.Nil(err)
	t.Log(response)
}
