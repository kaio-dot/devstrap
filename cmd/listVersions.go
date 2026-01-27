package cmd

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/platform"
	"github.com/kaio-dot/devstrap/internal/providers"

	"github.com/spf13/cobra"
)

var listVersionCmd = &cobra.Command{
	Use:     "lv [tool]",
	Aliases: []string{"list-version [tool]"},
	Short:   "Lista as versões disponíveis de uma ferramenta",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		p := platform.DetectPlatform()

		if !p.IsSupported() {
			fmt.Println("Plataforma não suportada:", p.OS, p.Arch)
			return
		}
		tool := args[0]

		provider, ok := providers.GetProvider(tool)

		if !ok {
			fmt.Println("Ferramenta ainda não suportada, tenha calma dev!")

			return
		}

		versions, err := provider.ListAvailableVersions(20, false)

		if err != nil {
			fmt.Println("Deu um erro ao listar as versões. Tenta de novo aí, blz?", err)
			return
		}

		fmt.Printf("Versões disponíveis para %s:\n", tool)

		for _, v := range versions {
			fmt.Println("\n", v)
		}
	},
}

func init() {
	rootCmd.AddCommand(listVersionCmd)
}
