/*
接口类型具体描述了一系列方法的集合，一个实现了这些方法的具体类型是这个接口类型的实例。

io.Writer类型是用的最广泛的接口之一，因为它提供了所有的类型写入bytes的抽象，包括文
件类型，内存缓冲区，网络链接，HTTP客户端，压缩工具，哈希等等。io包中定义了很多其
它有用的接口类型。Reader可以代表任意可以读取bytes的类型，Closer可以是任意可以关闭
的值，例如一个文件或是网络链接。（到现在你可能注意到了很多Go语言中单方法接口的命
名习惯）
*/

package interfacetypes

import ()

type Reader interface {
	Read(p []byte) (n int, err error)
}
type Closer interface {
	Close() error
}
type Writer interface {
	Write(p []byte) (n int, err error)
}
type ReadWriter interface {
	Reader
	Writer
}
type ReadWriterCloser interface {
	Reader
	Writer
	Closer
}

/*
上面用到的语法和结构内嵌相似，我们可以用这种方式以一个简写命名另一个接口，而不用
声明它所有的方法。这种方式本称为接口内嵌。尽管略失简洁，我们可以像下面这样，不使
用内嵌来声明io.Writer接口。
*/
type ReadWriter2 interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}

// 或者甚至使用种混合的风格：
type ReadWriter interface {
	Read(p []byte) (n int, err error)
	Writer
}

/*
上面3种定义方式都是一样的效果。方法的顺序变化也没有影响，唯一重要的就是这个集合里
面的方法。
*/
