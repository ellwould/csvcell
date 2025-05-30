# csvcell
#### CSV Cell is a Go package with functions to read from a CSV file and write to a CSV file

<br>

>[!IMPORTANT]
>The ReadCSV function returnes each row within the CSV file as an array
>
>Example of a [Go Program](https://github.com/ellwould/csvcell/blob/main/example/readexample.go) using the ReadCSV function:
>```go
>package main
>
>import (
>        "fmt"
>        "github.com/ellwould/csvcell"
>        "strings"
>)
>
>func main() {
>
>        dirPath := "/"
>        fileName := "example.csv"
>
>        // Reading a CSV file where each row has 3 columns
>        read := csvcell.ReadCSV(dirPath, fileName)
>        for _, read := range read {
>                column1 := strings.Join((read[0:][0:1]), ", ")
>                column2 := strings.Join((read[0:][1:2]), ", ")
>                column3 := strings.Join((read[0:][2:3]), ", ")
>                fmt.Println(column1, column2, column3)
>        }
>}
>```
>
><br><b>Output in CLI:</b>
>```bash
>Name Type Quantity
>Apple Fruit 1
>Pineapple Fruit 2
>Pear Fruit 6
>Orange Fruit 5
>Grape Fruit 3
>Carrot Vegetable 2
>Lemon Fruit 8
>Cherry Fruit 55
>Banana Fruit 34
>Broccoli Vegetable 9
>```
>
><br>The [example.csv](https://github.com/ellwould/csvcell/blob/main/example/example.csv) file used:
>
>| Name	| Type | Quantity |
>|------|------|----------|
>|Apple|Fruit|1|
>|Pineapple|Fruit|2|
>|Pear|Fruit|6|
>|Orange|Fruit|5|
>|Grape|Fruit|3|
>|Carrot|Vegetable|2|
>|Lemon|Fruit|8|
>|Cherry|Fruit|55|
>|Banana|Fruit|34|
>|Broccoli|Vegetable|9|
>

<br>

>[!IMPORTANT]
>Example of a [Go Program](https://github.com/ellwould/csvcell/blob/main/example/writeexample.go) using the WriteCSV function:
>```go
>package main
>
>import (
>        "github.com/ellwould/csvcell"
>)
>
>func main() {
>
>        dirPath := "/"
>        fileName := "example.csv"
>
>        // Data to write to CSV
>        produce := "Broccoli"
>        produceType := "Vegetable"
>        quantity := "9"
>
>        // Variable with all strings concatenated and commas included inbetween
>        data := produce + "," + produceType + "," + quantity
>
>        // WriteCSV function called from csvcell package
>        // (Directory Path, name of CSV file, commas to prepend, data to write to the file, commas to append)
>        csvcell.WriteCSV(dirPath, fileName, 0, data, 0)
>}
>```
>

<br>

>[!NOTE]
>For a list of abbreviations and there meanings used throughout this repository please refer to this [README](https://github.com/Ellwould/information_technology_and_telecommunication_abbreviations)

