package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/go-cli/internal/client"
)

// Subtract subtracts two numbers.
func Subtract(commander *Commander) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "subtract [num1] [num2]",
		Short:   "Subtracts two numbers",
		Long:    "Subtracts two numbers",
		Example: "  subtract --num1 1 --num2 2",
		RunE:    commander.WithClient(subtractNumbersE),
	}

	cmd.Flags().Int("num1", 0, "Number 1")
	cmd.Flags().Int("num2", 0, "Number 2")

	return cmd
}

func subtractNumbersE(c client.Interface, cmd *cobra.Command, args []string) error {
	num1, _ := cmd.Flags().GetInt("num1")
	num2, _ := cmd.Flags().GetInt("num2")

	result := c.SubtractNumbers(client.NumbersOptions{
		Number1: num1,
		Number2: num2,
	})

	log.Infoln(result)

	return nil
}
