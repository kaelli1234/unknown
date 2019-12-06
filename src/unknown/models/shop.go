package models

import (
    "time"

    "github.com/pkg/errors"

    "ccgwf/app"
)

var (
    errDBNotFound = errors.New("DB instance not init!")
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

    db.Order("id DESC").Limit(20).Find(&shops)
    return
}

func AddShop(shop *Shop) (int64, error) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return 0, errDBNotFound
    }

    err := db.Create(&shop).Error
    if err != nil {
        return 0, errors.WithStack(err)
    }

    return shop.ID, nil
}

func GetShopsByIDs(ids []int64) (shops []Shop) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return
    }

    db.Where(ids).Find(&shops)
    return
}
