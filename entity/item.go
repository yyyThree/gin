package entity

type Item struct {
	PrimaryKey
	StoreId int `gorm:"type:int(11);not null;uniqueIndex:store_item"`
	ItemId  string `gorm:"type:varchar(64);not null;uniqueIndex:store_item"`
	Name    string `gorm:"type:varchar(255);not null"`
	State   int `gorm:"type:tinyint(4);default:0;;not null"`
	Time
}
