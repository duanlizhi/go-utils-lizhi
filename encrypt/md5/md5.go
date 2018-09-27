/**
 * <p>Description: (md5加密) </>
 * @author lizhi_duan
 * @date 2018/9/26 18:34
 * @version 1.0
 */
package md5

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * <p>Description: (一句话描述该方法的作用) </p>
 * @param value 需要加密的明文字符串
 * @author lizhi_duan
 * @date 2018/9/26 18:38
 */
func Md5Encrypt(value string) string {
	md := md5.New()
	md.Write([]byte(value))
	cipherStr := md.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
