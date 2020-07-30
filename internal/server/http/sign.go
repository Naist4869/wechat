package http

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"wechat/internal/server/http/pkcs7"
)

// Sign 微信公众号 url 签名. https://www.programming-books.io/essential/go/hex-base64-encoding-b6e9fbb3165c4bcb907e469d86783aab
func Sign(strs ...string) (signature string) {
	sort.Strings(strs)
	tmpstr := strings.Join(strs, "")
	signature = fmt.Sprintf("%x", sha1.Sum([]byte(tmpstr)))
	return
}

// random  当前消息加密所用的 random, 16-bytes  cap为20 所以可以取出msgLen为[16:]
// rawXMLMsg  消息的明文文本, xml格式
// appId 当前消息加密所用的 AppId
func AESDecryptMsg(base64Msg []byte, aesKey []byte) (random, rawXMLMsg, appId []byte, err error) {
	var encryptedMsgLen int
	encryptedMsg := make([]byte, base64.StdEncoding.DecodedLen(len(base64Msg)))

	encryptedMsgLen, err = base64.StdEncoding.Decode(encryptedMsg, base64Msg)
	if err != nil {
		err = fmt.Errorf("AESDecryptMsg Decode base64Msg fail: %w", err)
		return
	}
	cipherText := encryptedMsg[:encryptedMsgLen]
	if len(cipherText) < pkcs7.BlockSize {
		err = fmt.Errorf("the length of ciphertext too short: %d", len(cipherText))
		return
	}
	if len(cipherText)&pkcs7.BlockMask != 0 {
		err = fmt.Errorf("ciphertext is not a multiple of the block size, the length is %d", len(cipherText))
		return
	}
	plaintext := make([]byte, len(cipherText)) // len(plaintext) >= BLOCK_SIZE
	// 解密
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		err = fmt.Errorf("AESDecryptMsg %w", err)
		return
	}
	mode := cipher.NewCBCDecrypter(block, aesKey[:16])
	mode.CryptBlocks(plaintext, cipherText)
	plaintext, err = pkcs7.Unpad(plaintext)
	if err != nil {
		err = fmt.Errorf("AESDecryptMsg %w", err)
		return
	}
	// 反拼接
	// len(plaintext) == 16+4+len(rawXMLMsg)+len(appId)
	if len(plaintext) <= 20 {
		err = fmt.Errorf("plaintext too short, the length is %d", len(plaintext))
		return
	}
	msgLen := binary.BigEndian.Uint32(plaintext[16:20])
	appIdOffset := 20 + msgLen
	if len(plaintext) <= int(appIdOffset) {
		err = fmt.Errorf("msg length too large: %d", msgLen)
		return
	}
	random = plaintext[:16:20]
	rawXMLMsg = plaintext[20:appIdOffset:appIdOffset]
	appId = plaintext[appIdOffset:]
	return
}

//AESEncryptMsg cipherText = AES_Encrypt[random(16B) + msg_len(4B) + rawXMLMsg + appId]
func AESEncryptMsg(random, rawXMLMsg []byte, appId string, aesKey []byte) (cipherText string, err error) {
	appIdOffset, plaintext := pkcs7.Pad(rawXMLMsg, appId)
	copy(plaintext[:16], random)
	binary.BigEndian.PutUint32(plaintext[16:20], uint32(len(rawXMLMsg)))
	copy(plaintext[20:], rawXMLMsg)
	copy(plaintext[appIdOffset:], appId)
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		err = fmt.Errorf("AESEncryptMsg %w", err)
		return
	}
	cipherTextBytes := make([]byte, len(plaintext))

	mode := cipher.NewCBCEncrypter(block, aesKey[:16])
	mode.CryptBlocks(cipherTextBytes, plaintext)
	cipherText = base64.StdEncoding.EncodeToString(cipherTextBytes)

	return
}
func makeNonce() (nonce string) {
	limit := big.NewInt(1)
	limit = limit.Lsh(limit, 64)
	n, err := rand.Int(rand.Reader, limit)
	if err != nil {
		return
	}
	nonce = n.String()
	return
}
