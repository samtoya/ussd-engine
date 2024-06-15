package dtos

import "encoding/json"

type AfricasTalkingDto struct {
	NetworkCode string `json:"networkCode,omitempty" form:"networkCode" param:"networkCode"`
	PhoneNumber string `json:"phoneNumber,omitempty" form:"phoneNumber" param:"phoneNumber"`
	ServiceCode string `json:"serviceCode,omitempty" form:"serviceCode" param:"serviceCode"`
	SessionId   string `json:"sessionId,omitempty" form:"sessionId" param:"sessionId"`
	Text        string `json:"text" form:"text" param:"text"`
}

func (a *AfricasTalkingDto) String() string {
	jb, _ := json.Marshal(a)
	return string(jb)
}
