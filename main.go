package main

import (
	"fmt"
	"time"
)

func main() {
    gk := time.Now().Format("2006.02.01")
    ck := time.Now().Add(time.Hour * 2).Format("2006.02.01");

    if gk < ck {
        fmt.Println("ok")
    } else {
        fmt.Println("Not OK")
    }
}