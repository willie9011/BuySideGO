package Structs

//----------------API結構體---------------------

type UrlQuery struct {
	UserName string `json:"username" form:"username" binding:"required"` //required 代表如果url列表沒有這項伺服器將會出錯
	UserID   string `json:"user_id" form:"user_id" binding:"required"`
}

type User struct {
	ID   string `json:"id" binding:"required" gorm:"primaryKey"`
	Name string `json:"name" binding:"required" gorm:"column:name"`
	Age  uint   `json:"age" binding:"required,gt=18" gorm:"column:age"`
}

//{"id":"123","name":"Alice","age":25}

type MultiUsers struct {
	UserList     []User `json:"user_list" binding:"gt=0,lt=4,dive"`
	UserListSize int    `json:"user_list_size" binding:"required,gt=0,lt=4"`
}

type Vegetable struct {
	Date   string  `json:"date" binding:"required"`
	Type   string  `json:"type" binding:"required"`
	ID     string  `json:"id" binding:"required"`
	Name   string  `json:"name" binding:"required"`
	Market string  `json:"market" binding:"required"`
	Up     float64 `json:"up" binding:"required"`
	Down   float64 `json:"down" binding:"required"`
	Avg    float64 `json:"avg" binding:"required"`
	Wave   string  `json:"wave" binding:"required"`
}
type Vegetables struct {
	VegeList     []Vegetable `json:"vege_list" binding:"dive"`
	VegeListSize int         `json:"vege_list_size" binding:"required,gt=0,"`
}

//測試資料:
//{
//  "user_list": [
//    {"id": "1", "name": "Alice", "age": 25},
//    {"id": "2", "name": "Bob", "age": 30},
//    {"id": "3", "name": "Charlie", "age": 28}
//  ],
//  "user_list_size": 3
//}
