package models

type Gender struct {
	MyModel
	FpId uint `sql:"index"`
	Name string
}
