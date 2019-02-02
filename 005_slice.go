package main

import "fmt"

func main() {
    //　创建一个cap为0, len为5的切片
    var a []int = make([]int, 5, 9)
    // 简化版本的创建一个len为100,cap=100的切片
    b := make([]int, 20)
    fmt.Println(a)
    fmt.Println(b)

    // 满容的切片
    var c []int = []int{1, 2, 3, 4, 5}
    fmt.Println(len(c), cap(c))

    // 创建空的切片
    var d []int
    var e = []int{}
    f := make([]int, 0)
    fmt.Println(len(d), cap(d))
    fmt.Println(len(e), cap(e))
    fmt.Println(len(f), cap(f))

    for index := range a {
        a[index] = index
    }
    fmt.Println(a)
    g := a //浅拷贝, g和a共享底层存储
    g[3] = 255
    fmt.Println(g)
    fmt.Println(a)

    // append的操作
    h := make([]int, 5)
    h = append(h, 42) // 触发h的扩容
    fmt.Println(len(h), cap(h))

    i := make([]int, 5, 6)
    _ = append(i, 42)
    fmt.Println(len(i), cap(i))
    i = append(i, 42)
    fmt.Println(len(i), cap(i))

    //切片
    var j = a[2:4] // j与a底层共享存储
    //j的cap要小于a
    fmt.Println(len(j), cap(j))
    var k = a[3:]
    var l = a[:3]
    var m = a[:] // 该句等价于m=a，底层也是共享存储的
    fmt.Println(a, k, l, m)

    // copy操作
    n := make([]int, 3, 20)
    u := copy(n, a) //copy的数据量是min(len(n), len(a))
    fmt.Println(u, n)

    //nil切片和空切片
    var p []int
    var q = *new([]int)
    var r = []int{}
    var t = make([]int, 0)

    fmt.Println(p == nil)
    fmt.Println(q == nil)
    fmt.Println(r == nil)
    fmt.Println(t == nil)
}
