package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
    Use:   "delete [name]",
    Short: "Delete a password entry",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        err := internal.DeleteEntry(args[0])
        if err != nil {
            fmt.Println("❌", err)
        } else {
            fmt.Println("🗑️ Deleted:", args[0])
        }
    },
}

func init() {
    rootCmd.AddCommand(deleteCmd)
}
