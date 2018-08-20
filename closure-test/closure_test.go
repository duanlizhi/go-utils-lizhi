/**
 * <p>Description: (闭包测试) </>
 * @author lizhi_duan
 * @date 2018/7/30 08:09
 * @version 1.0
 */
package closure

import (
	"fmt"
	"testing"
)

func TestInt64Sum(t *testing.T) {
	var arr = []int64{1, 2, 3, 4, 5}
	var a = Int64Sum(arr)
	var arr1 = []int64{2, 3, 4, 5, 6}
	var a1 = Int64Sum(arr1)
	fmt.Println("a:", a())
	fmt.Println("a1", a1())
}
func TestCount(t *testing.T) {
	arr := Count()
	f1 := arr[0]
	f2 := arr[1]
	f3 := arr[2]
	fmt.Println("0:", f1()) //16
	fmt.Println("1:", f2()) //16
	fmt.Println("2", f3())  //16
}
func TestCount1(t *testing.T) {
	arr := Count1()
	f1 := arr[0]
	f2 := arr[1]
	f3 := arr[2]
	fmt.Printf("数组函数第一个结果:%v\n", f1()) //1
	fmt.Printf("数组函数第二个结果:%v\n", f2()) //4
	fmt.Printf("数组函数第三个结果:%v\n", f3()) //9
}
