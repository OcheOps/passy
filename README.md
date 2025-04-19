# Passy 🔐

A lightweight, secure, and offline password manager for the terminal — built with Go.

---

## ✅ Features

- Add, get, list, delete, and update credentials
- Stores both **username** and **password**
- AES-256 encrypted vault file
- Master password authentication
- Secure password generation
- Clipboard copy (no terminal echo)
- Export credentials to CSV

---

## 🚀 Commands

| Command                | Description                           |
|------------------------|---------------------------------------|
| `passy add <name>`     | Add a new entry                       |
| `passy get <name>`     | Retrieve and copy password            |
| `passy list`           | Show all entry names                  |
| `passy delete <name>`  | Delete an entry                       |
| `passy gen <name>`     | Generate random password              |
| `passy update <name>`  | Update username or password           |
| `passy export <file>`  | Export all credentials to CSV         |

---

## 🧪 Example

```bash
$ passy add github
🔑 Master password: ********
👤 Enter username: johndoe
🔐 Enter password: ********
✅ Entry saved.

$ passy get github
🔑 Master password: ********
👤 Username: johndoe
📋 Password copied to clipboard!

$ passy export mysecrets.csv
✅ Exported to: mysecrets.csv
```

---

## 🛠️ Install

```bash
go build -o passy main.go
./passy
```

> Requires: Go 1.18+, `golang.design/x/clipboard`, and `github.com/spf13/cobra`

---

## 🔒 Security Note

Your vault is encrypted with AES-256 using your master password as a key. Make sure you do **not lose** your master password.

---

## 📦 TODO

- GUI shell (maybe with TUI or WebView)
- Cloud backup (opt-in)
- Touch ID/FaceID support

---

Made with ❤️ by Oche.
