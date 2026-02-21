package csvcell

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Function to read a CSV file
func ReadCSV(rootDirPath string, fileName string) (result [][18278]string) {

	// Go introduced OpenRoot in version 1.24, it restricts file operations to a single directory
	rootDir, err := os.OpenRoot(rootDirPath)
	if err != nil {
		panic("Directory path does not exist")
	}

	defer rootDir.Close()

	// Open the file
	file, err := rootDir.Open(fileName)
	if err != nil {
		log.Fatal("File " + fileName + " cannot be opened or does not exist")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	// For loop to scan lines in the file and add to a slice
	for scanner.Scan() {
		line := scanner.Text()
		// Create a slice and split each line into an index based on the , delimiter
		var splitSlice = []string{}
		splitSlice = strings.Split(line, ",")
		// Convert splitSlice into an array named splitArray
		splitArray := [18278]string{}
		copy(splitArray[:], splitSlice)
		// Append to array
		result = append(result, splitArray)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

// Function to write to a CSV file
func WriteCSV(rootDirPath string, fileName string, commaPrepend int, csvRow string, commaAppend int) {

	// Go introduced OpenRoot in version 1.24, it restricts file operations to a single directory
	rootDir, err := os.OpenRoot(rootDirPath)
	if err != nil {
		panic("Directory path does not exist")
	}

	defer rootDir.Close()

	// Open/create the file
	file, err := rootDir.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File " + fileName + " cannot be opened or does not exist")
	}

	defer file.Close()

	// Write the string to the file and add optional commas at the beginning or end of the row
	_, err = file.WriteString(strings.Repeat(",", commaPrepend) + csvRow + strings.Repeat(",", commaAppend) + "\n")
	if err != nil {
		panic(err)
	}
}

// Function to retrieve all contents from a plain text file
func FileData(rootDirPath string, fileName string) (result string) {

	// Go introduced OpenRoot in version 1.24, it restricts file operations to a single directory
	rootDir, err := os.OpenRoot(rootDirPath)

	if err != nil {
		panic("Directory path does not exist")
	}

	defer rootDir.Close()

	file, err := rootDir.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var fileSlice []string

	for scanner.Scan() {
		line := scanner.Text()
		fileSlice = append(fileSlice, line)
	}

	result = strings.Join(fileSlice, "")
	return result
}

// Function to retrieve all file and directory names from a specific directory
func DirList(rootDirPath string) (result []string) {
        rootDir, err := os.Open(rootDirPath)

        if err != nil {
                panic("Directory path does not exist")
        }

        defer rootDir.Close()

        fileList, err := rootDir.ReadDir(-1)

        if err != nil {
                log.Fatal(err)
        }

        for _, fileListLoop := range fileList {
                result = append(result, fileListLoop.Name())
        }
        return result
}
