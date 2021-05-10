<div align=center>
<img src="https://golang.org/lib/godoc/images/footer-gopher.jpg"/>
</div>
<div align=center>
<img src="https://img.shields.io/badge/golang-v1.15-blue"/>
<img src="https://img.shields.io/badge/gin-v1.6.3-lightBlue"/>
<img src="https://img.shields.io/badge/gorm-v1.21.5-lightBlue"/>
<img src="https://img.shields.io/badge/go redis-v8.8.2-lightBlue"/>
<img src="https://img.shields.io/badge/zap-v1.16.0-lightBlue"/>
<img src="https://img.shields.io/badge/amqp-v1.0.0-lightBlue"/>
</div>

## 一、项目介绍
本项目是基于`gin`框架开发的`web`框架，集成了`gorm`、`jwt`、`zap`、`redis`、`rabbitmq`等核心组件。<br>
示例代码已实现商品服务的完整功能，启动服务后，可导入`doc`目录下的`postman`测试用例和测试环境一键化运行通过。

## 二、功能点
- 支持`restful`路由
- 支持统一捕获`404`/`500`错误
- 支持`JWT`校验
- 支持统一接口输出
- 支持`gorm`使用
- 支持自动生成数据库结构体  
- 支持数据库迁移
- 支持统一接口参数校验  
- 支持<a href="https://github.com/yyyThree/zap" target="_blank">zap</a>日志收集
- 支持<a href="https://github.com/yyyThree/rabbitmq" target="_blank">rabbitmq</a>使用
- 支持`redis`使用 
- 支持`viper`配置文件解析
- 支持单元测试  
- 支持平滑关闭服务器

## 三、目录结构
```
├── config/         ----- 各项配置文件目录
│   └── config.go       ----- 基础配置
├── constant/       ----- 全局常量/变量定义文件目录
│   ├── common.go       ----- 通用定义
│   ├── db.go           ----- 数据库定义
│   └── type.go         ----- 通用定义类型
├── router/         ----- 路由文件目录
│   ├── group/          ----- 子路由文件目录 
│   │   └── no_router.go     ----- 404路由
│   └── router.go   ----- 核心路由文件 
├── middleware/     ----- 中间件文件目录
│   │   ├── jwt.go      ----- JWT统一检验
│   │   └── panic.go    ----- 500错误统一监听处理
├── controller/     ----- 接口控制器文件目录
├── service/        ----- 逻辑处理层文件目录
├── model/          ----- 各项数据模型文件目录
│   ├── db/             ----- 数据库连接文件目录
│   │   └── db.go           ----- 核心文件
│   ├── entity/         ----- 数据库结构体文件目录
│   │   ├── time.go         ----- 通用标准时间格式文件
│   │   └── config.yaml     ----- 数据库结构体生成配置文件
│   ├── field/          ----- 数据库字段对应关系文件目录
│   └── param/          ----- 接口传参结构体文件目录
├── dao/            ----- 数据库操作层文件目录
│   └── common.go       ----- 通用数据库操作方法
├── library/        ----- 类库目录
│   ├── log/            ----- 日志库
│   │   └── log.go          ----- 核心文件
│   ├── rabbitmq/       ----- rabbitmq库
│   │   ├── common/         ----- 常量配置文件目录
│   │   │   ├── msg.go          ----- 路由键值定义文件
│   │   │   └── queue.go        ----- 队列定义文件
│   │   ├── subscriber/     ----- 队列订阅者文件目录
│   │   ├── subscriber.go   -----  队列订阅者启动配置
│   │   └── rabbitmq.go     ----- 核心文件
│   ├── redis/          ----- redis库
│   │   └── redis.go        ----- 核心文件
│   ├── token/          ----- token库
│   │   └── token.go        ----- 核心文件
│   ├── valid           ----- 通用校验库
│   │   ├── validator.go    ----- 各项校验方法
│   │   └── valid.go        ----- 核心文件
├── helper/         ----- 辅助函数目录
├── output/         ----- 统一输出文件目录
│   ├── code/           -----  项目错误码文件目录
│   │   ├── code.go             ----- 项目错误码定义文件
│   │   └── common_string.go    ----- 通过go generate自动生成的错误码信息文件
│   ├── status.go       ----- 项目内部错误结构文件 
│   └── response.go     ----- 核心输出文件
├── migration/      ----- 数据库迁移文件目录
├── doc/            ----- 项目各类静态文档目录
├── test/           ----- 单元测试文件目录
│   ├── config.yaml     ----- 核心配置文件
│   ├── redis_test.go   ----- redis单元测试文件 
│   └── token_test.go   ----- jwt单元测试文件
├── config.yaml     ----- 核心配置文件
├── init.go         ----- 初始化操作文件
├── main.go         ----- 项目启动文件
├── go.mod 
└── go.sum
```

