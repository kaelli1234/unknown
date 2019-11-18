package controllers

import (
	"github.com/astaxie/beego"
	"github.com/json-iterator/go"

	"ccgwf/base"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type ExampleController struct {
	beego.Controller
	base.ClientController
}

/**
 * @api {get} /user/:id Request User information
 * @apiName GetUser
 * @apiGroup User
 *
 * @apiParam {Number} id Users unique ID.
 *
 * @apiSuccess {String} firstname Firstname of the User.
 * @apiSuccess {String} lastname  Lastname of the User.
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "firstname": "John",
 *       "lastname": "Doe"
 *     }
 *
 * @apiError UserNotFound The id of the User was not found.
 *
 * @apiErrorExample Error-Response:
 *     HTTP/1.1 404 Not Found
 *     {
 *       "error": "UserNotFound"
 *     }
 */
func (e *ExampleController) Home() {
	e.Initialize(e.Ctx)
	e.Output()
	return
}
