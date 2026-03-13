package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run gen_data.go <filename> <num_rows>")
		os.Exit(1)
	}

	filename := os.Args[1]
	numRows, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Invalid number of rows: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Could not create file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Header
	writer.Write([]string{"id", "name", "email", "role"})

	for i := 1; i <= numRows; i++ {
		writer.Write([]string{
			strconv.Itoa(i),
			fmt.Sprintf("User %d", i),
			fmt.Sprintf("user%d@example.com", i),
			"admin",
		})
	}

	fmt.Printf("Successfully generated %d rows in %s\n", numRows, filename)
}
