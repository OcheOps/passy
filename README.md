# 🔐 Passy — A Simple Terminal Password Manager

[![Build & Distribute](https://github.com/OcheOps/passy/actions/workflows/package-and-publish.yml/badge.svg)](https://github.com/OcheOps/passy/actions)

Passy is a lightweight, secure, and terminal-first password manager written in Go.  
It supports clipboard copying, username/password storage, entry updates, file export, and more — all in a fast CLI interface.

---

## ✨ Features

- 🔒 AES-GCM encrypted local storage
- 📋 Clipboard integration (with auto-clear)
- 📅 Export vault to plaintext
- 👤 Stores username + password per entry
- 🧄 Simple file-based vault (`~/.passy-vault.json`)
- 📦 Installable via `apt` or `dnf` package managers
- ⚡ Cross-platform binaries available for Windows, Linux, and macOS

---

## 🚀 Installation

### 🐧 Linux (.deb / .rpm)

#### 📦 APT (Debian/Ubuntu)
```bash
echo "deb [trusted=yes] https://ocheops.github.io/passy-apt-repo/ ./" | sudo tee /etc/apt/sources.list.d/passy.list
sudo apt update
sudo apt install passy
```

#### 📦 DNF (Fedora/RHEL)
```bash
sudo dnf config-manager --add-repo https://ocheops.github.io/passy-rpm-repo/passsy.repo
sudo dnf install passy
```

---

### 💻 Binaries

Download prebuilt binaries from [Releases](https://github.com/OcheOps/passy/releases):

```bash
# Example for Linux
wget https://github.com/OcheOps/passy/releases/download/v1.2.0/passy-linux-amd64
chmod +x passy-linux-amd64
sudo mv passy-linux-amd64 /usr/local/bin/passy
```

---

## 🥪 Usage

```bash
# Add a new password
passy add github

# Retrieve an entry and copy password to clipboard
passy get github

# List all entries
passy list

# Update an entry
passy update github

# Delete an entry
passy delete github

# Export vault to file
passy export vault.txt
```

---

## 📁 File Location

All entries are stored securely in:

```
~/.passy-vault.json
```

Encrypted using AES-GCM with your generated key.

---

## ⚙️ Development

### Build locally:
```bash
go build -o passy main.go
```

### Test:
```bash
go run main.go list
```

---

## ❤️ Contribute

Wanna help with cloud sync, GUI (Tauri), or encryption enhancements?  
PRs and issues are welcome!

---

## 📜 License

MIT — use it, hack it, ship it.

---

Built with Go & love by [@OcheOps](https://github.com/OcheOps) 🚰