package root

import (
	"github.com/sikalabsx/gohost/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "gohost",
	Short: "gohost, " + version.Version,
}
