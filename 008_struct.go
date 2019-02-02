package main

import "fmt"

type Point2d struct {
    x int32
    y int32
}

// 定义point2d的方法
func (p Point2d) GetX() int32{
    return p.x
}
func (p Point2d) GetY() int32{
    return p.y
}
func (p *Point2d) SetX(new_x int32) {
    p.x = new_x
}
func (p *Point2d) SetY(new_y int32) {
    p.y = new_y
}

func main() {
    // 显式指名所有成员的值
    var p1 Point2d = Point2d {
        x: 32,
        y: 43,
    }
    fmt.Printf("%+v\n", p1)

    //为部分成员显示赋值
    var p2 = Point2d {
        y: 43,
    }
    fmt.Printf("%+v\n", p2)
    //所有成员都不赋值
    var p3 = Point2d {32, 43}
    fmt.Printf("%+v\n", p3)

    // 创建一个结构体的指针变量

    var p4 = &Point2d {32, 43}
    var p5 = new(Point2d)
    fmt.Println(p4, p5)
    fmt.Printf("%+v\n", p4)
    fmt.Printf("%+v\n", p5)

    // 零值结构体
    var a Point2d
    var b = Point2d{}
    var c *Point2d = new(Point2d)
    fmt.Printf("%+v\n", a)
    fmt.Printf("%+v\n", b)
    fmt.Printf("%+v\n", c)

    // nil结构体
    var d *Point2d = nil
    fmt.Printf("%+v\n", d)

    // 包含切片的结构体
    type SliceStruct struct {
        value []int
    }
    var ss SliceStruct = SliceStruct {
        []int{0, 1, 2, 3},
    }
    fmt.Println(ss)

    // 包含数组的结构体
    type ArrayStruct struct {
        value [4]int
    }
    var as ArrayStruct = ArrayStruct {
        [...]int{0, 1, 2, 3},
    }
    fmt.Println(as)

    a.SetX(111)
    a.SetY(222)
    fmt.Println(a.GetX(), a.GetY())

    // 结构体中包含结构体，但是都有名的
    type Circle struct {
        center Point2d
        radius int32
    }
    var circle Circle = Circle {
        Point2d {11, 22},
        5,
    }
    fmt.Println(circle)
    fmt.Println(circle.center.GetX())

    // 包括无名结构体
    type Rectangle struct {
        Point2d // 这个成员是不具名的
        w int32
        h int32
    }
    var rect = Rectangle {
        Point2d {22, 33},
        100,
        100,
    }
    fmt.Println(rect)
    //对于不具名的成员的方法，可以当成自己的方法一样，直接访问
    fmt.Println(rect.x)
}
