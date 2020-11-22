package rule

type Rule struct {
	Id     int `gorm:"primary_key"`
	Object string
	Action string
	Title  string
}
