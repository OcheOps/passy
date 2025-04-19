package cmd

import (
    "fmt"
    "passy/internal"
    "github.com/spf13/cobra"
)

var length int

var genCmd = &cobra.Command{
    Use:   "gen [name]",
    Short: "Generate a strong random password",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        password := internal.GeneratePassword(length)
        internal.StoreEntry(name, "", password)
        internal.CopyToClipboard(password)
        fmt.Println("🔐 Generated and copied password for:", name)
    },
}

func init() {
    genCmd.Flags().IntVarP(&length, "length", "l", 16, "Password length")
    rootCmd.AddCommand(genCmd)
}
