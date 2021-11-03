package bit_operation

import (
	"fmt"
	"math"
)

//位运算

//给定两个整数 a 和 b ，求它们的除法的商 a/b ，要求不得使用乘号 '*'、除号 '/' 以及求余符号 '%' 。
//注意：
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2
//假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231, 231−1]。本题中，如果除法结果溢出，则返回

//切入点1：不准使用*、/、%运算符，那么可以考虑二进制正数位移运算(避免超时)
//切入点2：截取小数部分
//切入点3:限制为int32类型，如果除法溢出则返回2^31 − 1,可以使用math.MaxInt32
func divide(a int, b int) int {

	isPositive := true
	if a < 0 {
		isPositive = !isPositive
		a = -a
	}
	if b < 0 {
		isPositive = !isPositive
		b = -b
	}

	//1.a<b，返回0
	if a < b {
		return 0
	}
	//2.a>=b,位移后用被除数-除数直到不能减为止
	//17 = 5*2^1+5*2^0
	//25 = 5^2^2+5*2^1
	//count为最大项的指数
	res := 0
	for b <= a {
		count := 0

		for b<<count<<1 < a {
			count++
		}
		fmt.Println(count)
		for b<<count <= a {
			res = 1<<count + res
			a = a - b<<count
		}
	}

	//3.负数处理
	if !isPositive {
		res = -res
	} else {
		//4.正号溢出
		if res > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return res

}

func main() {
	fmt.Println(divide(17, 5))
}
