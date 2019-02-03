package main

import "fmt"

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}


func sign(a int) int {
    if a > 0 {
        return 1
    } else if a < 0 {
        return -1
    } else {
        return 0
    }
}

// 从0加到n，练习基本的for循环的语法
func sum1_to_n(n int) int {
    sum := 0
    for i := 0; i < n; i++ {
        sum += i
    }
    return sum
}

func test_for_loop() {
    i := 0
    // for后面可以什么都不加，类似c里的while(true)
    for {
        i++
        if i < 8 {
            continue
        } else if i < 10 {
            fmt.Println("Hello, world")
        } else {
            break
        }
    }
}

func main() {
	fmt.Println(sign(max(23, -32)))
    fmt.Println(sum1_to_n(100))
    test_for_loop()
}
