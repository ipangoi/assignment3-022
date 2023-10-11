package model

type Weather struct {
	ID    uint `gorm:"primaryKey" json:"id"`
	Wind  int  `json:"wind"`
	Water int  `json:"water"`
}
