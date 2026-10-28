# 🔐 Passy

[![Build & Distribute](https://github.com/OcheOps/passy/actions/workflows/package-and-publish.yml/badge.svg)](https://github.com/OcheOps/passy/actions)

**A small, terminal-first password manager written in Go.**

Passy keeps your credentials in a local, encrypted file and gets them out of your way fast: one command to save a login, one command to copy a password to your clipboard. No accounts, no cloud, no daemon — just a single binary.

```text
$ passy add github
🔑 Master password:
Enter username:
Enter password:
✅ Entry saved.

$ passy get github
🔑 Master password:
👤 Username: octocat
📋 Password copied to clipboard!
```

---

## ✨ Features

- 🔒 **Encrypted at rest** — usernames and passwords are encrypted with AES-256-GCM using a key derived from your master password
- 📋 **Clipboard integration** — `get` copies the password straight to your clipboard instead of printing it
- 🎲 **Password generator** — create strong random passwords of any length with `gen`
- 👤 **Username + password per entry**, with in-place updates
- 📤 **CSV export** for backups or migrating to another manager
- 🙈 **Hidden input** — nothing you type at a prompt is echoed to the terminal
- 📦 **Installable via `apt` and `dnf`**, or build for Linux, macOS and Windows yourself

---

## 🚀 Installation

### 📦 APT (Debian / Ubuntu)

```bash
echo "deb [trusted=yes] https://ocheops.github.io/passy-apt-repo/ ./" | sudo tee /etc/apt/sources.list.d/passy.list
```

```bash
sudo apt update && sudo apt install passy
```

### 📦 DNF (Fedora / RHEL)

```bash
sudo dnf config-manager --add-repo https://ocheops.github.io/passy-rpm-repo/passsy.repo
```

```bash
sudo dnf install passy
```

Both packages install the binary to `/usr/local/bin/passy`.

### 🛠️ From source

Requires Go 1.23+. On Linux, the clipboard library also needs the X11 headers (`sudo apt install libx11-dev` on Debian/Ubuntu).

```bash
git clone https://github.com/OcheOps/passy.git && cd passy
```

```bash
go build -o passy . && sudo mv passy /usr/local/bin/
```

To cross-compile for Linux, macOS (Intel + Apple Silicon) and Windows in one go:

```bash
./build-all.sh   # binaries land in ./dist
```

---

## 🧪 Usage

Every command asks for your **master password** first. The first time you run Passy, whatever master password you choose becomes the key for your vault — use the same one every time after that.

| Command | What it does |
| --- | --- |
| `passy add <name>` | Save a new entry; prompts for username and password |
| `passy get <name>` | Show the username and copy the password to the clipboard |
| `passy gen <name> [-l N]` | Generate a random password (default 16 chars), save it under `<name>` and copy it |
| `passy list` | List the names of all saved entries |
| `passy update <name>` | Change the username and/or password (leave a prompt empty to keep the current value) |
| `passy delete <name>` | Remove an entry |
| `passy export <file>` | Write all entries to a CSV file (`Service,Username,Password`) |
| `passy help [command]` | Show help for any command |

### Examples

```bash
# Generate a 24-character password for a new account
passy gen aws-console -l 24

# Back up everything
passy export passy-backup.csv
```

> ⚠️ `export` writes your passwords in **plain text**. Store the file somewhere safe and delete it when you're done.

---

## 📁 Where your data lives

Passy stores its vault in a file called **`passy.db` in the current working directory**. Entry names are stored as-is; usernames and passwords are encrypted and hex-encoded. The file is created with `0600` permissions.

Because the path is relative, running Passy from a different directory gives you a different (empty) vault. The simplest approach is to always run it from the same place, e.g. with this shell function in your `~/.bashrc` or `~/.zshrc`:

```bash
passy() { (cd ~ && command passy "$@"); }
```

---

## 🛡️ Security notes

Passy is a small personal project, and it's worth being clear about what it does and doesn't protect against:

- The encryption key is a single SHA-256 hash of the master password (no salt or slow KDF), so a weak master password can be brute-forced if someone gets hold of `passy.db`. **Use a long, unique master password.**
- Encryption uses a fixed nonce rather than a random one per value.
- A wrong master password isn't rejected — entries just decrypt to empty values.
- Entry names (e.g. `github`) are not encrypted.
- Copied passwords stay on the clipboard until you overwrite them.

Hardening these (Argon2/scrypt key derivation, random nonces, master-password verification, clipboard auto-clear) is on the wish list — contributions welcome.

---

## ⚙️ Development

```text
passy/
├── main.go            # entry point
├── cmd/               # one file per CLI command (Cobra)
├── internal/
│   ├── vault.go       # encryption, storage, clipboard
│   ├── input.go       # hidden terminal prompts
│   └── utils.go       # password generator
├── build-all.sh       # cross-platform build script
└── .github/workflows/ # packaging & repo publishing
```

Run without installing:

```bash
go run . list
```

### Releasing

Pushing a tag like `v1.2.0` triggers the GitHub Actions workflow, which builds the binary, packages it as `.deb` and `.rpm` with [fpm](https://github.com/jordansissel/fpm), and publishes them to the [APT](https://github.com/OcheOps/passy-apt-repo) and [RPM](https://github.com/OcheOps/passy-rpm-repo) repositories.

```bash
git tag v1.2.0 && git push origin v1.2.0
```

---

## ❤️ Contributing

Issues and pull requests are welcome! Ideas that would make a big difference:

- Stronger key derivation and random nonces (see [Security notes](#️-security-notes))
- A fixed vault location (e.g. `~/.config/passy/`) or a `--vault` flag
- Clipboard auto-clear
- Import from CSV / other password managers
- Cloud sync or a GUI

---

## 📜 License

[MIT](LICENSE) — use it, hack it, ship it.

---

Built with Go & ❤️ by [@OcheOps](https://github.com/OcheOps)
