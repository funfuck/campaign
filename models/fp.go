package models

import (
	"time"
	r "campaign/databases/redis"
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"campaign/vars"
)

type TFlashPoint struct {
	MyModel
	Name string
	StartDate time.Time
	EndDate time.Time
	MinAge uint
	MaxAge uint
	MinReg time.Time
	MaxReg time.Time
	Repeat uint
	IsActive bool
	Prizes []*TPrize
	Adss []*TFpAds
	Operators []*TOperator
	Genders []*Gender
}

func GetMatchedCampaign(param *vars.ParamGetPrizeBody) *[]TFlashPoint{
	// compare : campaignId, campaignDuration, age, register date, operator
	return getFpRedis(param)
}

func getFpRedis(param *vars.ParamGetPrizeBody) *[]TFlashPoint {
	session := r.Conn()
	defer session.Close()

	all, _ := session.Do("keys", "dooadsCampaignFp_*")
	s := all.([]interface{})

	fps := make([]TFlashPoint, 0)
	for _, v := range s {
		key, _ := redis.String(v, nil)
		item, _ := redis.String(session.Do("GET", key))

		fp := TFlashPoint{}
		json.Unmarshal([]byte(item), &fp)

		if validateAds(param, &fp) &&
			validateAge(param, &fp) &&
			validateRegisterDate(param, &fp) &&
			validateGender(param, &fp) &&
			validateOperator(param, &fp) {

			fps = append(fps, fp)
		}
	}

	return &fps
}

func validateAge(param *vars.ParamGetPrizeBody, fp *TFlashPoint) bool {
	return !(param.MemberProfile.Age < fp.MinAge) && !(param.MemberProfile.Age > fp.MaxAge)
}

func validateRegisterDate(param *vars.ParamGetPrizeBody, fp *TFlashPoint) bool {
	return  !param.MemberProfile.RegisterDate.Before(fp.MinReg) && !param.MemberProfile.RegisterDate.After(fp.MaxReg)
}

func validateGender(param *vars.ParamGetPrizeBody, fp *TFlashPoint) bool {
	for _, v := range fp.Genders {
		if v.ID == param.MemberProfile.Gender {
			return true
		}
	}
	return false
}

func validateOperator(param *vars.ParamGetPrizeBody, fp *TFlashPoint) bool {
	for _, v := range fp.Operators {
		if v.ID == param.MemberProfile.Operator {
			return true
		}
	}
	return false
}

func validateAds(param *vars.ParamGetPrizeBody, fp *TFlashPoint) bool {
	for _, v := range fp.Adss {
		if v.ID == param.AdsId {
			return true
		}
	}
	return false
}