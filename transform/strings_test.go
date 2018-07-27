/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/7/27 08:45
 * @version 1.0
 */
package transform

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBaseType2String(t *testing.T) {
	var i = 8
	fmt.Print(reflect.TypeOf(i))
	string := BaseType2String(i)
	fmt.Println(string, reflect.TypeOf(string))
	var b = 33.55
	fmt.Println(reflect.TypeOf(b))
	sf := BaseType2String(b)
	fmt.Println(sf, reflect.TypeOf(sf))
}
func TestString2Int(t *testing.T) {
	var a = "2.35"
	fmt.Println(String2Int(a))
}
