/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the list",
	Long:  `Delete a task by inputting the ID`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Invalid ID entered. Enter a numerical value")
			return
		}

		file, err := LoadFile(Filename)
		if err != nil {
			fmt.Printf("Error loading file: %s\n", err)
		}

		csvReader := csv.NewReader(file)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Printf("Error reading csv file: %s\n", err)
		}

		flag := false

		var new [][]string
		for _, record := range records {
			rec_id, _ := strconv.Atoi(record[0])
			if id != rec_id {
				new = append(new, record)
			} else {
				flag = true
			}
		}

		if !flag {
			fmt.Println("Record not found. Check your ID")
			return
		}

		CloseFile(file)
		os.Remove(Filename)

		file, err = LoadFile(Filename)
		if err != nil {
			fmt.Printf("Error loading file: %s\n", err)
		}
		defer CloseFile(file)

		for _, record := range new {
			line := strings.Join(record, ",")
			line += "\n"
			file.WriteString(line)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
