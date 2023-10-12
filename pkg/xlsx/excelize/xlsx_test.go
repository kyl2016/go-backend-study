package main

import (
	"fmt"
	"testing"

	"github.com/xuri/excelize/v2"
)

func TestGetStyle(t *testing.T) {
	f, _ := excelize.OpenFile("./实名认证服务器数据统计-0429.xlsx")
	firstSheet := f.GetSheetName(0)
	styleID, _ := f.GetCellStyle(firstSheet, "A8")
	x := f.Styles.CellXfs.Xf[styleID]
	fmt.Println(*x.BorderID)
	fmt.Println(*x.Alignment)
	fmt.Println(*x.FillID)
	fmt.Println(x.Alignment.Horizontal)
}
