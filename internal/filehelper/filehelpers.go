package filehelper

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xuri/excelize/v2"
)

func writeToPlainTextFile(filePath string, content string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Println("Content written to", filePath)
	return nil
}

func readFromPlainTextFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func readExcelFile(filePath string) {
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	sheetList := xlsx.GetSheetList()
	fmt.Println("Sheets:", sheetList)

	sheetName := sheetList[0]
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows {
		for _, cell := range row {
			fmt.Print(cell, "\t")
		}
		fmt.Println()
	}
}

func writeExcelFile(filePath string) {
	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	xlsx.SetSheetName("Sheet1", sheetName)

	data := [][]interface{}{
		{"Name", "Age", "City"},
		{"John Doe", 30, "New York"},
		{"Jane Doe", 25, "Los Angeles"},
		{"Bob Smith", 35, "Chicago"},
	}

	for rowIdx, row := range data {
		for colIdx, cell := range row {
			cellAddr := columnIndexToLetter(colIdx+1) + fmt.Sprint(rowIdx+1)
			xlsx.SetCellValue(sheetName, cellAddr, cell)
		}
	}

	if err := xlsx.SaveAs(filePath); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Excel file saved successfully.")
}

func columnIndexToLetter(index int) string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letter := ""
	for index > 0 {
		index--
		letter = string(alphabet[index%26]) + letter
		index /= 26
	}
	return letter
}
