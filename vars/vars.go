package vars

import (
	"time"
)

/*
=======================================================
 */

type RequestGetPointBody struct {
	RefId string
	MemberId string
}

/*
=======================================================
 */

type MemberProfile struct {
	Age uint
	Gender uint
	Operator uint
	RegisterDate time.Time
}

type ParamGetPrizeBody struct {
	MemberProfile *MemberProfile
	AdsId uint
	WatchedCount uint
}

/*
=======================================================
 */

//type Prize struct {
//	Id uint
//	Name string
//	PrizePoint float32
//	Volume uint
//	CreateBy uint
//	LastUpdBy uint
//}
//
//type FPPrize struct {
//	Id uint
//	Name string
//	Repeat uint
//	Prize []*Prize
//}

/*
=======================================================
 */

type Response struct {
	Success bool
	ResultCode string
	Path string
	ErrorDescription string
	DeveloperMessage string
	TimeStamp time.Time
	Result interface{}
	ServerError string
	Method string
	Header interface{}
	RequestBody interface{}
	LogMessage string
}
