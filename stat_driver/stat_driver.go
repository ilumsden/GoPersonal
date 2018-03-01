package main

import (
    "fmt"
    "stat_driver/stat"
)

func main() {
    avg := stat.Average(2, 5, 3, 1, 9, 10, 6, 8, 4, 15)
    fmt.Printf("The average is %f\n", avg)
    return
}
