package cmd

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omegion/go-cli/internal/info"
)

// Version prints version/build.
func Version(commander *Commander) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version/build number",
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Infoln(fmt.Sprintf("%s %s\n", info.AppName, info.Version))

			return nil
		},
	}

	return cmd
}
