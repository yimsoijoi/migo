package svc_test

import (
	"github.com/yimsoijoi/migo/svc"
	"testing"
)

func TestToResult(t *testing.T) {
	models := []svc.DBModel{
		{TableName: "A", ColumnName: "A.a"},
		{TableName: "A", ColumnName: "A.b"},
		{TableName: "A", ColumnName: "A.b"},
		{TableName: "B", ColumnName: "B.a"},
		{TableName: "B", ColumnName: "B.b"},
		{TableName: "C", ColumnName: "C.a"},
	}

	t.Run("test ToExcelModel() Happy", func(t *testing.T) {
		results := svc.ToExcelModel(models)
		if len(results) != 3 {
			t.Error("invalid result length:", len(results))
		}
		if len(results[0].Rows) != 3 {
			t.Error("invalid rows length of results[0]:", len(results[0].Rows))
		}
		if len(results[1].Rows) != 2 {
			t.Error("invalid rows length of results[1]:", len(results[1].Rows))
		}
		if len(results[2].Rows) != 1 {
			t.Error("invalid rows length of results[2]:", len(results[2].Rows))
		}
	})
}
