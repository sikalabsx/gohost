package host

import (
	"fmt"
	"net"
	"os"

	"github.com/sikalabsx/gohost/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "host <hostname>",
	Short: "DNS lookup",
	Args:  cobra.ExactArgs(1),
	Run: func(c *cobra.Command, args []string) {
		hostname := args[0]

		addrs, err := net.LookupHost(hostname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		for _, addr := range addrs {
			fmt.Println(addr)
		}
	},
}

func init() {
	root.Cmd.AddCommand(Cmd)
}
