package main

import "fmt"

func main() {
    //go中字符串使用utf8编码，汉字通常占3个字节
    var s = "Hello,世界"
    fmt.Println(s, len(s))

    for i := 1; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Printf("\n")

    // rune_value是字符对应的unicode编码
    for codepoint, rune_value := range s {
        fmt.Printf("%d: %d, ", codepoint, rune_value)
    }
    fmt.Printf("\n")

    // 字符串可以像slice一样进行下标范围切片
    var substr = s[0:11]
    fmt.Println(substr)

    // 将字符串转换为字符切片
    var a = []byte(s)
    fmt.Println(a)
}
