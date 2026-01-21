package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devstrap",
	Short: "Bootstrap your dev environment",
}

var devStrapCmd = &cobra.Command{
	Use:   "devstrap",
	Short: "Mostra o banner em ASCII para o usuário",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
		
		██████╗	██████╗ ██║	  ██║███████╗████████ ██████  ███████   ███╗  
		██╔══██╗██╔════╝██║   ██║██╔════╝╚══██╔══╝██╔══██╗██╔══██╗██╔══██╗
		██║  ██║█████╗  ██║   ██║███████╗   ██║   ██████╔╝███████║██████╔╝
		██║  ██║██╔══╝  ╚██╗ ██╔╝╚════██║   ██║   ██╔══██╗██╔══██║██╔═══╝
		██████╔╝███████╗ ╚████╔╝ ███████║   ██║   ██║  ██║██║  ██║██║
		╚═════╝ ╚══════╝  ╚═══╝  ╚══════╝   ╚═╝   ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝
		____________________________________________________________________

		Bootstrap your dev environment with DEVSTRAP!
	`)
	},
}

func init() {
	rootCmd.AddCommand(devStrapCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
