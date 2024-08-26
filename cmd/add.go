package cmd

import (
	"encoding/csv"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a task",
	Long:    `Create a new task`,
	Example: "  gotodo add [Task_Name]",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]

		if len(strings.TrimSpace(name)) == 0 {
			fmt.Println("Name is requred")
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

		var id int
		if len(records) == 0 {
			id = 1
		} else {
			id, _ = strconv.Atoi(records[len(records)-1][0])
			id += 1
		}
		file.WriteString(fmt.Sprintf("%d,%s,%s,false\n", id, name, time.Now().Format(time.RFC3339)))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
