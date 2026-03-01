package main

import "fmt"

func main() {
    conditionValue := 10

    if conditionValue > 5 {
        fmt.Println("OK!")
    } else {
        fmt.Println("Not good!: Doesn't match any condition!")
    } else {
        fmt.Println("The condition value is out of range!")
    }
}
