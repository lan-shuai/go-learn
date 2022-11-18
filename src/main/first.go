package main // 第一行代码 package main 定义了包名。你必须在源文件中非注释的第一行指明这个文件属于哪个包，package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。

import "fmt"

// aa := "hello"

func main(){
	a := "hello"
	fmt.Println("Hello, go")
	fmt.Println(a)
}