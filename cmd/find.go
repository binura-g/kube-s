/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Usage:\n\tkube-s <ResourceKind> <Pattern>")
			os.Exit(1)
		}
		kind := args[0]
		pattern := args[1]

		output, err := exec.Command("kubectl", "config", "get-clusters").Output()
		if err != nil {
			log.Fatalf("Failed to retrieve cluster list from kubectl: %v\n", err)
		}

		clusters := strings.Split(strings.Trim(string(output), "\n"), "\n")[1:]
		//log.Printf("searching for %s/%s in %d clusters...", kind, pattern, len(clusters))

		var wg sync.WaitGroup

		for _, cluster := range clusters {
			wg.Add(1)
			go func(cluster string) {
				defer wg.Done()

				if err := exec.Command("kubectl", "config", "use-context", cluster).Run(); err != nil {
					log.Fatalf("cannot switch to cluster %q: %v", cluster, err)
				}

				c1 := exec.Command("kubectl", "get", kind, "--all-namespaces",  "--no-headers")
				c2 := exec.Command("findstr", pattern)

				pr, pw := io.Pipe()
				c1.Stdout = pw
				c2.Stdin = pr

				c2.Stdout = os.Stdout

				c1.Start()
				c2.Start()

				go func() {
					defer pw.Close()
					c1.Wait()
				}()

				c2.Wait()
			}(cluster)

			wg.Wait()
		}
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
