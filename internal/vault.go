package internal

import (
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "os"

    "golang.design/x/clipboard"
)

var dbFile = "passy.db"
var vault map[string]Entry
var masterKey []byte

type Entry struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

func init() {
    clipboard.Init()
}

func InitVault(password string) {
    masterKey = deriveKey(password)
    vault = make(map[string]Entry)
    loadVault()
}

func deriveKey(password string) []byte {
    hash := sha256.Sum256([]byte(password))
    return hash[:]
}

func encrypt(text string) string {
    block, _ := aes.NewCipher(masterKey)
    aesGCM, _ := cipher.NewGCM(block)
    nonce := []byte("passypassypass")[:aesGCM.NonceSize()]
    ciphertext := aesGCM.Seal(nil, nonce, []byte(text), nil)
    return hex.EncodeToString(ciphertext)
}

func decrypt(enc string) string {
    data, _ := hex.DecodeString(enc)
    block, _ := aes.NewCipher(masterKey)
    aesGCM, _ := cipher.NewGCM(block)
    nonce := []byte("passypassypass")[:aesGCM.NonceSize()]
    plain, err := aesGCM.Open(nil, nonce, data, nil)
    if err != nil {
        return ""
    }
    return string(plain)
}

func loadVault() {
    data, err := os.ReadFile(dbFile)
    if err == nil {
        _ = json.Unmarshal(data, &vault)
    }
}

func saveVault() error {
    data, err := json.MarshalIndent(vault, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(dbFile, data, 0600)
}

func StoreEntry(name, username, password string) error {
    vault[name] = Entry{
        Username: encrypt(username),
        Password: encrypt(password),
    }
    return saveVault()
}

func RetrieveEntry(name string) (string, string, error) {
    enc, ok := vault[name]
    if !ok {
        return "", "", fmt.Errorf("entry '%s' not found", name)
    }
    return decrypt(enc.Username), decrypt(enc.Password), nil
}

func DeleteEntry(name string) error {
    if _, ok := vault[name]; !ok {
        return fmt.Errorf("entry '%s' not found", name)
    }
    delete(vault, name)
    return saveVault()
}

func ListEntries() map[string]Entry {
    return vault
}

func CopyToClipboard(text string) {
    clipboard.Write(clipboard.FmtText, []byte(text))
}
