package decrypt

import (
	"crypto/sha256"
	"github.com/gek64/gek/gCrypto"
	"github.com/gek64/gek/gCrypto/padding"
	"golang.org/x/crypto/chacha20poly1305"
	"os"
)

// FromBytes 从比特切片解密
func FromBytes(ciphertext []byte, key []byte) (plaintext []byte, err error) {
	// 通过密钥长度判断是否使用解密
	switch len(key) {
	case 0:
		return ciphertext, nil
	default:
		key = padding.ZeroPadding(key, chacha20poly1305.KeySize)
		key = key[0:chacha20poly1305.KeySize]
		return gCrypto.NewChaCha20Poly1305WithHashAD(key, sha256.New()).Decrypt(ciphertext)
	}
}

// FromFile 从文件解密
func FromFile(filepath string, encryptionKey []byte) (plaintext []byte, err error) {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return FromBytes(d, encryptionKey)
}
