package entity

type TimeZone struct {
	Id          int `gorm:"primary_key"`
	Title       string
	DiffHour    int
	DiffMinutes int
	Ident       string
}
