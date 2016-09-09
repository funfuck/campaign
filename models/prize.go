package models

type TPrize struct {
	MyModel
	FpId uint `sql:"index"`
	Name string
	Volume uint
	PrizeBaht uint
}
