package main

import (
	"github.com/jazzray/sparql"
	"time"
	"log"
)

func main() {

	repo, err := sparql.NewRepo("http://localhost:7200/repositories/test",
		sparql.Timeout(time.Millisecond*1500),
	)
	if err != nil {
		log.Fatal(err)
	}

	res, err := repo.Query("SELECT * WHERE { ?s ?p ?o } LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}

	res.Solutions()

}