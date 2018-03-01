package main

import (
    "io"
    "fmt"
    "github/container"
    "github/sort/sortalg"
)

func main() {
    var v container.Vector
    var data []int
    var tmp int
    for {
        _, err := fmt.Scanf("%d", &tmp)
        if err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println(err)
            return
        } else if tmp == ' ' {
            continue
        } else {
            data = append(data, tmp)
        }
    }
    v = container.NewVector(data...)
    sortalg.Quicksort(&v, 0, v.Length()-1)
    data_out := v.Data()
    for _, d := range data_out {
        fmt.Printf("%d ", d)
    }
    fmt.Printf("\n")
}
