package main

import (
        "fmt"
        "github.com/ellwould/csvcell"
        "strings"
)

func main() {
        dirPath := "/"
        fileName := "example.csv"

        // Reading a CSV file where each row has 3 columns
        read := csvcell.ReadCSV(dirPath, fileName)
        for _, read := range read {
                column1 := strings.Join((read[0:][0:1]), ", ")
                column2 := strings.Join((read[0:][1:2]), ", ")
                column3 := strings.Join((read[0:][2:3]), ", ")
                fmt.Println(column1, column2, column3)
        }
}
