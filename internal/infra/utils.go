package infra

import "fmt"

func generateInstanceNames(component string, count int) []string {
	if component == "worker" || component == "workers" {
		names := []string{}
		for i := 1; i <= count; i++ {
			names = append(names, fmt.Sprintf("worker-%d", i))
		}
		return names
	}

	return []string{component}
}
