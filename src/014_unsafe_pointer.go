package main

import (
    "fmt"
    "unsafe"
)

func main() {
    type Rect struct {
        w int
        h int
    }
    // 通过指针来操作struct对象
    var r Rect = Rect{1, 2}
    fmt.Println(&r)
    fmt.Println(unsafe.Pointer(&r))

    // 只用unsafe.Pointer可以与各种指针类型之间进行转换
    var pw *int = (*int)(unsafe.Pointer(&r))
    // 可以用unsafe.Offsetof来取得结构体成员在结构体中的偏移量
    var ph *int = (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&r)) + unsafe.Offsetof(r.h)))
    *pw = 11
    *ph = 22
    fmt.Println(r)
}
