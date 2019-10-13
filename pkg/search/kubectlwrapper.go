package search

import (
	"fmt"
	"os/exec"
	"strings"
)

func listClusters() ([]string, error) {
	output, err := exec.Command("kubectl", "config", "get-clusters").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s %v", output, err)
	}

	clusters := strings.Split(strings.Trim(string(output), "\n"), "\n")[1:]
	return clusters, nil
}

func listResources(cluster, kind string) ([]byte, error) {
	kubectlArgs := []string{
		"get", kind,
		"--all-namespaces",
		"--no-headers",
		"-o=custom-columns=NAMESPACE:.metadata.namespace,NAME:.metadata.name",
		fmt.Sprintf("--context=%s", cluster),
	}
	return exec.Command("kubectl", kubectlArgs...).CombinedOutput()
}
