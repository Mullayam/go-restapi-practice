package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/Mullayam/text-search-engine/utils"
)

func main() {
	var indexPath, dataPath string
	flag.StringVar(&indexPath, "index", "index.json", "Path to the index file")
	flag.StringVar(&dataPath, "data", "data.json", "Path to the data file")
	flag.Parse()

	docs, err := utils.LoadDocuments(indexPath)
	if err != nil {
		log.Fatalf("Error loading index: %v", err)
	}
	log.Printf("Index loaded with %d documents", len(docs))
	start := time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Index built in %v", time.Since(start))
	start = time.Now()
	matchIDs := idx.Search("The quick brown fox")
	log.Printf("Search completed in %v", time.Since(start))
	log.Printf("Found %d matches", len(matchIDs))

	for _, doc := range matchIDs {
		log.Printf("Matched document: %v", docs[doc].Text)
	}

}
