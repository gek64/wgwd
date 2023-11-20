package preload

import (
	"github.com/gek64/gek/gCrypto"
	"golang.org/x/crypto/chacha20poly1305"
)

const AssociatedDataSize = 8

func GetDecryptedPreload(ciphertext []byte, key []byte, associatedDataSize uint) (preload []byte, err error) {
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
