package controllers

import (
    "github.com/astaxie/beego"
    "github.com/json-iterator/go"

    "ccgwf/base"
    Error "ccgwf/error"

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
 *        "id": "1",
 *        "name": "味千拉面",
 *        "distance": "500",
 *        "star": "4.5",
 *      }
 *    ]
 *  }
**/
func (e *ExampleController) ShopList() {

    e.Initialize(e.Ctx)
    e.Output(Error.SUCCESS, models.GetShops())
    return
}

type ShopAddRequest struct {
    UID      string  `json:"uid"`
    Name     string  `json:"name"`
    Distance int64   `json:"distance"`
    Star     float32 `json:"star"`
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
 *   "distance": "500",
 *   "star": "4.5",
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
 *     "sid1",
 *     "sid2",
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
 *      "vid": "VID"
 *    }
 *  }
 */
func (e *ExampleController) VoteAdd() {
    e.Initialize(e.Ctx)
    e.Output()
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
 *          "id": "1",
 *          "name": "味千拉面",
 *          "distance": "500",
 *          "star": "4.5",
 *        },
 *        {
 *          "id": "2",
 *          "name": "小杨生煎",
 *          "distance": "300",
 *          "star": "5.0",
 *        }
 *      ]
 *    }
 *  }
 */
func (e *ExampleController) VoteGet() {
    e.Initialize(e.Ctx)
    e.Output()
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
 *   "vid": "vid",
 *   "sid": "sid"
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
func (e *ExampleController) VotePost() {
    e.Initialize(e.Ctx)
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
 *        "sid1": 10,
 *        "sid2": 1,
 *      ]
 *    }
 *  }
 */
func (e *ExampleController) VoteResult() {
    e.Initialize(e.Ctx)
    e.Output()
    return
}
