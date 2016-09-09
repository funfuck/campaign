package models

type TOperator struct {
	MyModel
	FpId uint `sql:"index"`
	Name string
}
