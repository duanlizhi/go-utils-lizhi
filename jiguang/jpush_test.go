/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/9 21:53
 * @version 1.0
 */
package jiguang

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewPushObject(t *testing.T) {
	object, _ := NewPushObject(nil, nil)
	byte, _ := json.Marshal(object)
	fmt.Println(string(byte))
}
