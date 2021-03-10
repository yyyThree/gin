package entity

// ItemSearches 商品搜索表		
type ItemSearches struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Appkey    string    `gorm:"uniqueIndex:item_search;column:appkey;type:varchar(64);not null" json:"appkey"`
	Channel   int       `gorm:"uniqueIndex:item_search;column:channel;type:int(11);not null" json:"channel"`
	ItemID    string    `gorm:"uniqueIndex:item_search;column:item_id;type:varchar(64);not null" json:"item_id"` // 商品ID
	SkuID     string    `gorm:"uniqueIndex:item_search;column:sku_id;type:varchar(64);not null" json:"sku_id"`   // 商品SKU_ID
	ItemName  string    `gorm:"column:item_name;type:varchar(255);not null" json:"item_name"`                    // 商品名称
	SkuName   string    `gorm:"column:sku_name;type:varchar(255);not null" json:"sku_ame"`                      // 商品SKU名称
	Barcode   string    `gorm:"column:barcode;type:varchar(50);not null" json:"barcode"`                        // 条形码
	ItemState int8      `gorm:"column:item_state;type:tinyint(4);not null" json:"item_state"`                    // 商品状态 0 正常 1 已删除 2 已彻底删除
	SkuState  int8      `gorm:"column:sku_state;type:tinyint(4);not null" json:"sku_state"`                      // sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除
	UpdatedAt DateTime `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	CreatedAt DateTime `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
}

// TableName get sql table name.获取数据库表名
func (m *ItemSearches) TableName() string {
	return "item_searches"
}

// ItemSearchesColumns get sql column name.获取数据库列名
var ItemSearchesColumns = struct {
	ID        string
	Appkey    string
	Channel   string
	ItemID    string
	SkuID     string
	ItemName  string
	SkuName   string
	Barcode   string
	ItemState string
	SkuState  string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	Appkey:    "appkey",
	Channel:   "channel",
	ItemID:    "item_id",
	SkuID:     "sku_id",
	ItemName:  "item_name",
	SkuName:   "sku_name",
	Barcode:   "barcode",
	ItemState: "item_state",
	SkuState:  "sku_state",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// Items 商品表		
type Items struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Appkey    string    `gorm:"uniqueIndex:item;column:appkey;type:varchar(64);not null" json:"appkey"`
	Channel   int       `gorm:"uniqueIndex:item;column:channel;type:int(11);not null" json:"channel"`
	ItemID    string    `gorm:"uniqueIndex:item;column:item_id;type:varchar(64);not null" json:"item_id"` // 商品ID
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`                      // 商品名称			
	Photo     string    `gorm:"column:photo;type:varchar(512);not null" json:"photo"`                    // 商品主图			
	Detail    string    `gorm:"column:detail;type:text;not null" json:"detail"`                          // 商品详情			
	State     int8      `gorm:"column:state;type:tinyint(4);not null" json:"state"`                      // 商品状态 0 正常 1 已删除 2 已彻底删除			
	UpdatedAt DateTime `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	CreatedAt DateTime `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Items) TableName() string {
	return "items"
}

// ItemsColumns get sql column name.获取数据库列名
var ItemsColumns = struct {
	ID        string
	Appkey    string
	Channel   string
	ItemID    string
	Name      string
	Photo     string
	Detail    string
	State     string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	Appkey:    "appkey",
	Channel:   "channel",
	ItemID:    "item_id",
	Name:      "name",
	Photo:     "photo",
	Detail:    "detail",
	State:     "state",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

// Skus [...]		
type Skus struct {
	ID        int       `gorm:"primaryKey;column:id;type:int(11);not null" json:"id"`
	Appkey    string    `gorm:"uniqueIndex:sku;column:appkey;type:varchar(64);not null" json:"appkey"`
	Channel   int       `gorm:"uniqueIndex:sku;column:channel;type:int(11);not null" json:"channel"`
	ItemID    string    `gorm:"uniqueIndex:sku;column:item_id;type:varchar(64);not null" json:"item_id"` // 商品ID
	SkuID     string    `gorm:"uniqueIndex:sku;column:sku_id;type:varchar(64);not null" json:"sku_id"`   // 商品SKU_ID
	Name      string    `gorm:"column:name;type:varchar(255);not null" json:"name"`                     // 商品SKU名称			
	Photo     string    `gorm:"column:photo;type:varchar(512);not null" json:"photo"`                   // 商品SKU主图			
	Barcode   string    `gorm:"column:barcode;type:varchar(50);not null" json:"barcode"`               // 条形码
	State     int8      `gorm:"column:state;type:tinyint(4);not null" json:"state"`                     // sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除			
	UpdatedAt DateTime `gorm:"column:updated_at;type:datetime;not null" json:"updated_at"`
	CreatedAt DateTime `gorm:"column:created_at;type:datetime;not null" json:"created_at"`
}

// TableName get sql table name.获取数据库表名
func (m *Skus) TableName() string {
	return "skus"
}

// SkusColumns get sql column name.获取数据库列名
var SkusColumns = struct {
	ID        string
	Appkey    string
	Channel   string
	ItemID    string
	SkuID     string
	Name      string
	Photo     string
	Barcode   string
	State     string
	UpdatedAt string
	CreatedAt string
}{
	ID:        "id",
	Appkey:    "appkey",
	Channel:   "channel",
	ItemID:    "item_id",
	SkuID:     "sku_id",
	Name:      "name",
	Photo:     "photo",
	Barcode:   "barcode",
	State:     "state",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}
