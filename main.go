package cmd

import (
	log "github.com/sirupsen/logrus"
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
}

func (c *Commander) setLogger() {
	logLevel, _ := c.Root.Flags().GetString("logLevel")

	level, err := log.ParseLevel(logLevel)
	if err != nil {
		cobra.CheckErr(err)
	}

	log.SetLevel(level)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
		FullTimestamp:   true,
	})
}

// Init is entrypoint for the commands.
func (c *Commander) Init() *Commander {
	cobra.OnInitialize(func() {
		c.setLogger()
	})

	c.setDefaultFlags()

	c.SetPersistentFlags(func(c *Commander) {
		c.Root.PersistentFlags().String("test", "info", "Set the logging level. One of: debug|info|warn|error")
	})

	return c
}

// Execute executes Cobra commands.
func (c *Commander) Execute() error {
	return c.Root.Execute()
}
