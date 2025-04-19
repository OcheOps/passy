package internal

import (
    "fmt"
    "golang.org/x/term"
    "syscall"
)

func PromptPassword(prompt string) string {
    fmt.Print(prompt)
    bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    fmt.Println()
    if err != nil {
        return ""
    }
    return string(bytePassword)
}
