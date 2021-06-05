package cmd

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestCommander_NewCommander_Empty(t *testing.T) {
	commander := Commander{}

	assert.Equal(t, (*cobra.Command)(nil), commander.Root)
}

func TestCommander_SetCommand(t *testing.T) {
	commander := NewCommander(&cobra.Command{
		Use:          "test",
		Short:        "Test.",
		Long:         "Test longer.",
		SilenceUsage: true,
	})

	commands := []*cobra.Command{
		{
			Use: "cmd-1",
		},
		{
			Use: "cmd-2",
		},
	}

	commander.SetCommand(commands...)

	assert.Len(t, commander.Root.Commands(), 2)
}

func TestCommander_Setup(t *testing.T) {
	commander := NewCommander(&cobra.Command{
		Use:          "test",
		Short:        "Test.",
		Long:         "Test longer.",
		SilenceUsage: true,
	})

	commander.Init()

	commander.Root.SetArgs([]string{"version"})

	_, err := commander.Root.ExecuteC()

	expectedLogLevelFlag, _ := commander.Root.Flags().GetString("logLevel")
	expectedLogFormatFlag, _ := commander.Root.Flags().GetString("logFormat")

	assert.Equal(t, nil, err)
	assert.Equal(t, "info", log.GetLevel().String())
	assert.Equal(t, "info", expectedLogLevelFlag)
	assert.Equal(t, "text", expectedLogFormatFlag)
}

func TestCommander_SetPersistentFlags(t *testing.T) {
	commander := NewCommander(&cobra.Command{
		Use:          "test",
		Short:        "Test.",
		Long:         "Test longer.",
		SilenceUsage: true,
	})

	commander.Init()

	commander.Root.SetArgs([]string{"version"})

	_, err := commander.Root.ExecuteC()

	expectedLogLevelFlag, _ := commander.Root.Flags().GetString("logLevel")

	assert.Equal(t, nil, err)
	assert.Equal(t, "info", log.GetLevel().String())
	assert.Equal(t, "info", expectedLogLevelFlag)
}
