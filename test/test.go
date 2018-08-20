/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/7/31 08:00
 * @version 1.0
 */
package test

import (
	"fmt"
	"reflect"
)

var v = "1,2,3"

func V() {
	fmt.Printf("外部v实体%v,类型是%v\n", v, reflect.TypeOf(v))
	v := []int{1, 2, 3}
	fmt.Printf("外部v实体%v,类型是%v\n", v, reflect.TypeOf(v))
	if v != nil {
		v := 123
		fmt.Printf("内部重新赋值v:%v,类型是%v\n", v, reflect.TypeOf(v))
	}
	fmt.Printf("外部v实体%v,类型是%v\n", v, reflect.TypeOf(v))
}
