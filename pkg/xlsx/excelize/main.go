package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	// file := ""
	// r, err := excelize.OpenFile(file)

	f := excelize.NewFile()
	// Create a new sheet.
	firstSheet := "Sheet1"

	f.NewSheet(firstSheet)

	f.InsertRow(firstSheet, 3)
	f.InsertRow(firstSheet, 3)

	styleID, _ := f.GetCellStyle("Sheet1", "A1")
	x := f.Styles.CellXfs.Xf[styleID]
	fmt.Println(x)

	style, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			Horizontal: "right",
			Vertical:   "top",
		},
		Border: []excelize.Border{
			{Type: "left", Color: "606060", Style: 1},
			{Type: "top", Color: "606060", Style: 1},
			// {Type: "bottom", Color: "FFFF00", Style: 4},
			// {Type: "right", Color: "FF0000", Style: 4},
			// {Type: "diagonalDown", Color: "A020F0", Style: 7},
			// {Type: "diagonalUp", Color: "A020F0", Style: 8},
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStr(firstSheet, "A3", "test")
	f.SetCellStr(firstSheet, "A4", "test")
	f.SetCellStr(firstSheet, "A5", "test")
	f.SetCellStr(firstSheet, "A6", "test")
	f.InsertRow(firstSheet, 5)
	f.SetCellStr(firstSheet, "A100", "test")
	err = f.SetCellStyle(firstSheet, "A3", "A100", style)
	if err != nil {
		fmt.Println(err)
	}
	err = f.SetRowStyle(firstSheet, 1, 100, style)
	if err != nil {
		fmt.Println(err)
	}
	// rightWithBorder, _ := f.GetCellStyle(firstSheet, "A8")

	secondSheet := "Sheet2"
	index2 := f.NewSheet(secondSheet)
	f.SetCellValue("Sheet2", "A2", "Hello world.")

	// Set active sheet of the workbook.
	f.SetActiveSheet(index2)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
