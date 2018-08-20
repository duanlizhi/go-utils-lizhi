/**
 * <p>Description: (帮助理解闭包的概念) </>
 * @author lizhi_duan
 * @date 2018/7/30 08:00
 * @version 1.0
 */
package closure

import (
	"log"
)

func Int64Sum(arr []int64) func() int64 {
	return func() int64 {
		var sum = int64(0)
		for k, v := range arr {
			log.Println(k, v)
			sum += v
		}
		return sum
	}
}

/**
 * <p>Description: (闭包中不能使用循环变量) </p>
 * <p>Description: (此包中引用了循环变量，导致最后函数执行时循环变量的值都为最后的4) </p>
 * @author lizhi_duan
 * @date 2018/7/30 21:28
 */
func Count() []func() int {
	var arr []func() int
	for i := 1; i <= 3; i++ {
		arr = append(arr, func() int {
			log.Println(i, "------", i*i)
			return i * i
		})
	}
	return arr
}

/**
 * <p>Description: (如果一定要引用循环变量怎么办？方法是再创建一个函数，用该函数的参数绑定循环变量当前的值，无论该循环变量后续如何更改，已绑定到函数参数的值不变：) </p>
 * <p>Description: 匿名函数
	func(n int) func() int {
			return func() int {
				log.Println("匿名函数中绑定的n：", n)
				return n * n
			}
		}(i)
	</p>
 * @author lizhi_duan
 * @date 2018/7/30 22:20
*/
func Count1() []func() int {
	var arr []func() int
	for i := 1; i <= 3; i++ {
		arr = append(arr, func(n int) func() int {
			return func() int {
				log.Println("匿名函数中绑定的n：", n)
				return n * n
			}
		}(i))
	}
	return arr
}
