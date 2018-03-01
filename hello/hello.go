package main

import "fmt"

func main() {
    name := "Ian Lumsden"
    time := [2]int{9, 54}
    fmt.Printf("Hello World!\nMy name is %s.\nThe time is %d:%d\n", name,
               time[0], time[1])
    return
}
