package models

type TFpAds struct {
	MyModel
	FpId uint `sql:"index"`
	Name string
	Length float32
}