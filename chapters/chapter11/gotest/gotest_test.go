/*
go test命令是一个按照一定的约定和组织的测试代码的驱动程序。在包目录内，所有以
_test.go为后缀名的源文件并不是go build构建包的一部分，它们是go test测试的一部分。

在*_test.go文件中，有三种类型的函数：测试函数、基准测试函数、示例函数。

一个测试函数是以Test为函数名前缀的函数，用于测试程序的一些逻辑行为是否正确；
go test命令会调用这些测试函数并报告测试结果是PASS或FAIL。

基准测试函数是以Benchmark为函数名前缀的函数，它们用于衡量一些函数的性能；
go test命令会多次运行基准函数以计算一个平均的执行时间。

go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的
main包用于调用相应的测试函数，然后构建并运行、报告测试结果，最后清理测试中生成的
临时文件。
*/
package gotest

import (
	"testing" // 每个测试函数必须导入testing包。
)

// 每个测试函数必须导入testing包。
func TestName(t *testing.T) {

}

// 测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头
func TestSin(t *testing.T) { /* ... */ }
func TestCos(t *testing.T) { /* ... */ }
func TestLog(t *testing.T) { /* ... */ }
