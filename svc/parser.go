package svc

func ToExcelModel(models []DBModel) []ExcelModel {
	results := []ExcelModel{}
	index := map[string]int{}
	for _, model := range models {
		tableName := model.TableName
		if _, ok := index[tableName]; !ok {
			results = append(results, ExcelModel{
				TableName: tableName,
				Rows:      []Row{getRow(model)},
			})
			index[tableName] = len(results) - 1
		} else {
			results[index[tableName]].Rows = append(results[index[tableName]].Rows, getRow(model))
		}
	}
	return results
}

func getRow(model DBModel) Row {
	return Row{
		ColumnName:   model.ColumnName,
		Constraint:   model.Constraint,
		DataType:     model.DataType,
		Size:         model.Size,
		Required:     model.Required,
		ExampleValue: model.ExampleValue,
		DefaultValue: model.DefaultValue,
		Index:        model.Index,
	}
}
