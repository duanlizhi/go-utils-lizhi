/**
 * <p>Description: (字符串转换工具类) </>
 * @author lizhi_duan
 * @date 2018/7/27 08:10
 * @version 1.0
 */
package transform

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
)

/**
 * <p>Description: (根据传入的基础数据类型，比如int、int64、uint32等转string) </p>
 * @author lizhi_duan
 * @date 2018/7/27 23:50
 */
func BaseType2String(base interface{}) string {
	//reflect.TypeOf() go使用反射获取类型，或者使用断言assert
	fmt.Println("type:", reflect.TypeOf(base))
	switch base.(type) {
	case string:
		return base.(string)
	case int:
		return strconv.Itoa(base.(int))
	case int64:
		return strconv.FormatInt(base.(int64), 10)
	case float64:
		// 'b' (-ddddp±ddd，二进制指数)
		// 'e' (-d.dddde±dd，十进制指数)
		// 'E' (-d.ddddE±dd，十进制指数)
		// 'f' (-ddd.dddd，没有指数)
		// 'g' ('e':大指数，'f':其它情况)
		// 'G' ('E':大指数，'f':其它情况)
		return strconv.FormatFloat(base.(float64), 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(base.(float64), 'f', -1, 32)
	default:
		return ""
	}
}

/**
 * <p>Description: (将string转成int) </p>
 * @author lizhi_duan
 * @date 2018/7/27 23:58
 */
func String2Int(s string) int {
	//fmt.Println(reflect.TypeOf(s))
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal("string to int failed !", err)
	}
	return i
}
func String2Int32(s string) int32 {
	//第二个参数为基数（2~36），第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		log.Fatal("string to int64 failed!", err)
	}
	return int32(i)
}
func String2Int64(s string) int64 {
	//第二个参数为基数（2~36），第三个参数位大小表示期望转换的结果类型，其值可以为0, 8, 16, 32和64，分别对应 int, int8, int16, int32和int64
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatal("string to int64 failed!", err)
	}
	return i
}
