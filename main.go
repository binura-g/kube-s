package main

import (
	"github.com/binura-g/kube-s/pkg/search"
	"log"
	"os"
)

func main() {
	log.SetFlags(0)
	if len(os.Args) != 3 {
		log.Println("Usage:\n\tkube-s <ResourceKind> <Pattern>\n\tEg. kube-s pods my-app")
		os.Exit(1)
	}

	kind := os.Args[1]
	pattern := os.Args[2]
	resultsChannel := make(chan search.Result)
	go search.Run(kind, pattern, resultsChannel)

	for r := range resultsChannel {
		log.Printf("%s\t%s\t\t%s\n", r.Cluster, r.Ns, r.Name)
	}

	os.Exit(0)
}
