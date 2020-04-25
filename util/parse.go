package util

import (
	"fmt"
	"strings"
)

func Parse(data []map[string]interface{}) [][]string {
	var rows [][]string
	rows = append(rows, []string{"Id", "Date",
		"No1", "No2", "No3", "No4", "No5", "No6", "SNo",
		"SB Code", "SB Name (English)", "SB Name (Chinese)",
		"Investment", "P1", "P1 Unit", "P2", "P2 Unit",
		"P3", "P3 Unit", "P4", "P4 Unit", "P5", "P5 Unit",
		"P6", "P6 Unit", "P7", "P7 Unit"})

	for key, entry := range data {
		var row []string

		fmt.Println("Reading Value for Key :", key)

		row = append(row, entry["id"].(string))
		row = append(row, entry["date"].(string))

		no := strings.Split(entry["no"].(string), "+")
		row = append(row, no...)

		row = append(row, entry["sno"].(string))
		row = append(row, entry["sbcode"].(string))
		row = append(row, entry["sbnameE"].(string))
		row = append(row, entry["sbnameC"].(string))
		row = append(row, entry["inv"].(string))
		row = append(row, entry["p1"].(string))
		row = append(row, entry["p1u"].(string))
		row = append(row, entry["p2"].(string))
		row = append(row, entry["p2u"].(string))
		row = append(row, entry["p3"].(string))
		row = append(row, entry["p3u"].(string))
		row = append(row, entry["p4"].(string))
		row = append(row, entry["p4u"].(string))
		row = append(row, entry["p5"].(string))
		row = append(row, entry["p5u"].(string))
		row = append(row, entry["p6"].(string))
		row = append(row, entry["p6u"].(string))
		row = append(row, entry["p7"].(string))
		row = append(row, entry["p7u"].(string))

		printSlice(row)
		rows = append(rows, row)
	}
	return rows
}

func printSlice(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
