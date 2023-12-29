package svc

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func BuildXlsx(models []ExcelModel) *excelize.File {
	file := excelize.NewFile()

	for _, model := range models {
		sheetName := model.TableName
		_ = file.NewSheet(sheetName)

		// Set header column
		file.SetCellValue(sheetName, "A1", "column_name")
		file.SetCellValue(sheetName, "B1", "constraint")
		file.SetCellValue(sheetName, "C1", "data_type")
		file.SetCellValue(sheetName, "D1", "size")
		file.SetCellValue(sheetName, "E1", "required")
		file.SetCellValue(sheetName, "F1", "example_value")
		file.SetCellValue(sheetName, "G1", "default_value")
		file.SetCellValue(sheetName, "H1", "index")

		// Set record to rows
		for i, row := range model.Rows {
			file.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), row.ColumnName)
			file.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), row.Constraint)
			file.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), row.DataType)
			file.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), row.Size)
			file.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), row.Required)
			file.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), row.ExampleValue)
			file.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), row.DefaultValue)
			file.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), row.Index)
		}
	}
	return file
}
