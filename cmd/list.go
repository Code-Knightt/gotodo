/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long:  `List all your pending tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := cmd.Flags().GetBool("incomplete")
		if err != nil {
			fmt.Println("Error parsing flag. Please try again")
			return
		}

		file, err := LoadFile(Filename)
		if err != nil {
			fmt.Printf("Error loading file: %s\n", err)
		}
		defer CloseFile(file)

		csvReader := csv.NewReader(file)
		records, err := csvReader.ReadAll()
		if err != nil {
			fmt.Printf("Error reading csv file: %s\n", err)
		}

		if len(records) == 0 {
			fmt.Printf("No tasks here")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 10, 0, 2, ' ', 0)
		defer w.Flush()
		fmt.Fprintln(w, "ID\tTask\tCreated At\tCompleted?")

		for _, record := range records {
			id, err := strconv.Atoi(record[0])
			if err != nil {
				id = 1
			}

			recTime, err := time.Parse(time.RFC3339, record[2])
			if err != nil {
				recTime = time.Now()
			}

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\n", id, record[1], timediff.TimeDiff(recTime), record[3])
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolP("incomplete", "i", false, "Show only incomplete tasks")
}
