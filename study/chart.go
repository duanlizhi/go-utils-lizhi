/**
 * <p>Description: (一句话描述一下该文件的作用) </>
 * @author lizhi_duan
 * @date 2018/8/2 22:45
 * @version 1.0
 */
package study

import (
	"bufio"
	"fmt"
	"os"
)

func Chart() {
	//从标准输入中读取数据
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	//读取数据直到碰到 \n 为止
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred: %s\n", err)
		//异常退出
		os.Exit(1)
	} else {
		//用切片操作删除最后的\n
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! What can I do for you?\n", name)
	}
}
