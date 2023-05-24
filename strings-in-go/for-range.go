package main

import "fmt"

func main() {
	const c = "日本語"
	for i, r := range c {
        fmt.Printf("%U starts at byte position %d\n", r, i)
    }
}