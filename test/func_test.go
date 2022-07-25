package test

import "testing"


//单元测试
func TestFib(t *testing.T) {
	//单个值进行测试
	var (
		in = 7
		result = 13
	)
	actual := Fib(in)
	if actual != result {
		t.Errorf("Fib(%d) = %d; Result %d", in, actual, result)
	}
}

func TestFibMore(t *testing.T) {
	//定义多个测试值的接口体
	var fibTest = []struct{
		in int
		result int
	}{
		{1,1},
		{2,1},
		{3,2},
		{4,3},
		{5,6},
	}

     //进行测试
     for _,test := range fibTest {
     	actual := Fib(test.in)
     	if actual != test.result {
     		t.Errorf("Fib(%d) = %d; Result %d", test.in, actual, test.result)
		}

	 }

}

//基准测试
func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(10)
	}
}



