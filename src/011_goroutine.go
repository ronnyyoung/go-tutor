package main

import (
    "fmt"
    "time"
    "runtime"
)

// 测试主协程退出，将导入子协程提前退出
func test1() {
    // 启动一个协程，先Sleep 1秒
    go func() {
        time.Sleep(time.Second)
        fmt.Println("go routine")
    }()
    fmt.Println("exec main") //这句话可以和子协和同时执行
    // 防止主协程提前退出
    time.Sleep(time.Second)
    time.Sleep(time.Second)
    fmt.Println("exit main")
}

// 协程的调试机制
func test2() {
    n := 7 // 在8核的机器上,超过8时,主协程就没有机会执行了
    // 启动n个死循环的协程
    for i := 0; i < n; i++ {
        go func() {
            fmt.Println("exec sub routine")
            // 死循环
            for {}
        }()
    }
    for {
        time.Sleep(time.Second)
        fmt.Println("exec main")
    }
}

// 重新设置可用线程数来观察协程的调度机制
func test3() {
    fmt.Println(runtime.GOMAXPROCS(0))
    runtime.GOMAXPROCS(4)
    fmt.Println(runtime.GOMAXPROCS(0))
    n := 7 // 现在最大线程变为了4，所以主协程就没有机会再执行了
    // 启动n个死循环的协程
    for i := 0; i < n; i++ {
        go func() {
            fmt.Println("exec sub routine")
            // 死循环
            for {}
        }()
    }
    // 如果不加sleep，主协程是可以直接抢占到cpu的,然后导致上面只有3个子协程可以执行
    // 猜测是因为子协作的启动需要时间
    //time.Sleep(time.Second)
    fmt.Println("exec main")
    for {}
}

func main() {
    test1()
    test3()
}
