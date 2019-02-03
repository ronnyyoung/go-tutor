package main

import "fmt"
import "time"
import "sync"

func task(wc *sync.WaitGroup) {
    defer wc.Done()
    fmt.Println("do task")
    time.Sleep(2 * time.Second)
}

func producer(ch1 chan int, ch2 chan int) {
    for i := 0; i < 100; i++ {
        select {
            case ch1 <- i:
                fmt.Printf("produce: %d\n", i)
            case ch2 <- i + 100:
                fmt.Printf("produce: %d\n", i + 100)
        }
    }
}

func comsumer(ch1 chan int, ch2 chan int) {
    for {
        select {
        case i := <- ch1 :
            fmt.Printf("comsume: %d\n", i)
        case i := <- ch2:
            fmt.Printf("comsume: %d\n", i)
        }
    }
}


func send_data_with_drop(c chan int) {
    for i := 0 ; i < 100; i++ {
        select {
        case c <- i:
            fmt.Printf("send data: %d\n", i)
        default:
            fmt.Printf("drop the data: %d\n", i)
        }
    }
}

func recv_data(c chan int) {
    for data := range c {
        fmt.Printf("recv data: %d\n", data)
    }
}

func main() {
    // 创建通道
    // 创建一个非缓冲型的通道
    var c1 chan int = make(chan int)
    fmt.Println(len(c1), cap(c1))
    // 创建一个大小为1024的缓冲型的通道
    var c2 chan int = make(chan int, 1024)
    fmt.Println(len(c2), cap(c2))

    // 往通道里写数据
    for i := 0; i < 5; i++ {
        c2 <- i + 1
    }
    // 直接从通道中读数据
    for len(c2) > 0 {
        data := <- c2
        fmt.Println(data)
    }

    for i := 0; i < 5; i++ {
        c2 <- i + 5
    }
    close(c2)
    // c2 <- 42 // 往已经关闭的通道内写数据，会抛异常
    // 使用for-range从通道里读数据
    for data := range c2 {
        fmt.Println(data)
    }
    time.Sleep(time.Second)

    // 使用WaitGroph来回收子协程
    var wc  = new(sync.WaitGroup)
    wc.Add(1) // 设置需要回收子协程的个数
    go task(wc)
    wc.Wait()

    // 用Select来处理多个chan的多路复用
    // 只要有一个chan可以写入就不会阻塞producer
    // 只要有一个chan可以读取就不会阻塞comsumer
    var chn1 chan int = make(chan int)
    var chn2 chan int = make(chan int)
    go producer(chn1, chn2)
    go comsumer(chn1, chn2)
    time.Sleep(time.Second * 2) // 让上面的协程跑完

    // 使用select的default分支来实现非阻塞
    var chn chan int = make(chan int, 10)
    go send_data_with_drop(chn)
    go recv_data(chn)
    time.Sleep(time.Second * 2) // 让上面的协程跑完
}
