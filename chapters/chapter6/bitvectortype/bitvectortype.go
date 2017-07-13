/*
Go语言里的集合一般会用map[T]bool这种形式来表示，T代表元素类型。集合用map类型来表
示虽然非常灵活，但我们可以以一种更好的形式来表示它。例如在数据流分析领域，集合元
素通常是一个非负整数，集合会包含很多元素，并且集合会经常进行并集、交集操作，这种
情况下，bit数组会比map表现更加理想。(译注：这里再补充一个例子，比如我们执行一个
http下载任务，把文件按照16kb一块划分为很多块，需要有一个全局变量来标识哪些块下载完
成了，这种时候也需要用到bit数组)
一个bit数组通常会用一个无符号数或者称之为“字”的slice或者来表示，每一个元素的每一位都
表示集合里的一个值。当集合的第i位被设置时，我们才说这个集合包含元素i。下面的这个程
序展示了一个简单的bit数组类型，并且实现了三个函数来对这个bit数组来进行操作
*/

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
package bitvectortype

import ()

type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte('}')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main1() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	x.UnionWith(&y)
	fmt.Println(x.String())           // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
}