## 四、使用文档
### 一、配置文件解析（`config.yaml`）
```
app:
  env: debug # 项目运行模式 debug/production
  token_secret: token # jwt密钥
  
http:
  port: 10080 # 项目启动端口
  read_time_out: 10 # 接口读取请求最大超时时间(s)
  write_time_out: 10 # 接口返回最大超时时间(s)
  shutdown_time_out: 10 # 平滑关闭等待时间(s)
  
db:
  master: # 主数据库配置
    host: 127.0.0.1
    port: 3355
    user: go
    password: go
    max_open_connections: 10 # 最大连接数
    max_idle_connections: 5 # 最大空闲连接数，即始终保持连接的数量
    max_connection_idle_time: 300 # 连接可复用的最大时间(s)
    
  slave:
    host: 127.0.0.1
    port: 3355
    user: go
    password: go
    max_open_connections: 10
    max_idle_connections: 2
    max_connection_idle_time: 300
    
redis:
  address: 127.0.0.1:8003
  password: ""
  db: 0
  connect_timeout: 10 # 连接超时时间(s)
  read_timeout: 5 # 读取数据超时时间(s)
  write_timeout: 5 # 写入数据超时时间(s)
  pool_size: 5 # 连接池数量
  
log:
  out: file # 日志写入方式：file 文件 stdout 屏幕 redis redis队列
  dir: ./log/ # out 为 file 时，需要提供存储的文件夹
  redisKey: go # out 为 redis 时，需要提供redis列表的key
  
rabbitmq:
  host: 127.0.0.1
  port: 5673
  user: go
  password: go
  vhost: go
  admin_user: goadmin # vhost对应管理员账号
  admin_password: goadmin # vhost对应管理员密码
  ex_direct: go.direct # 基础直连交换机名称，业务系统默认使用
  ex_topic: go.topic # 主题交换机名称
  ex_death_letter: go.dl # 死信交换机名称
  ttl_queue_msg: 86400000 # 队列中消息有效期(ms)
  ttl_msg: 86400000 # 每条消息的有效期(ms)
  log_dir: ./log/rabbitmq # 日志存储文件夹地址
```
### 二、路由（`router/group/`）
```go
func (item *Item) Router(router *gin.Engine) {
	itemRouter := router.Group("/item")
	itemRouter.Use(new(middleware.Item).Handler())
	itemController := new(controller.Item)
	{
		itemRouter.POST("/add", itemController.Add)
		itemRouter.PUT("/update", itemController.Update)
		itemRouter.DELETE("/delete", itemController.Delete)
		itemRouter.PATCH("/recover", itemController.Recover)
		itemRouter.GET("/get", itemController.Get)
		itemRouter.GET("/search", itemController.Search)
	}
}
```
### 三、中间件（`middleware/`）
```go
func (item *Item) Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("item中间件开始")

		// before request

		c.Next()

		// after request
		fmt.Println("item中间件结束")
	}
}
```
### 四、控制器（`controller/`）
```go
type Item struct {
}
// 添加商品
func (item *Item) Add(c *gin.Context) {
	params := param.ItemAdd{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, err := (&service.Item{}).Add(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	output.Response(c, &output.SucResponse{
		Data: data,
	}, nil)
	return
}

// 更新商品
func (item *Item) Update(c *gin.Context) {
	params := param.ItemUpdate{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Update(params)
	output.Response(c, nil, err)
	return
}

// 删除商品
func (item *Item) Delete(c *gin.Context) {
	params := param.ItemDelete{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Delete(params)
	output.Response(c, nil, err)
	return
}

// 恢复商品
func (item *Item) Recover(c *gin.Context) {
	params := param.ItemRecover{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	err := (&service.Item{}).Recover(params)
	output.Response(c, nil, err)
	return
}

// 获取商品详情
func (item *Item) Get(c *gin.Context) {
	params := &param.ItemGet{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, err := (&service.Item{}).Get(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	output.Response(c, &output.SucResponse{
		Data: helper.FilterStructByFields(data, params.Fields, field.GetItemFields()),
	}, nil)
	return
}

// 搜索商品列表
func (item *Item) Search(c *gin.Context) {
	params := &param.ItemSearch{}
	if err := c.ShouldBind(&params); err != nil {
		output.Response(c, nil, output.Error(code.ParamBindErr))
		return
	}
	helper.AppendTokenParams(c, &params.Common)

	data, total, err := (&service.ItemSearch{}).Search(params)
	if err != nil {
		output.Response(c, nil, err)
		return
	}

	output.Response(c, &output.ListResponse{
		SucResponse: &output.SucResponse{
			Data: helper.FilterStructsByFields(data, params.Fields, field.GetItemFields()),
		},
		Total: total,
	}, nil)
	return
}
```
### 五、业务逻辑（`service/`）
```go
type ItemSearch struct {
}

// 搜索商品列表
func (itemSearch *ItemSearch) Search(params *param.ItemSearch) (data []ItemDetail, total int, err error) {
	if err = itemSearch.checkSearch(params); err != nil {
		return
	}

	// step1 初始化默认值
	params.Fields = helper.GetString(params.Fields, "*")

	// step2 搜索item_id
	// 构建搜索条件
	sqlBuild := itemSearch.buildSearch(params)
	itemIds, total, _ := dao.NewItemSearch().SearchItem(sqlBuild)

	if helper.IsEmpty(itemIds) {
		return
	}

	// step3 根据itemIds获取商品详情
	for _, itemId := range itemIds {
		item, _ := (&Item{}).Get(&param.ItemGet{
			ItemId: itemId,
			Fields: params.Fields,
			Common: params.Common,
		})
		if !helper.IsEmpty(item) {
			data = append(data, item)
		}
	}

	return
}
```
### 六、数据库结构体文件生成（`model/entity/`）
```go
// 提前安装gormt：go get -u -v github.com/xxjwxc/gormt@master

cd model/entity
gormt // 文件将于 model/entity 下生成
```
### 七、数据库操作（`dao/`）
```go
type item struct {
    dao
}
func NewItem(txs ...*gorm.DB) *item {
	return &item{dao{
		Tx:     GetTx(txs...),
		DbName: constant.DbServiceItems,
	}}
}

func (item *item) Insert(insert *entity.Items) (data entity.Items, err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}

	err = db.Create(&insert).Error
	if err != nil {
		return
	}
	data = *insert
	return
}

func (item *item) GetOne(fields []string, where map[string]interface{}) (data entity.Items, err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}
	err = db.Select(fields).
		Where(where).
		Limit(1).
		Find(&data).Error
	return
}

func (item *item) Update(update map[string]interface{}, where map[string]interface{}, limit int) (err error) {
	db, err := item.GetDb()
	if err != nil {
		return
	}
	err = db.Model(&entity.Items{}).
		Where(where).
		Limit(limit).
		Updates(update).Error
	if err != nil {
		return
	}
	return
}

```
### 八、统一输出
1. 错误码（`output/code`）
   - 错误码文件以`_code`结尾
   - 错误码全局唯一
   - 对应的注释即为错误码对应的错误信息 
