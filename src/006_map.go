package main

import "fmt"

func main() {
    //　创建map
    var a map[string]string = map[string]string{}
    // 插入元素
    a["name"] = "yang yansheng"
    a["email"] = "yangyansheng@example.com"
    a["age"] = "26"

    var b = map[string]string{
        "name": "yang yansheng",
        "email": "yangyansheng@example.com",
        "age": "26",
    }
    fmt.Println(a)
    fmt.Println(b)

    // 使用make创建
    var c = make(map[string]int)
    // 提前申请容量，避免动态扩容
    var d = make(map[string]int, 16)
    fmt.Println(c, len(c))
    fmt.Println(c, len(d))

    // 删除元素
    delete(a, "email")
    fmt.Println(a)

    // 遍历map
    for key,value := range b {
        fmt.Println(key, ":", value)
    }

    // 判断key值是否存在
    name, ok := a["no_exist_key"]
    if ok {
        fmt.Println(name)
    } else {
       fmt.Println("key值不存在") 
    }
}
