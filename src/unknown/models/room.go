package models

// import (
// 	"time"

// 	"ccgwf/app"
// )

// // CREATE TABLE `shops` (
// //   `id` int(11) NOT NULL AUTO_INCREMENT,
// //   `uid` varchar(32) NOT NULL,
// //   `name` varchar(64) NOT NULL DEFAULT '',
// //   `distance` int(11) NOT NULL,
// //   `star` float(5,2) NOT NULL DEFAULT 0,
// //   `updated_at` timestamp NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp,
// //   `created_at` timestamp NOT NULL DEFAULT current_timestamp,
// //   PRIMARY KEY (`id`)
// // ) ENGINE=InnoDB DEFAULT CHARSET=utf8;

// type Room struct {
// 	ID        int       `json:"id"`
// 	UID       string    `json:"uid"`
// 	Distance  int       `json:"distance"`
// 	Star      float32   `json:"name"`
// 	UpdatedAt time.Time `json:"updated_at"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// func GetAllRooms() (result []Room) {
// 	db := app.Kernel().GetDBClient()
// 	if db == nil {
// 		return
// 	}

// 	db.Table("room").Find(&result)
// 	return
// }
