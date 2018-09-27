/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/26 17:09
 * @version 1.0
 */
package base64

import (
	"fmt"
	"testing"
)

func TestCreatePbkdf2Base64(t *testing.T) {
	CreatePbkdf2Base64("123456")
}

func TestCheckPbkdf2Base64(t *testing.T) {
	key := "123456"
	base64Key := "EfZBvFoRwGfF7LU+vuuUY6YfQWziRZnkXqhc0Hg9td4="
	saltKey := "RUPljFHBRom4dlkqzL/6kMpLbu871RuBFHqM+VZYBfY="
	flag := CheckPbkdf2Base64(key, base64Key, saltKey)
	fmt.Printf("验证是否正确：%v\n", flag)
}
