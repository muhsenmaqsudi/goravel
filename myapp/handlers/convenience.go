package handlers

import "github.com/muhsenmaqsudi/goravel"

func (h *Handlers) encrypt(text string) (string, error) {
	enc := goravel.Encryption{Key: []byte(h.App.EncryptionKey)}
	encrypted, err := enc.Encrypt(text)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}

func (h *Handlers) decrypt(crypto string) (string, error) {
	enc := goravel.Encryption{Key: []byte(h.App.EncryptionKey)}

	encrypted, err := enc.Decrypt(crypto)
	if err != nil {
		return "", err
	}
	return encrypted, nil
}