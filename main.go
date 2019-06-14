// im-infra is a command line tool for managing docker-cloud infrastructure and deployment
package main

import (
	"fmt"
	"os"

	"github.com/daneroo/im-infra/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
