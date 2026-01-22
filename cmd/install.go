package cmd

import (
	"fmt"

	"github.com/kaio-dot/devstrap/internal/providers"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:     "install [tool]",
	Aliases: []string{"i"},
	Short:   "Instala a ferramenta em ambiente local",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		tool := args[0]

		provider, ok := providers.GetProvider(tool)

		if !ok {
			fmt.Println("Ferramenta ainda não suportada.Tenha paciência dev!", tool)
			fmt.Printf("tool: '%s'\n", tool)
			fmt.Printf("registry: %#v\n", providers.DebugRegistry())

			return
		}

		err := provider.Install("latest")
		if err != nil {
			fmt.Println("Não sei por que, mas deu erro. Adivinha aí:", err)
		}
	},
}

func init() {
	fmt.Println("NodeProvider registrado")
	rootCmd.AddCommand(installCmd)
}
