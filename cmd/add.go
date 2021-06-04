package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/go-cli/internal/client"
)

// Add adds two numbers.
func Add(commander *Commander) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add [num1] [num2]",
		Short:   "Adds two numbers",
		Long:    "Adds two numbers",
		Example: "  add --num1 1 --num2 2",
		RunE:    commander.WithClient(addNumbersE),
	}

	cmd.Flags().Int("num1", 0, "Number 1")
	cmd.Flags().Int("num2", 0, "Number 2")

	return cmd
}

func addNumbersE(c client.Interface, cmd *cobra.Command, args []string) error {
	num1, _ := cmd.Flags().GetInt("num1")
	num2, _ := cmd.Flags().GetInt("num2")

	result := c.AddNumbers(client.NumbersOptions{
		Number1: num1,
		Number2: num2,
	})

	log.Infoln(result)

	return nil
}
