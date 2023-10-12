package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	files := []string{
		"/Users/kyl/Downloads/BEW/1000 Basic English Words 1_Word List_ENG.xlsx",
		"/Users/kyl/Downloads/BEW/1000 Basic English Words 2_Word List_ENG.xlsx",
		"/Users/kyl/Downloads/BEW/1000 Basic English Words 3_Word List_ENG.xlsx",
		"/Users/kyl/Downloads/BEW/1000 Basic English Words 4_Word List_ENG.xlsx",
	}

	sb := strings.Builder{}
	count := 0

	for _, file := range files {
		xlFile, err := xlsx.OpenFile(file)
		if err != nil {
			panic(err)
		}

		for _, sheet := range xlFile.Sheets {
			header := sheet.Rows[4]

			numberIndex := -1
			exampleIndex := -1
			for k, cell := range header.Cells {
				if cell.String() == "Number" {
					numberIndex = k
				} else if cell.String() == "Example" {
					exampleIndex = k
				}
			}

			for _, row := range sheet.Rows[5:] {
				text := ""
				for k, c := range row.Cells {
					if k == numberIndex {
						text += c.String() + "\t"
					} else if k == exampleIndex {
						text += c.String() + "\t"
					}
				}
				if strings.TrimSpace(text) != "" {
					count++
					sb.WriteString(text + "\n")
				}
			}
		}
	}

	ioutil.WriteFile("w.csv", []byte(sb.String()), os.ModePerm)
	fmt.Println(count)
}
