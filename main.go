package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func main() {
	log.SetFlags(0)

	if len(os.Args) != 3 {
		fmt.Println("Usage:\n\tkube-s <ResourceKind> <Pattern>\n\tEg. kube-s pods my-app")
		os.Exit(1)
	}

	kind := os.Args[1]
	pattern := os.Args[2]

	resultsCh := make(chan string)
	go search(kind, pattern, resultsCh)
	for result := range resultsCh {
		log.Println(result)
	}
	os.Exit(0)
}

func search(kind, pattern string, resultCh chan string) {
	clusters, err := listClusters()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Searching for all [%s] Resources with Names matching %q in %d cluster(s)...\n", kind, pattern, len(clusters))

	kubectlOutputCh := make(chan []byte, len(clusters))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(clusters))

		for _, cluster := range clusters {
			go func(cluster string) {
				defer wg.Done()
				kubectlArgs := []string{"get", kind, "--all-namespaces", "--no-headers", fmt.Sprintf("--context=%s", cluster)}
				output, err := exec.Command("kubectl", kubectlArgs...).CombinedOutput()
				if err != nil {
					log.Printf("Error: %q %s\n", cluster, output)
				}
				kubectlOutputCh <- output
			}(cluster)
		}
		wg.Wait()
		close(kubectlOutputCh)
	}()

	var resultWg sync.WaitGroup
	for output := range kubectlOutputCh {
		resultWg.Add(1)
		go func(in []byte) {
			defer resultWg.Done()
			scanner := bufio.NewScanner(bytes.NewReader(in))
			for scanner.Scan() {
				s := scanner.Text()
				if strings.Contains(s, pattern) {
					resultCh <- s
				}
			}
		}(output)
	}

	resultWg.Wait()
	close(resultCh)
}

func listClusters() ([]string, error) {
	output, err := exec.Command("kubectl", "config", "get-clusters").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %v", output, err)
	}

	clusters := strings.Split(strings.Trim(string(output), "\n"), "\n")[1:]
	return clusters, nil
}
