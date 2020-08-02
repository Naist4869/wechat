package pkcs7

import "fmt"

const (
	BlockSize = 32            // PKCS#7
	BlockMask = BlockSize - 1 // BLOCK_SIZE 为 2^n 时, 可以用 mask 获取针对 BLOCK_SIZE 的余数
)

func Pad(rawXMLMsg []byte, appId string) (int, []byte) {
	appIdOffset := 20 + len(rawXMLMsg)
	contentLen := appIdOffset + len(appId)
	amountToPad := BlockSize - contentLen&BlockMask
	plaintextLen := contentLen + amountToPad
	plaintext := make([]byte, plaintextLen)
	// PKCS#7 补位
	for i := contentLen; i < plaintextLen; i++ {
		plaintext[i] = byte(amountToPad)
	}
	return appIdOffset, plaintext
}

func Unpad(x []byte) ([]byte, error) {
	// last byte is number of suffix bytes to remove
	n := int(x[len(x)-1])
	if n < 1 || n > BlockSize {
		return nil, fmt.Errorf("the amount to pad is incorrect: %d", n)
	}
	return x[:len(x)-n], nil
}
