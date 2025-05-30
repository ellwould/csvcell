package main

import (
        "github.com/ellwould/csvcell"
)

func main() {

        dirPath := "/"
        fileName := "example.csv"

        // Data to write to CSV
        produce := "Broccoli"
        produceType := "Vegetable"
        quantity := "9"

        // Variable with all strings concatenated and commas included inbetween
        data := produce + "," + produceType + "," + quantity

        // WriteCSV function called from csvcell package
        // (Directory Path, name of CSV file, commas to prepend, data to write to the file, commas to append)
        csvcell.WriteCSV(dirPath, fileName, 0, data, 0)
}
