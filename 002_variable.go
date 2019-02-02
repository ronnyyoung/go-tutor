package main

import "fmt"


// 定义一个全局的变量

var global_s int = 233 // 首字母小写的全局变量，只是包内可以访问
var Global_s int = 233 // 公开的全局变量
const const_global_s int = 233

func main() {
    // !定义变量的三种方式

    // 定义变量的完全体
    // 适合强制变量类型的类型，以及全局变量等
    var s int = 42
    // 省略变量类型
    var s1 = 42
    // 定义变量的简化体，适合局部变量
    s2 := 42

    fmt.Println(global_s, Global_s, const_global_s)
    fmt.Println(s, s1, s2)

    // ============================================
    // 定义不同类型的变量
    var a int8 = 1
    var b int16 = 2
    var c int32 = 3
    var d int64 = 4

    fmt.Println(a, b, c, d)
    var ua uint8 = 1
    var ub uint16 = 2
    var uc uint32 = 3
    var ud uint64 = 4
    fmt.Println(ua, ub, uc, ud)

    // int型的字节数取决于是32位系统(4)还是64位系统(8)
    var e int = 5
    var ue uint = 5
    fmt.Println(e, ue)

    var f bool = true
    fmt.Println(f)

    // 相当于c语言里的char
    var g byte = 'a'
    var h string = "hello,world"
    fmt.Println(g, h)

    var i float32 = 3.14
    var j float64 = 3.1415926
    fmt.Println(i, j)

    // 定义指针
    var k int = 12
    var p *int = &k
    fmt.Println(k, p)
}
