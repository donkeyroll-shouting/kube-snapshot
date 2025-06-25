package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

const FlagNameKubeconfig = "kubeconfig"

var rootCmd = &cobra.Command{
	Use:   "ksnap",
	Short: "A Kubernetes snapshot tool",
	Long:  `ksnap is a CLI tool for taking snapshots of Kubernetes clusters.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Taking a snapshot of the Kubernetes cluster...")
	},
}

func main() {
	// Add k8s generic CLI flags.
	configFlags := genericclioptions.NewConfigFlags(true)
	configFlags.AddFlags(rootCmd.Flags())

	// Loop through the flags and mark them as hidden unless they are "kubeconfig".
	rootCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Name != FlagNameKubeconfig {
			flag.Hidden = true
		}
	})

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
