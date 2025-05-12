package cmd

import (
	"cloudctl/internal/infra"
	"fmt"

	"github.com/spf13/cobra"
)

var workerCount int

var createWorkersCmd = &cobra.Command{
	Use:   "workers",
	Short: "Create cluster workers --count (max 3)",
	RunE: func(cmd *cobra.Command, args []string) error {

		if workerCount < 1 || workerCount > 3 {
			return fmt.Errorf("--count must be between 1 - 3")
		}

		if err := infra.CreateComponent("worker", workerCount); err != nil {
			return fmt.Errorf("worker creation failed: %w", err)
		}

		fmt.Printf("[cloudctl] Created %d worker node(s). \n", workerCount)
		return nil
	},
}

func init() {
	createWorkersCmd.Flags().IntVar(&workerCount, "count", 1, "Num workers (1-3)")
	createCmd.AddCommand(createWorkersCmd)
}
