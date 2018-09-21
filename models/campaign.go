package models

type Campaign struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

func (c Campaign) GetById(campaignId string) *Campaign {
	return &Campaign{
		"1",
		"测试活动",
		"2018-08-01",
		"2018-10-01",
	}
}
