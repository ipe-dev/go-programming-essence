package main

import "fmt"

type F struct {
	Name string
	Age  int
}

// Stringerインターフェースを実装する
// type Stringer interface {
// 	String() string
// }
func (f *F) String() string {
	return fmt.Sprintf("NAME=%q,AGE=%d", f.Name, f.Age)
}

func main() {
	f := &F{
		Name: "John",
		Age:  20,
	}
	// fmtパッケージに渡したときのフォーマットが変わる
	fmt.Printf("%v\n", f)
}
