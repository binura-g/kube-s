package search

import (
	"bufio"
	"bytes"
	"log"
	"strings"
	"sync"
)

type Result struct {
	Cluster string
	Ns      string
	Name    string
}

type kubectlOutput struct {
	Data    []byte
	Cluster string
}

func Run(kind, pattern string, resultCh chan Result) {
	clusters, err := listClusters()
	if err != nil {
		log.Fatal(err)
	}

	kubectlOutputCh := make(chan kubectlOutput, len(clusters))
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(clusters))

		for _, cluster := range clusters {
			go func(cluster string) {
				defer wg.Done()
				output, err := listResources(cluster, kind)
				if err != nil {
					log.Printf("Error: %q %s\n", cluster, output)
				}
				kubectlOutputCh <- kubectlOutput{output, cluster}
			}(cluster)
		}
		wg.Wait()
		close(kubectlOutputCh)
	}()

	var resultWg sync.WaitGroup
	for output := range kubectlOutputCh {
		resultWg.Add(1)

		go func(output kubectlOutput) {
			defer resultWg.Done()

			scanner := bufio.NewScanner(bytes.NewReader(output.Data))
			for scanner.Scan() {
				s := scanner.Text()
				if strings.Contains(s, pattern) {
					split := strings.Split(s, " ")
					resultCh <- Result{
						Cluster: output.Cluster,
						Ns:      split[0],
						Name:    split[len(split)-1],
					}
				}
			}
		}(output)
	}

	resultWg.Wait()
	close(resultCh)
}
