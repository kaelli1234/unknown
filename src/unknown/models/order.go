package models

// import (
// 	"fmt"
// 	"strings"
// 	"time"

// 	"ccgwf/app"

// 	"mrs/common"
// )

// type Order struct {
// 	ID        int       `json:"id"`
// 	RoomId    int       `json:"room_id"`
// 	Date      string    `json:"date"`
// 	Block     int       `json:"block"`
// 	Uid       uint64    `json:"uid"`
// 	Username  string    `json:"username"`
// 	Avatar    string    `json:"avatar"`
// 	Subject   string    `json:"subject"`
// 	Remark    string    `json:"remark"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// func GetAllOrders() (result []Order) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	db.Table("order").Find(&result)
// 	return
// }

// func GetOrderByParams(id, date, block string) (result Order) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	db.Table("order").Where("room_id = ? AND date = ? AND block = ?", id, date, block).First(&result)
// 	return
// }

// func UpdateOrder(id, date, block, uid, username, avatar string) (result Order) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	sql := fmt.Sprintf("INSERT INTO `mrs`.`order` (`room_id`, `date`, `block`, `uid`, `username`, `avatar`) VALUES ('%v', '%v', '%v', '%v', '%v', '%v') ON DUPLICATE KEY UPDATE uid = '%v', username = '%v', avatar = '%v'", id, date, block, uid, username, avatar, uid, username, avatar)
// 	db.Exec(sql)
// 	return
// }

// func GetOrdersByRoomAndDate(id, date string) (result []Order) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	db.Table("order").Where("room_id = ? AND date = ?", id, date).Find(&result)
// 	return
// }

// func InsertOrdersByRoomAndDate(id, date string) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	sql := "INSERT INTO `mrs`.`order` (`room_id`, `date`, `block`) VALUES "
// 	var vals []string
// 	for i := 1; i <= common.Timeblock; i++ {
// 		vals = append(vals, fmt.Sprintf("('%v', '%v', '%v')", id, date, i))
// 	}
// 	sql = sql + strings.Join(vals, ", ")
// 	db.Exec(sql)
// 	return
// }

// func InsertBySql(sql string) error {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return common.ErrDBNotFound
// 	}

// 	return db.Exec(sql).Error
// }

// func GetOrderById(id string) (result Order) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	db.Table("order").Where("id = ?", id).First(&result)
// 	return
// }

// func DelOrderById(id string) error {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return common.ErrDBNotFound
// 	}

// 	return db.Exec("DELETE FROM `mrs`.`order` WHERE `id` = ?", id).Error
// }
