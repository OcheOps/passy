package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
    Use:   "get [name]",
    Short: "Retrieve username and password for an entry",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        user, pwd, err := internal.RetrieveEntry(name)
        if err != nil {
            fmt.Println("❌", err)
            return
        }
        fmt.Println("👤 Username:", user)
        fmt.Println("📋 Password copied to clipboard!")
        internal.CopyToClipboard(pwd)
    },
}

func init() {
    rootCmd.AddCommand(getCmd)
}
