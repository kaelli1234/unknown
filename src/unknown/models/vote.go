package models

import (
    "time"

    "github.com/pkg/errors"

    "ccgwf/app"
)

type Vote struct {
    ID        int64     `json:"id"`
    UID       string    `json:"uid"`
    Subject   string    `json:"subject"`
    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
}

type VoteOption struct {
    ID        int64     `json:"id"`
    VID       int64     `gorm:"column:vid" json:"vid"`
    SID       int64     `gorm:"column:sid" json:"sid"`
    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
}

type VoteResult struct {
    ID        int64     `json:"id"`
    VID       int64     `gorm:"column:vid" json:"vid"`
    SID       int64     `gorm:"column:sid" json:"sid"`
    UID       string    `json:"uid"`
    UpdatedAt time.Time `json:"updated_at"`
    CreatedAt time.Time `json:"created_at"`
}

type GroupByVidVoteResult struct {
    SID   int64 `gorm:"column:sid" json:"sid"`
    Total int64 `json:"total"`
}

func AddVote(vote *Vote) (int64, error) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return 0, errDBNotFound
    }

    err := db.Create(&vote).Error
    if err != nil {
        return 0, errors.WithStack(err)
    }

    return vote.ID, nil
}

func AddVoteOptions(voteOptions []*VoteOption) error {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return errDBNotFound
    }

    tx := db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    if err := tx.Error; err != nil {
        return errors.WithStack(err)
    }

    for _, v := range voteOptions {
        if err := tx.Create(&v).Error; err != nil {
            tx.Rollback()
            return errors.WithStack(err)
        }
    }

    return tx.Commit().Error
}

func GetVoteOptionsByVID(vid int64) (voteOptions []VoteOption) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return
    }

    db.Where("vid = ?", vid).Find(&voteOptions)
    return
}

func AddVoteResult(voteResult *VoteResult) (int64, error) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return 0, errDBNotFound
    }

    err := db.Create(&voteResult).Error
    if err != nil {
        return 0, errors.WithStack(err)
    }

    return voteResult.ID, nil
}

func GetVoteResultGroupByVID(vid int64) (result []GroupByVidVoteResult) {

    db := app.Kernel().GetDBClient()
    if db == nil {
        return
    }

    db.Table("vote_results").Where("vid = ?", vid).Select("sid, count(sid) AS total").Group("sid").Scan(&result)
    return
}
