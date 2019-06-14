package cmd

import (
	"fmt"
	"strings"

	"github.com/daneroo/im-infra/dc"
	"github.com/spf13/cobra"
)

// upCmd represents the up command
var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Bring up the infrastructure",
	Long:  `Bring up the infrastructure`,
	Run:   doit,
}

func init() {
	RootCmd.AddCommand(upCmd)
}

func doit(cmd *cobra.Command, args []string) {
	fmt.Println("up called: " + strings.Join(args, " "))

	dc.Up()
	// cmd.PersistentFlags.

}
