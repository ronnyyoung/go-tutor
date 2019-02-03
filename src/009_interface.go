package main

import "fmt"

// 定义二元操作算子的接口
type  BinaryOperator interface {
    run() int32
}

// 定义加法算子
type Add struct {
    a int32
    b int32
}
// 二元运算接口实现
func (ad Add) run() int32 {
    return ad.a + ad.b ;
}

// 定义减法算法
type Sub struct {
    a int32
    b int32
}
// 二元运算接口实现
func (su Sub) run() int32 {
    return su.a - su.b;
}

// 定义二元算子的计算流
type Flow struct {
    operators []BinaryOperator
}
func (f Flow) run() {
    for _, op := range f.operators {
        fmt.Println(op.run())
    }
}


func main() {
    // 空接口使用
    var a = make([]int, 3)
    fmt.Println(a)
    var b interface{}
    b = a
    var c = b.([]int)
    c[0]= 32
    fmt.Println(c)
    fmt.Println(a)

    var add = Add{1, 2}
    var sub = Sub{1, 2}

    var flow Flow = Flow{
        make([]BinaryOperator, 0, 10),
    }
    flow.operators = append(flow.operators, add)
    flow.operators = append(flow.operators, sub)
    flow.run()
}
