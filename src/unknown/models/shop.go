package models

import (
    "fmt"
    "time"

    "ccgwf/app"
)

type Shop struct {
    ID        int64     `json:"id"`
    UID       string    `json:"uid"`
    Name      string    `json:"name"`
    Distance  int64     `json:"distance"`
    Star      float32   `json:"star"`
    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
}

func GetShops() (shops []Shop) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return
    }

    db.Find(&shops)
    return
}

func AddShop(shop *Shop) (int64, error) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return 0, fmt.Errorf("GetDBClient error")
    }

    err := db.Create(&shop).Error
    if err != nil {
        return 0, fmt.Errorf("AddShopetDb err: %v", err)
    }

    return shop.ID, nil
}
