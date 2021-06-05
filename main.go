package cmd

import (
	"github.com/spf13/cobra"
)

// Commander is a struct for command system.
type Commander struct {
	Root *cobra.Command
}

// NewCommander is a factory for Commander.
func NewCommander(cmd *cobra.Command) *Commander {
	return &Commander{Root: cmd}
}

// SetCommand sets Root command.
func (c *Commander) SetCommand(cmds ...*cobra.Command) *Commander {
	c.Root.AddCommand(cmds...)

	return c
}

// SetPersistentFlags sets persistent flags.
func (c *Commander) SetPersistentFlags(p func(c *Commander)) *Commander {
	p(c)

	return c
}

func (c *Commander) setDefaultFlags() {
	c.Root.PersistentFlags().String("logLevel", "info", "Set the logging level. One of: debug|info|warn|error")
	c.Root.PersistentFlags().String("logFormat", "text", "Set the logging format. One of: text|json (default \"text\")")
}

func (c *Commander) setLogger() {
	logLevel, _ := c.Root.Flags().GetString("logLevel")
	logFormat, _ := c.Root.Flags().GetString("logFormat")

	InitLogger(logLevel, logFormat)
}

// Init is entrypoint for the commands.
func (c *Commander) Init() *Commander {
	cobra.OnInitialize(func() {
		c.setLogger()
	})

	c.setDefaultFlags()

	return c
}

// Execute executes Cobra commands.
func (c *Commander) Execute() error {
	return c.Root.Execute()
}
