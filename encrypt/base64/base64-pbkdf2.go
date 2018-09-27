/**
 * <p>Description: (pbkdf2加密算法加盐后base64编码) </>
 * @author lizhi_duan
 * @date 2018/9/26 16:55
 * @version 1.0
 */
package base64

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/pbkdf2"
)

/**
 * <p>Description: (使用pbkdf2加密) </p>
 * @author lizhi_duan
 * @date 2018/9/26 17:20
 */
func pbkdf2Key(key string) ([]byte, []byte) {
	salt := make([]byte, 32)
	rand.Reader.Read(salt)
	bytes := pbkdf2.Key([]byte(key), salt, 1024, 32, sha256.New)
	return bytes, salt
}

/**
 * <p>Description: (用于验证pbkdf2加密传是否正确传入) </p>
 * @author lizhi_duan
 * @date 2018/9/26 17:26
 */
func pbkdf2Decode(key string, saltKey string) []byte {
	salt, _ := base64.StdEncoding.DecodeString(saltKey)
	bytes := pbkdf2.Key([]byte(key), salt, 1024, 32, sha256.New)
	return bytes
}

/**
 * <p>Description: (基于pbkdf2算法加盐生成密码，最后进行base64编码的key值和加盐值的key) </p>
 * @param key 需要加密的明文字符串
 * @author lizhi_duan
 * @date 2018/9/26 16:58
 */
func CreatePbkdf2Base64(key string) (base64Key string, saltKey string) {
	bytes, salt := pbkdf2Key(key)
	base64Key = base64.StdEncoding.EncodeToString(bytes)
	saltKey = base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("明文加密后的值（base64编码）：%s\n加密使用的盐值（base64编码）: %s\n", base64Key, saltKey)
	return
}

/**
 * <p>Description: (检测经pbkdf2加密字符串后base64key是否匹配) </p>
 * @param plaintext 明文
 * @param base64Key 加密后的key值
 * @param saltKey base64转码后的盐值
 * @author lizhi_duan
 * @date 2018/9/26 17:16
 */
func CheckPbkdf2Base64(plaintext string, base64Key string, saltKey string) (flag bool) {
	flag = false
	decode := pbkdf2Decode(plaintext, saltKey)
	if key := base64.StdEncoding.EncodeToString(decode); key == base64Key {
		flag = true
	}
	return
}
