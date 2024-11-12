package cmd

import (
	"fmt"

	"github.com/gkwa/anyhobbit/core"
	"github.com/spf13/cobra"
)

var zooCmd = &cobra.Command{
	Use:   "zoo",
	Short: "Show all animal configurations",
	Long:  "Show all available Renovate configurations ordered by animal name",
	RunE: func(cmd *cobra.Command, args []string) error {
		configs, err := core.ListAllConfigs()
		if err != nil {
			return err
		}

		for _, c := range configs {
			fmt.Printf("[%s]\n", c.Name)
			for _, line := range c.Config {
				if line != "" {
					fmt.Printf("[%s] %s\n", c.Name, line)
				}
			}
			fmt.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(zooCmd)
}
