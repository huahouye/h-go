/*
这个模板先打印匹配到的issue总数，然后打印每个issue的编号、创建用户、标题还有存在的
时间。对于每一个action，都有一个当前值的概念，对应点操作符，写作“.”。当前值“.”最初被
初始化为调用模板是的参数，在当前例子中对应github.IssuesSearchResult类型的变量。模板
中 {{.TotalCount}}  对应action将展开为结构体中TotalCount成员以默认的方式打印的值。模
板中 {{range .Items}}  和 {{end}}  对应一个循环action，因此它们直接的内容可能会被展开多
次，循环每次迭代的当前值对应当前的Items元素的值。
在一个action中， |  操作符表示将前一个表达式的结果作为后一个函数的输入，类似于UNIX
中管道的概念。在Title这一行的action中，第二个操作是一个printf函数，是一个基于
fmt.Sprintf实现的内置函数，所有模板都可以直接使用。对于Age部分，第二个动作是一个叫
daysAgo的函数，通过time.Since函数将CreatedAt成员转换为过去的时间长度：
*/
package templates

import (
	"time"
)

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}
