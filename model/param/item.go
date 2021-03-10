package param

type ItemAdd struct {
	Name string   `form:"name" json:"name"`
	Photo string  `form:"photo" json:"photo"`
	Detail string `form:"detail" json:"detail"`
	Common
}