/*
JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。在类似的协
议中，JSON并不是唯一的一个标准协议。 XML（§7.14）、ASN.1和Google的Protocol
Buffers都是类似的协议，并且有各自的特色，但是由于简洁性、可读性和流行程度等原因，
JSON是应用最广泛的一个。
Go语言对于这些标准格式的编码和解码都有良好的支持，由标准库中的encoding/json、
encoding/xml、encoding/asn1等包提供支持（译注：Protocol Buffers的支持由
github.com/golang/protobuf 包提供），并且这类包都有着相似的API接口。本节，我们将对
重要的encoding/json包的用法做个概述。
JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象——Unicode本文编
码。它可以用有效可读的方式表示第三章的基础数据类型和本章的数组、slice、结构体和
map等聚合数据类型。
基本的JSON类型有数字（十进制或科学记数法）、布尔值（true或false）、字符串，其中字
符串是以双引号包含的Unicode字符序列，支持和Go语言类似的反斜杠转义特性，不过JSON
使用的是\Uhhhh转义数字来表示一个UTF-16编码（译注：UTF-16和UTF-8一样是一种变长的
编码，有些Unicode码点较大的字符需要用4个字节表示；而且UTF-16还有大端和小端的问
题），而不是Go语言的rune类型。
这些基础类型可以通过JSON的数组和对象类型进行递归组合。一个JSON数组是一个有序的
值序列，写在一个方括号中并以逗号分隔；一个JSON数组可以用于编码Go语言的数组和
slice。一个JSON对象是一个字符串到值的映射，写成以系列的name:value对形式，用花括号
包含并以逗号分隔；JSON的对象类型可以用于编码Go语言的map类型（key类型是字符串）
和结构体。例如
boolean true
number -273.15
string "She said \"Hello, BF\""
array ["gold", "silver", "bronze"]
object {"year": 1980,
"event": "archery",
"medals": ["gold", "silver", "bronze"]}
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	main1()
	main2()
}

// marshal，go 对象 -> json 字符串
func main1() {
	fmt.Println("main1()")
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}
	fmt.Printf("%#v\n", movies)

	// 这样的数据结构特别适合JSON格式，并且在两种之间相互转换也很容易。将一个Go语言中
	// 类似movies的结构体slice转为JSON的过程叫编组（marshaling）。编组通过调用
	// json.Marshal函数完成
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	// Marshal函数返还一个编码后的字节slice，包含很长的字符串，并且没有空白缩进；我们将它
	// 折行以便于显示
	fmt.Printf("%s\n", data)
	// 为了生成便于阅读的格式，另
	// 一个json.MarshalIndent函数将产生整齐缩进的输出
	data2, err2 := json.MarshalIndent(movies, "", " ")
	if err2 != nil {
		log.Fatalf("JSON marshaling failed: %s", err2)
	}
	fmt.Printf("%s\n", data2)

	/*
		只有导出的结构体成员才会被编码，这也就是我们为什么选择用大写字
		母开头的成员名称。其中Year名字的成员在编码后变成了released，还有Color成员
		编码后变成了小写字母开头的color。这是因为构体成员Tag所导致的。一个构体成员Tag是和
		在编译阶段关联到该成员的元信息字符串
		Year int `json:"released"`
		Color bool `json:"color,omitempty"`
		结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值
		对序列；因为值中含义双引号字符，因此成员Tag一般用原生字符串面值的形式书写。json开
		头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的
		包也遵循这个约定。成员Tag中json对应值的第一部分用于指定JSON对象的名字，
		语言中的TotalCount成员对应到JSON中的total_count对象。Color成员的Tag还带了一个额外
		的omitempty选项，表示当Go语言结构体成员为空或零值时不生成JSON对象（这里false为零
		值）。果然，Casablanca是一个黑白电影，并没有输出Color成员。
	*/
}

// unmarshal，json 字符串 -> go 对象 slice
func main2() {
	fmt.Println("main2()")
	// 当Unmarshal函数调用返回，slice将被只含有Title信息值填
	// 充，其它JSON成员将被忽略。
	data := []byte(`[{"Title":"Casablanca","released":1942,"Actors":["Humphrey Bogart","Ingrid Bergman"]},{"Title":"Cool Hand Luke","released":1967,"color":true,"Actors":["Paul Newman"]},{"Title":"Bullitt","released":1968,"color":true,"Actors":["Steve McQueen","Jacqueline Bisset"]}]`)
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles) // "[{Casablanca} {Cool Hand Luke} {Bullitt}]"
}
