package main

import "fmt"

func main() {
    var a [10]int
    var b [10]int = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    c := [10]int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
    // go的数组可以直接打印
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)

    var d [8]int
    // d = c // 编译错误，d和c的类型不同


    //使用range对数组进行遍历
    for index := range  d {
        d[index] = index
    }
    fmt.Println(d)

    for index, value := range d {
        fmt.Println(index, value)
    }

}
