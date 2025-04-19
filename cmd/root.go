package cmd

import (
    "fmt"
    "os"
    "passy/internal"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "passy",
    Short: "Passy is a simple CLI password manager",
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        master := internal.PromptPassword("🔑 Master password: ")
        internal.InitVault(master)
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
