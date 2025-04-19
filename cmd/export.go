package cmd

import (
    "encoding/csv"
    "fmt"
    "os"
    "passy/internal"
    "github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
    Use:   "export [file]",
    Short: "Export all entries to CSV (username, password)",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        file, err := os.Create(args[0])
        if err != nil {
            fmt.Println("❌ Could not create file:", err)
            return
        }
        defer file.Close()

        writer := csv.NewWriter(file)
        defer writer.Flush()

        writer.Write([]string{"Service", "Username", "Password"})
        for name := range internal.ListEntries() {
            user, pass, _ := internal.RetrieveEntry(name)
            writer.Write([]string{name, user, pass})
        }
        
        fmt.Println("✅ Exported to:", args[0])
    },
}

func init() {
    rootCmd.AddCommand(exportCmd)
}
