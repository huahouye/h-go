package main

import (
	"fmt"
)

// 定义接口
type Speaker interface {
	Say(string)
	Listen(string) string
	Interrupt(string)
}

// WaLan 实现接口 Speaker
type WanLan struct {
	msg string
}

func (w *WanLan) Say(msg string) {
	fmt.Printf("WanLan Say: %s\n", msg)
}

func (w *WanLan) Listen(msg string) string {
	w.msg = msg
	return msg
}

func (w *WanLan) Interrupt(msg string) {
	w.Say(msg)
}

// JiangLou 实现接口 Speaker
type JiangLou struct {
	msg string
}

func (j *JiangLou) Say(msg string) {
	fmt.Printf("JiangLou Say: %s\n", msg)
}

func (j *JiangLou) Listen(msg string) string {
	j.msg = msg
	return msg
}

func (j *JiangLou) Interrupt(msg string) {
	j.Say(msg)
}

func main() {

	wl := &WanLan{}
	jl := &JiangLou{}

	var person Speaker
	person = wl
	person.Say("I am WanLan")
	person = jl
	person.Say("I am JiangLou")

}
