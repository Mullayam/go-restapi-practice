package utils

import (
	"compress/gzip"
	"encoding/json"
	"log"
	"os"
)

type Document struct {
	ID    int    `json:"id"`
	Text  string `json:"text" yaml:"text" toml:"text" bson:"text" xml:"text"`
	URL   string `json:"url" yaml:"url" toml:"url" bson:"url" xml:"url"`
	Title string `json:"title" yaml:"title" toml:"title" bson:"title" xml:"title"`
}

func LoadDocuments(path string) ([]Document, error) {

	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		log.Fatalf("Error creating gzip reader: %v", err)
	}
	dec := json.NewDecoder(gz)
	dunp := struct {
		Documents []Document `json:"docs"`
	}{}
	if err := dec.Decode(&dunp); err != nil {
		log.Fatalf("Error decoding json: %v", err)
	}
	docs := dunp.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return dunp.Documents, nil

}
