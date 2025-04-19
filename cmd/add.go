package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
    Use:   "add [name]",
    Short: "Add a new password entry with username",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        user := internal.PromptPassword("Enter username: ")
        pwd := internal.PromptPassword("Enter password: ")
        err := internal.StoreEntry(name, user, pwd)
        if err != nil {
            fmt.Println("❌ Error saving entry:", err)
        } else {
            fmt.Println("✅ Entry saved.")
        }
    },
}

func init() {
    rootCmd.AddCommand(addCmd)
}
