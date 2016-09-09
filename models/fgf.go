package models

import (
	"time"
	"encoding/json"
	r "campaign/databases/redis"
	"github.com/garyburd/redigo/redis"
	"sort"
)

type TFriendGetFriend struct {
	MyModel
	Name string
	StartDate time.Time
	EndDate time.Time
	Image string
	DetailTh string
	DetailEn string
	LabelTh string
	LabelEn string
	ImageSuccess string
	DetailSuccessTh string
	DetailSuccessEn string
	DownLinePoint int
	IsActive bool
	Uplines TUplines
	MatchedUpline interface{} `gorm:"ignoring"`
}

func GetActiveFGF() *TFriendGetFriend {
	session := r.Conn()
	defer session.Close()
	campaign, _ :=  redis.String(session.Do("GET", "dooadsCampaignFgf"))

	if campaign != "" {
		fgf := TFriendGetFriend{}
		json.Unmarshal([]byte(campaign), &fgf)
		return &fgf
	} else {
		return nil
	}
}

func MatchedPoint(uplines TUplines, friendNo int) interface{} {
	sort.Sort(uplines)
	for i := len(uplines) - 1; i >= 0; i-- {
		if friendNo >= uplines[i].MemberNo {
			return uplines[i]
		}
	}
	return 0.0
}

func IsInCampaignDuration(t *time.Time, fgf *TFriendGetFriend) bool {
	return !t.Before(fgf.StartDate) && !t.After(fgf.EndDate)
}