// 创建人：lg
// 创建时间：2021/5/17
package main

import "fmt"

func main() {
	// var p People = Student{}   // 编译报错(Student未实现People，因为Speak方式是指针接收器，如果是对象接收器则编译通过)
	var p People = &Student{}  // 可以编译通过，对象接收器也可以编译通过
	s := Student{}
	fmt.Println(s.Speak("hello"))
	fmt.Println(p.Speak("bitch"))
}

type People interface {
	Speak(string) string
}

type Student struct {

}

func (s *Student) Speak(think string) (talk string) {
	talk = "Hi"
	if think == "bitch" {
		talk = "You are a good girl"
	}
	return
}

//  两种实现不可以同时存在，
//  Go 语言的编译器会在结构体类型和指针类型都实现一个方法时报错 —— method redeclared。
// func (s Student) Speak(think string) (talk string) {
// 	talk = "Hi"
// 	if think == "bitch" {
// 		talk = "You are a good girl"
// 	}
// 	return
// }

