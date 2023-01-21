package main

import (
    "fmt"

    "github.com/BurntSushi/toml"
)

type Config struct {
    Title  string
    Author struct {
        Name  string
        Email string
    }
}

func main() {
    var config Config
    _, err := toml.DecodeFile("config.toml", &config)
    if err != nil {
        panic(err)
    }
    fmt.Println("Title:", config.Title)
    fmt.Println("Author:", config.Author.Name)
    fmt.Println("Email:", config.Author.Email)
}