2. 错误码使用
   ```go
   // 错误码
   err := output.Error(code.ParamBindErr)
    
   // 错误码 + 错误数据
   err := output.Error(code.IllegalParams).WithDetails(err)
    ```
3. 基于错误码文件及其注释，自动生成错误码对应的错误信息文件
    ```go
    cd output/code
    go generate
    ```
4. 多种输出方式
    ```go
    // 输出错误
    output.Response(c, nil, err)
   
    // 成功输出
    output.Response(c, nil, err)
   
    // 带数据返回的成功输出
    output.Response(c, &output.SucResponse{
        Data: data,
    }, nil)
   
    // 带数据返回的列表成功输出 
    output.Response(c, &output.ListResponse{
        SucResponse: &output.SucResponse{
            Data: helper.FilterStructsByFields(data, params.Fields, field.GetItemFields()),
        },
        Total: total,
    }, nil)     
    ```
### 九、数据库迁移
- 在`migration`目录下添加`sql`文件
- `sql`文件以时间戳开头
- 执行迁移时将按照时间戳的大小顺序排序后执行
- 已执行过的迁移文件不会再次执行
### 十、redis使用
```go
res, err := redis.GetConn().Set(redis.GetCtx(), key, 1, 0).Result()
if err != nil {
    t.Fatal("redis设置失败", err)
}
fmt.Println("redis设置成功\n", res)
```
### 十一、rabbitmq使用
1. 定义路由键值(`library/common/msg.go`)
    ```go
    // 消息路由
    const (
        ItemSync = "item.sync" // 商品数据同步
    )
    
    // 死信消息路由
    const (
        ItemDl = "item.dl" // 商品通用
    )
    ```
2. 定义队列(`library/common/queue.go`)
    ```go
    type Queue struct {
        Name   string   // 队列名
        Keys   []string // 队列绑定的路由键值
        DlxKey string   // 队列绑定的死信队列路由
    }
    
    // 带死信参数的直连交换机队列
    var QueueDirectWithDlList = []Queue{SyncItemSearch}
    
    var (
        SyncItemSearch = Queue{
            Name:   "syncItemSearch",
            Keys:   []string{ItemSync},
            DlxKey: ItemDl,
        }
    )
    
    // 死信队列
    var QueueDlList = []Queue{ItemDlQueue}
    
    var (
        ItemDlQueue = Queue{
            Name: "itemDl",
            Keys: []string{ItemDl},
        }
    )
    ```
