package svc

type DBModel struct {
	TableName    string `db:"table_name"`
	ColumnName   string `db:"column_name"`
	Constraint   string `db:"constraint"`
	DataType     string `db:"data_type"`
	Size         string `db:"size"`
	Required     string `db:"required"`
	ExampleValue string `db:"example_value"`
	DefaultValue string `db:"default_value"`
	Index        string `db:"index"`
}

type ExcelModel struct {
	TableName string
	Rows      []Row
}

type Row struct {
	ColumnName   string
	Constraint   string
	DataType     string
	Size         string
	Required     string
	ExampleValue string
	DefaultValue string
	Index        string
}
