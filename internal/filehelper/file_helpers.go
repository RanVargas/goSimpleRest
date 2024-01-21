package filehelper

import (
	"booklib/pkg/model"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func GetFilesInThisDirectory(filepath string) []string {
	files := make([]string, 0)
	entries, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		files = append(files, e.Name())
	}
	return files
}

func WriteToPlainTextFile(filePath string, content string) error {
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

func ReadFromPlainTextFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func ReadExcelFile(filePath string) []model.Book {

	books := make([]model.Book, 0)
	xlsx, err := excelize.OpenFile(filePath)
	if err != nil {
		fmt.Println(err)
		return []model.Book{}
	}

	sheetList := xlsx.GetSheetList()
	fmt.Println("Sheets:", sheetList)

	sheetName := sheetList[0]
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return []model.Book{}
	}

	for rowIdx, row := range rows {
		if rowIdx == 0 {
			continue
		}
		for cellIdx, cell := range row {
			var book model.Book
			if cellIdx == 0 {
				i, err := strconv.Atoi(cell)
				if err != nil {
					log.Println("Error retrieving ID from row: " + fmt.Sprint(rowIdx))
				}
				book.ID = i
			} else if cellIdx == 1 {
				book.Author = cell
			} else if cellIdx == 2 {
				book.Title = cell
			}
			books = append(books, book)
		}
	}

	return books
}

func WriteExcelFile(filePath string, books []model.Book) {
	xlsx := excelize.NewFile()
	sheetName := "ExistentBooks"
	xlsx.SetSheetName("Sheet1", sheetName)

	//for
	celCount := 0
	for rowIdx, book := range books {
		cellAddr := columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
		if rowIdx == 0 {
			celCount++
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
			xlsx.SetCellValue(sheetName, cellAddr, "ID")
			celCount++
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
			xlsx.SetCellValue(sheetName, cellAddr, "Author")
			celCount++
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
			xlsx.SetCellValue(sheetName, cellAddr, "Title")
			celCount++
			celCount = 1
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+2)
			xlsx.SetCellValue(sheetName, cellAddr, book.ID)
			celCount++
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+2)
			xlsx.SetCellValue(sheetName, cellAddr, book.Author)
			celCount++
			cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+2)
			xlsx.SetCellValue(sheetName, cellAddr, book.Title)
			continue
		}
		celCount++
		cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
		xlsx.SetCellValue(sheetName, cellAddr, book.ID)
		celCount++
		cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
		xlsx.SetCellValue(sheetName, cellAddr, book.Author)
		celCount++
		cellAddr = columnIndexToLetter(celCount) + fmt.Sprint(rowIdx+1)
		xlsx.SetCellValue(sheetName, cellAddr, book.Title)
		celCount = 0
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
