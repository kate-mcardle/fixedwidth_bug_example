package main

import (
    "fmt"
    "io"
    "log"
    "os"
	fixedwidth "github.com/ianlopshire/go-fixedwidth"	
)

type person struct {
    ID        int     `fixed:"1,5"`
    FirstName string  `fixed:"6,15"`
    LastName  string  `fixed:"16,25"`
    Grade     float64 `fixed:"26,30"`
    Age       int    `fixed:"31,33"`
    Alive     string    `fixed:"34,39"`
    Github    string    `fixed:"40,41"`
}

func main() {
    r, err := os.Open("test_file.txt")
    if err != nil {
        fmt.Printf("Error opening file")
        return
    }

    decoder := fixedwidth.NewDecoder(r)
    for {
        var element person
        err := decoder.Decode(&element)
        if err == io.EOF {
            fmt.Printf("Got EOF")
            break
        }
        if err != nil {
            log.Fatal(err)
        }
        fmt.Printf("%+v\n", element)
    }

}