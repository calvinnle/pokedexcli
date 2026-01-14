package main

import "fmt"

func commandClear(_ *config, _ string) error {
    fmt.Println("all clear sir")
    return nil
}
