/*
使用命令 go get  可以下载一个单一的包或者用 ...  下载整个子目录里面的每个包。

这种包叫internal包，一个internal包只能被和internal目录有同一个父目录的包所导入。例
如，net/http/internal/chunked内部包只能被net/http/httputil或net/http包导入，但是不能被
net/url包导入。不过net/url包却可以导入net/http/httputil包。
*/
package packages

import (
	"strings"
)

func test() {
	strings.Compare("123", "123")
}
