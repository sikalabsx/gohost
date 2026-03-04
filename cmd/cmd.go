package cmd

import (
	"github.com/sikalabsx/gohost/cmd/root"
	_ "github.com/sikalabsx/gohost/cmd/version"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.Cmd.Execute())
}
