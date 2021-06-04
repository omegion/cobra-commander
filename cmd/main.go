package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/go-cli/internal/client"
)

// Commander is a struct for command system.
type Commander struct {
	Root     *cobra.Command
	LogLevel string
}

// NewCommander is a factory for Commander.
func NewCommander() *Commander {
	return &Commander{}
}

// SetRootCommand sets Root command.
func (c *Commander) SetRootCommand() {
	c.Root = &cobra.Command{
		Use:          "go-cli",
		Short:        "Go CLI application template",
		Long:         "Go CLI application template for Go projects.",
		SilenceUsage: true,
	}
}

func (c *Commander) setPersistentFlags() {
	c.Root.PersistentFlags().String("logLevel", "info", "Set the logging level. One of: debug|info|warn|error")
}

func (c *Commander) setLogger() {
	c.LogLevel, _ = c.Root.Flags().GetString("logLevel")

	level, err := log.ParseLevel(c.LogLevel)
	if err != nil {
		cobra.CheckErr(err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
}

// Setup is entrypoint for the commands.
func (c *Commander) Setup() {
	cobra.OnInitialize(func() {
		c.setLogger()
	})

	c.SetRootCommand()
	c.setPersistentFlags()

	c.Root.AddCommand(Version(c))
	c.Root.AddCommand(Add(c))
	c.Root.AddCommand(Subtract(c))
}

// WithClient is a wrapper for testing.
func (c *Commander) WithClient(
	fn func(c client.Interface, cmd *cobra.Command, args []string) error,
) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		c := client.NewClient()

		return fn(c, cmd, args)
	}
}
