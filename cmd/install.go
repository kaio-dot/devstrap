package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install",
	Aliases: []string{"i"},
	Short:   "Instala a ferramenta em ambiente local",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tool := args[0]
		fmt.Println("Iniciando a instalação")
		fmt.Println("Instalando:", tool)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}
