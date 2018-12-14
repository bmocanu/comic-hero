package main

import (
	"comic-hero/src/retrieve"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.InfoLevel)

	var issue, _ = retrieve.SinfestIssue()
	fmt.Println(issue)

	issue, _ = retrieve.DilbertIssue()
	fmt.Println(issue)
}
