package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mammenj/fulltextsearch/indexer"
	"github.com/mammenj/fulltextsearch/loader"
)

func main() {
	var datapath, search_query string
	flag.StringVar(&datapath, "p", "enwiki-latest-abstract1.xml.gz", "wiki abstract dump path")
	//flag.StringVar(&query, "q", "Small wild cat", "search query")
	flag.Parse()

	log.Println("Starting full text search in memory...")

	start := time.Now()
	docs, err := loader.LoadDocument(datapath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("1> Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idDoc := indexer.NewIndex(docs)
	//idDoc.Index(docs)
	log.Printf("2> Indexed %d documents in %v", len(docs), time.Since(start))

	// get search query
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(">> Enter search query:: ")
		scanner.Scan()
		search_query = scanner.Text()
		if len(search_query) != 0 {
			start = time.Now()
			matchedIDs := idDoc.Search(search_query)
			log.Printf("3> Found %d documents in %v", len(matchedIDs), time.Since(start))

			for _, id := range matchedIDs {
				doc := docs[id]
				log.Printf("docID:: %d\t Text::%s\n", id, doc.Text)
			}
			log.Printf("------------------------*****************************************************-------------------------------------")
		}
	}

}
