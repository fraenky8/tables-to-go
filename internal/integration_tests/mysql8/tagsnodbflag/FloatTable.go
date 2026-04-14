package dto

import (
	"database/sql"
)

type FloatTable struct {
	Col                      sql.NullFloat64
	FloatNn                  float64
	FloatNnUnique            float64
	FloatNnCheck             float64
	FloatNnRef               float64
	FloatNnDefConst          float64
	FloatNnDefFunc           float64
	FloatNnUniqueCheck       float64
	FloatUnique              sql.NullFloat64
	FloatUniqueCheck         sql.NullFloat64
	FloatUniqueRef           sql.NullFloat64
	FloatUniqueDefConst      sql.NullFloat64
	FloatUniqueDefFunc       sql.NullFloat64
	FloatCheck               sql.NullFloat64
	FloatCheckRef            sql.NullFloat64
	FloatCheckDefConst       sql.NullFloat64
	FloatCheckDefFunc        sql.NullFloat64
	FloatRef                 sql.NullFloat64
	FloatRefUniqueCheck      sql.NullFloat64
	FloatDefConst            sql.NullFloat64
	FloatDefConstUniqueCheck sql.NullFloat64
	FloatDefFunc             sql.NullFloat64
	FloatDefFuncUniqueCheck  sql.NullFloat64
}
