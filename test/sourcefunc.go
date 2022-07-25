package test

//待测试的函数
func Fib(n int)  int {
	if n< 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
