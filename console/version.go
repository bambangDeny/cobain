package console

import (
	"fmt"

	"github.com/bambangDeny/cobain/config"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print fer version",
	Long:  `print version of fer`,
	Run:   printVersion,
}

func printVersion(cmd *cobra.Command, args []string) {
	fmt.Println(config.Version)
}
