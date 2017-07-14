/*
实际上它就是interface类型，这个类型有一个返回错误信息的单一方法
*/
package errorinterface

import ()

/*
创建一个error最简单的方法就是调用errors.New函数，它会根据传入的错误信息返回一个新
的error。整个errors包仅只有4行：
*/
func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	text string
}

func (e *errorString) Error() string {
	return e.text
}
