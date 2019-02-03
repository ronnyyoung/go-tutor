package main

import "fmt"
import "errors"
import "os"


func fileio() {
    // 打开文件，返回文件句柄和错误信息
    var f, err = os.Open("010_error_exception.go");
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    // 将文件的关闭延时到作用域结束，下面的代码如果中间抛异常,Close也会执行
    defer f.Close()

    var content []byte = []byte{}
    // buff是一次读取的缓存
    var buff = make([]byte, 100)
    for {
        n, err := f.Read(buff)
        if n > 0 {
            // 注意这里将切片append到切片的用法
            content = append(content, buff[:n]...)
        }
        if err != nil {
            break
        }
    }
    fmt.Println(string(content))
}

func fileio_exception_panic() {
    // 打开文件，返回文件句柄和错误信息
    var f, err = os.Open("010_error_exception1.go");
    defer f.Close()
    if err != nil {
        panic(err)
    }
}

func fileio_exception_recover() {
    // 打开文件，返回文件句柄和错误信息
    var f, err = os.Open("010_error_exception1.go");
    defer f.Close()
    // 定义处理异常的无名函数，并defer到作用域结束时调用
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()

    if err != nil {
        panic(err)
    }
}

func main() {

    // error实际是个interface的类型，定义了Error() string函数
    //使用errors package中的New函数来返回一个error对象
    var err error = errors.New("error msg.")
    fmt.Println(err)

    // 使用fmt.Errorf来灵活定制化错误msg
    var err_msg error = fmt.Errorf("%s", "formated error msg.")
    fmt.Println(err_msg)

    fileio()
    // fileio_exception_panic()
    fileio_exception_recover()
}
