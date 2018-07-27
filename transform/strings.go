/**
 * <p>Description: (字符串转换工具类) </>
 * @author lizhi_duan
 * @date 2018/7/27 08:10
 * @version 1.0
 */
package transform

import "strconv"

func BaseType2String(base interface{}) string {
	//reflect.TypeOf() go使用反射获取类型，或者使用断言assert
	switch base.(type) {
	case string:
		return base.(string)
	case int:
		return strconv.Itoa(base.(int))
	case int64:
		return strconv.FormatInt(base.(int64),10)
	case float64:
		return strconv.FormatFloat(base.(float64),'E',-1,64)
	case float32:
		return strconv.FormatFloat(base.(float64),'E',-1,32)
	default:
		return ""
	}

}
