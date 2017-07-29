// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import (
	"io/ioutil"
	"net/http"
	"time"
)

// A Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	cache map[string]result
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// NOTE: not concurrency-safe!
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

// 这个函数的调用本身开销是比较大的，所以我们尽量尽量避免在不必要的时候反复调用
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
	/*
		ReadAll会返回两个结果，一个[]byte数组和一个错误，不过
		这两个对象可以被赋值给httpGetBody的返回声明里的interface{}和error类型，所以我们也就
		可以这样返回结果并且不需要额外的工作了。
	*/
}

/*
下面是一个使用Memo的例子。对于流入的URL的每一个元素我们都会调用Get，并打印调用
延时以及其返回的数据大小的log
*/
