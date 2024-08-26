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

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:     "complete",
	Short:   "Complete a task",
	Long:    `Complete a task by inputting the ID`,
	Example: "  gotodo complete [ID]",
	Args:    cobra.ExactArgs(1),
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

		for _, record := range records {
			rec_id, _ := strconv.Atoi(record[0])
			if id == rec_id {
				record[3] = "true"
				flag = true
				break
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

		for _, record := range records {
			line := strings.Join(record, ",")
			line += "\n"
			file.WriteString(line)
		}
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
