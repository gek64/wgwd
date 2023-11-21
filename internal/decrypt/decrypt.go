package decrypt

import (
	"github.com/gek64/gek/gCrypto"
	"golang.org/x/crypto/chacha20poly1305"
	"os"
)

const AssociatedDataSize = 8

// FromBytes 从比特切片解密
func FromBytes(ciphertext []byte, key []byte, associatedDataSize uint) (plaintext []byte, err error) {
	// 通过密钥长度判断是否使用解密
	switch len(key) {
	case 0:
		return ciphertext, nil
	default:
		key = gCrypto.KeyZeroPadding(key, chacha20poly1305.KeySize)
		key = gCrypto.KeyCropping(key, chacha20poly1305.KeySize)
		return gCrypto.NewChaCha20Poly1305(key, associatedDataSize).Decrypt(ciphertext)
	}
}

// FromFile 从文件解密
func FromFile(filepath string, encryptionKey []byte) (plaintext []byte, err error) {
	d, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return FromBytes(d, encryptionKey, AssociatedDataSize)
}
