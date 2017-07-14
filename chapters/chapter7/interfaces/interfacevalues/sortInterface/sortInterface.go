/*
前面我们已经把接口值的技巧都讲完了，让我们来看更多的一些在Go标准库中的重要接口类
型。在下面的三章中，我们会看到接口类型是怎样用在排序，web服务，错误处理中的。

Go语言的sort.Sort函数不会对具体的序列和它的元素做任何假设。相反，它
使用了一个接口类型sort.Interface来指定通用的排序算法和可能被排序到的序列类型之间的约
定。这个接口的实现由序列的具体表示和它希望排序的元素决定，序列的表示经常是一个切
片。
*/
package sortInterface

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

/*
一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换
两个元素的方式；这就是sort.Interface的三个方法：
*/
type Interface interface {
	Len() int
	Less(i, j int) bool // i, j are indices of sequence elements
	Swap(i, j int)
}

/*
为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个
实例应用sort.Sort函数。思考对一个字符串切片进行排序，这可能是最简单的例子了。下面是
这个新的类型StringSlice和它的Len，Less和Swap方法
*/
type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// 现在我们可以通过像下面这样将一个切片转换为一个StringSlice类型来进行排序
func main1() {
	names := []string{"1", "2", "3"}
	sort.Sort(StringSlice(names))
	// 对字符串切片的排序是很常用的需要，所以sort包提供了StringSlice类型，也提供了Strings函
	// 数能让上面这些调用简化成sort.Strings(names)
	sort.Strings(names)
}

//////////////////////

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

/*
printTracks函数将播放列表打印成一个表格。一个图形化的展示可能会更好点，但是这个小程
序使用text/tabwriter包来生成一个列是整齐对齐和隔开的表格，像下面展示的这样。注意到
*tabwriter.Writer是满足io.Writer接口的。它会收集每一片写向它的数据；它的Flush方法会格
式化整个表格并且将它写向os.Stdout（标准输出）。
*/
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

/*
为了能按照Artist字段对播放列表进行排序，我们会像对StringSlice那样定义一个新的带有必
须Len，Less和Swap方法的切片类型
*/
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
func main2() {
	sort.Sort(byArtist(tracks))
	// 对tracks进行逆向排序
	sort.Sort(sort.Reverse(byArtist(tracks)))
}

/*
为了使用方便，sort包为[]int,[]string和[]float64的正常排序提供了特定版本的函数和类型。对
于其他类型，例如[]int64或者[]uint，尽管路径也很简单，还是依赖我们自己实现。
*/
