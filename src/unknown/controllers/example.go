package controllers

import (
    "github.com/astaxie/beego"
    "github.com/json-iterator/go"

    "environment/tool"

    "ccgwf/base"
    Error "ccgwf/error"
    "ccgwf/logs"

    "unknown/models"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ExampleController struct {
    beego.Controller
    base.ClientController
}

/**
 * @api {get} /shops 获取店铺列表
 * @apiName 获取店铺列表
 * @apiGroup Shop
 *
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *    "datas": [
 *      {
 *        "id": 1,
 *        "name": "味千拉面",
 *        "distance": 500,
 *        "star": 4.5,
 *      }
 *    ]
 *  }
**/
func (e *ExampleController) ShopList() {

    e.Initialize(e.Ctx)
    e.Output(Error.SUCCESS, models.GetShops())
    return
}

/**
 * @api {post} /shops 新增店铺
 * @apiName 新增店铺
 * @apiGroup Shop
 *
 * @apiSuccessExample Request-Example:
 * Content-Type: application/json
 * {
 *   "uid": "UID",
 *   "name": "味千拉面",
 *   "distance": 500,
 *   "star": 4.5,
 * }
 *
 * @apiSuccessExample Success-Response:
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *  }
 */

type ShopAddRequest struct {
    UID      string  `json:"uid"`
    Name     string  `json:"name"`
    Distance int64   `json:"distance"`
    Star     float32 `json:"star"`
}

func (e *ExampleController) ShopAdd() {

    e.Initialize(e.Ctx)

    var request ShopAddRequest
    if err := json.Unmarshal(e.Ctx.Input.RequestBody, &request); err != nil {
        e.Output(Error.PARAM_ERROR, err.Error())
        return
    }

    _, err := models.AddShop(&models.Shop{
        UID:      request.UID,
        Name:     request.Name,
        Distance: request.Distance,
        Star:     request.Star,
    })
    if err != nil {
        e.Output(Error.SERVER_INTERVAL_ERROR, err.Error())
        return
    }

    e.Output()
    return
}

/**
 * @api {post} /votes 发起投票
 * @apiName 发起投票
 * @apiGroup Vote
 *
 * @apiSuccessExample Request-Example:
 * Content-Type: application/json
 * {
 *   "uid": "UID",
 *   "subject": "今天吃什么",
 *   "shops": [
 *     sid1,
 *     sid2,
 *   ],
 * }
 *
 * @apiSuccessExample Success-Response:
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *    "datas": {
 *      "vid": VID
 *    }
 *  }
 */

type VoteAddRequest struct {
    UID     string  `json:"uid"`
    Subject string  `json:"subject"`
    Shops   []int64 `json:"shops"`
}

func (e *ExampleController) VoteAdd() {

    e.Initialize(e.Ctx)

    var request VoteAddRequest
    if err := json.Unmarshal(e.Ctx.Input.RequestBody, &request); err != nil {
        e.Output(Error.PARAM_ERROR, err.Error())
        return
    }

    vid, err := models.AddVote(&models.Vote{
        UID:     request.UID,
        Subject: request.Subject,
    })
    if err != nil {
        e.Output(Error.SERVER_INTERVAL_ERROR, err.Error())
        return
    }

    voteOptions := make([]*models.VoteOption, len(request.Shops))
    for k, v := range request.Shops {
        voteOptions[k] = &models.VoteOption{
            VID: vid,
            SID: v,
        }
    }

    if err := models.AddVoteOptions(voteOptions); err != nil {
        e.Output(Error.SERVER_INTERVAL_ERROR, err.Error())
        return
    }

    e.Output(Error.SUCCESS, map[string]int64{"vid": vid})
    return
}

/**
 * @api {get} /votes/:id 获取投票详情
 * @apiName 获取投票详情
 * @apiGroup Vote
 *
 * @apiSuccessExample Success-Response:
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *    "datas": {
 *      "subject": "今天吃什么",
 *      "shops": [
 *        {
 *          "id": 1,
 *          "name": "味千拉面",
 *          "distance": 500,
 *          "star": 4.5,
 *        },
 *        {
 *          "id": 2,
 *          "name": "小杨生煎",
 *          "distance": 300,
 *          "star": 5.0,
 *        }
 *      ]
 *    }
 *  }
 */

type VoteGetResponse struct {
    Subject string        `json:"subject"`
    Shops   []models.Shop `json:"shops"`
}

func (e *ExampleController) VoteGet() {

    e.Initialize(e.Ctx)

    id := e.Ctx.Input.Param(":id")

    options := models.GetVoteOptionsByVID(tool.StringToInt64(id))

    sids := make([]int64, len(options))
    for k, v := range options {
        sids[k] = v.SID
    }
    logs.Debug("sids %v", sids)

    e.Output(Error.SUCCESS, &VoteGetResponse{
        Shops: models.GetShopsByIDs(sids),
    })
    return
}

/**
 * @api {post} /votes/post 投票
 * @apiName 投票
 * @apiGroup Vote
 *
 * @apiSuccessExample Request-Example:
 * Content-Type: application/json
 * {
 *   "uid": "UID",
 *   "vid": vid,
 *   "sid": sid
 * }
 *
 * @apiSuccessExample Success-Response:
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *  }
 */

type VotePostRequest struct {
    UID string `json:"uid"`
    VID int64  `json:"vid"`
    SID int64  `json:"sid"`
}

func (e *ExampleController) VotePost() {

    e.Initialize(e.Ctx)

    var request VotePostRequest
    if err := json.Unmarshal(e.Ctx.Input.RequestBody, &request); err != nil {
        e.Output(Error.PARAM_ERROR, err.Error())
        return
    }

    _, err := models.AddVoteResult(&models.VoteResult{
        VID: request.VID,
        SID: request.SID,
        UID: request.UID,
    })
    if err != nil {
        e.Output(Error.SERVER_INTERVAL_ERROR, err.Error())
        return
    }

    e.Output()
    return
}

/**
 * @api {get} /votes/:id/result 获取投票结果
 * @apiName 获取投票结果
 * @apiGroup Result
 *
 * @apiSuccessExample Success-Response:
 *  {
 *    "trans_id": "01f144d839c50a4ec8256afeccf5887a",
 *    "err_code": "0",
 *    "err_msg": "成功",
 *    "timestamp": "1575518905",
 *    "time_cost": "0.00004",
 *    "datas": {
 *      "name": "今天吃什么",
 *      "shops": [
 *        {
 *          "id": 1,
 *          "name": "味千拉面",
 *          "distance": 500,
 *          "star": 4.5,
 *          "total": 1,
 *        },
 *        {
 *          "id": 2,
 *          "name": "小杨生煎",
 *          "distance": 300,
 *          "star": 5.0,
 *          "total": 2,
 *        }
 *      ]
 *    }
 *  }
 */

type VoteResultResponse struct {
    Name  string             `json:"name"`
    Shops []*VoteResultShops `json:"shops"`
}

type VoteResultShops struct {
    ID       int64   `json:"id"`
    Name     string  `json:"name"`
    Distance int64   `json:"distance"`
    Star     float32 `json:"star"`
    Total    int64   `json:"total"`
}

func (e *ExampleController) VoteResult() {

    e.Initialize(e.Ctx)

    id := e.Ctx.Input.Param(":id")

    results := models.GetVoteResultGroupByVID(tool.StringToInt64(id))
    logs.Debug("results %+v", results)

    sids := make([]int64, len(results))
    sidsMap := make(map[int64]int64, len(results))
    for k, v := range results {
        sids[k] = v.SID
        sidsMap[v.SID] = v.Total
    }

    logs.Debug("sidsMap %+v", sidsMap)

    shops := models.GetShopsByIDs(sids)
    datas := make([]*VoteResultShops, len(shops))
    for k, v := range shops {
        datas[k] = &VoteResultShops{
            ID:       v.ID,
            Name:     v.Name,
            Distance: v.Distance,
            Star:     v.Star,
            Total:    sidsMap[v.ID],
        }
    }

    e.Output(Error.SUCCESS, datas)
    return
}