3. 定义队列消费者(`library/subscriber/`)
    ```go
    // 同步商品搜索数据
    func SyncItemSearch() {
        queue := common.SyncItemSearch
        go func() {
            _ = rabbitmq.Subscribe(queue.Name, func(msg amqp.Delivery) {
                params := param.ItemSync{}
                _ = json.Unmarshal(msg.Body, &params)
    
                if helper.HasAnyEmpty(params.ItemId, params.SyncType) {
                    rabbitmq.Reject(msg)
                    return
                }
                err := (&service.ItemSearch{}).Sync(params)
                if err != nil {
                    log.GetLogger().Info("SyncItemSearch", zap.BaseMap{
                        "queue":  queue,
                        "params": params,
                        "error":  err,
                    })
                    fmt.Println("SyncItemSearch err：", err.Error())
                    rabbitmq.Nack(msg)
                    return
                }
                rabbitmq.Ack(msg)
            })
        }()
    }
    ```
4. 定义需要启动的消费者(`library/subscriber.go`)
    ```go
    // 消息订阅
    // 需要启动的的订阅者放这里
    var Subscribers = []func(){
        // 商品模块
        subscriber.SyncItemSearch,
    
        // 死信队列
        subscriber.ItemDl,
    }
    ```
5. 声明队列和消费者（启动服务时自动执行）
6. 发布消息
    ```go
    _ = rabbitmq.PublishWithConfirm(common.ItemSync, helper.StructToJson(param.ItemSync{
        ItemId:   params.ItemId,
        SyncType: constant.ItemSyncTypeRecover,
        Common:   params.Common,
    }))
    ```
### 十二、zap日志使用
```go
log.GetLogger().Error("itemRecover", zap.BaseMap{
    "appkey":  params.AppKey,
    "channel": params.Channel,
    "item_id": params.ItemId,
    "err":     err,
})
```
### 十三、单元测试
1. 配置`test/config.yaml`
2. 执行测试
   ```go
   cd test
   // jwt测试
   go test token_test.go -v
   
   // 测试redis
   go test redis_test.go -v
   ```
### 十四、编译
```go
// make help 查看更多命令
make // 编译 Go 代码, 生成二进制文件
make start // 启动服务
```

## 五、示例功能（商品服务）
1. 数据结构
    ```go
    items：商品表
        id
        appkey varchar(64) 
        channel int(11)
        item_id varchar(64) 商品ID
        name varchar(255) 商品名称
        photo varchar(512) 商品主图
        detail text 商品详情
        state tinyint(4) 商品状态 0 正常 1 已删除 2 已彻底删除
        updated_at datetime
        created_at datetime
        索引：appkey + channel + item_id 唯一
       
    skus：商品sku表
        id
        appkey varchar(64)
        channel int(11)
        item_id varchar(64) 商品ID
        sku_id varchar(64) 商品SKU_ID
        name varchar(255) 商品SKU名称
        photo varchar(512) 商品SKU主图
        barcode varchar(50) 条形码
        state tinyint(4) sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除
        updated_at datetime
        created_at datetime   
        索引：appkey + channel + item_id + sku_id 唯一
    
    item_searches：商品搜索表
        id
        appkey varchar(64)
        channel int(11)
        item_id varchar(64) 商品ID
        sku_id varchar(64) 商品SKU_ID
        item_name varchar(255) 商品名称
        sku_name varchar(255) 商品SKU名称
        barcode varchar(50) 条形码
        item_state tinyint(4) 商品状态 0 正常 1 已删除 2 已彻底删除
        sku_state tinyint(4) sku状态 0 正常 1 已删除 2 已彻底删除 3 业务上删除
        updated_at datetime
        created_at datetime   
        索引：appkey + channel + item_id + sku_id 唯一
    ```   
2. 接口
   
   | 接口 | 方法 | 说明 |
   | :-----| :---- | :----: |
   | item/add       | POST      | 添加商品      |
   | item/update    | PUT       | 更新商品      |
   | item/delete    | DELETE    | 删除商品      |
   | item/recover   | PATCH     | 恢复商品      |
   | item/get       | GET       | 获取商品详情  |
   | item/search    | GET       | 搜索商品列表  |
3. 测试运行
    1. 配置`config.yaml`
    2. 运行项目（根目录下执行）
          ```go
          go run . 
          // or
          fresh // 需要安装fresh命令
          ```       
    3. 生成`token`
        - 执行单元测试`go test token_test.go -v`，获取最新可用的`token`
    3. 导入`postman`测试用例
        - 导入`doc/gin.postman_collection.json`和`doc/gin.postman_environment.json`
        - 替换环境变量的`token`值
        - 运行整个用例集合
