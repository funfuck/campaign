package models

//import "github.com/jinzhu/gorm"

type TUpline struct {
	MyModel
	FgfId uint `sql:"index"`
	//TFriendGetFriend `gorm:"ignoring"`
	MemberNo int
	UlType string
	Point float32
	Formula string
}

func init() {

}

type TUplines []*TUpline

func (s TUplines) Len() int {
	return len(s)
}

func (s TUplines) Less(i, j int) bool {
	return s[i].Point < s[j].Point
}

func (s TUplines) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
