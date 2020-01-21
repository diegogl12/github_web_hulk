package services

import (
    "os"
    "log"
    "encoding/csv"
)

var data = [][]string{{"Repository", "Hello Readers of"}, {"TimeStamp", "golangcode.com"}}

func CreateAndWriteCsv() {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, value := range data {
        err := writer.Write(value)
        checkError("Cannot write to file", err)
    }
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
