package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
    Use:   "update [name]",
    Short: "Update username or password for an entry",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        oldUser, _, err := internal.RetrieveEntry(name)
        if err != nil {
            fmt.Println("❌", err)
            return
        }

        fmt.Println("👤 Current username:", oldUser)
        newUser := internal.PromptPassword("🔁 New username (leave empty to keep): ")
        newPass := internal.PromptPassword("🔁 New password (leave empty to keep): ")

        if newUser == "" {
            newUser = oldUser
        }

        if newPass == "" {
            _, oldPass, _ := internal.RetrieveEntry(name)
            newPass = oldPass
        }

        err = internal.StoreEntry(name, newUser, newPass)
        if err != nil {
            fmt.Println("❌ Failed to update:", err)
        } else {
            fmt.Println("✅ Entry updated.")
        }
    },
}

func init() {
    rootCmd.AddCommand(updateCmd)
}
