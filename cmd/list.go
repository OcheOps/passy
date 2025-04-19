package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use:   "list",
    Short: "List all saved entries",
    Run: func(cmd *cobra.Command, args []string) {
        for k := range internal.ListEntries() {
            fmt.Println("🔐", k)
        }
    },
}

func init() {
    rootCmd.AddCommand(listCmd)
}
