package model

type House struct {
	EstateOrg string  `json:"estateOrg"` // 开发商名
	Neighborhood string `json:"neighborhood"` // 小区名
	BuildingId string `json:"buildingId"` // 楼号
	RoomId string `json:"roomId"` // 房屋号
}

