package main

import (
    "fmt"
    "rsc.io/quote"
    "net/http"
)

func main() {
    resp, err := http.Get("http://example.com/")
    if err != nil {
        fmt.Println("请求错误")
    }
    fmt.Print(resp)
    fmt.Println(quote.Hello())
}
