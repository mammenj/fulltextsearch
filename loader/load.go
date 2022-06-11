package loader

import (
	"compress/gzip"
	"encoding/xml"
	"os"
)

// struct representing wikipedia abstract document.
type Document struct {
	Title string `xml:"title" json:"title"`
	URL   string `xml:"url" json:"url"`
	Text  string `xml:"abstract" json:"abstract"`
	ID    int
}

// data example: https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract1.xml.gz
func LoadDocument(path string) ([]Document, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, err
	}
	defer gz.Close()
	dec := xml.NewDecoder(gz)
	dump := struct {
		Documents []Document `xml:"doc"`
	}{}
	if err := dec.Decode(&dump); err != nil {
		return nil, err
	}
	docs := dump.Documents
	for i := range docs {
		docs[i].ID = i
	}
	return docs, nil
}
