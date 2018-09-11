/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/9/11 06:40
 * @version 1.0
 */
package transform

/**
 * <p>Description: (将整数转换成2进制数) </p>
 * @author lizhi_duan
 * @date 2018/9/11 6:41
 */
func ConvertToBin(n int64) (bin string) {
	result := ""
	if n == 0 {
		result = BaseType2String(0)
	}
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = BaseType2String(lsb) + result
	}

	return result
}
