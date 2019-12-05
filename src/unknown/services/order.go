// package services

// import (
// 	"fmt"
// 	"strings"

// 	BaseApp "ccgwf/app"
// 	"ccgwf/open"

// 	"environment/tool"

// 	"mrs/app"
// 	"mrs/common"
// 	"mrs/models"
// )

// func OrderRoom(params common.OrderRequest) error {
// 	redisCli := BaseApp.Kernel().GetRedisClient()
// 	sql := "INSERT INTO `mrs`.`order` (`room_id`, `date`, `block`, `uid`, `username`, `avatar`, `subject`, `remark`) VALUES "
// 	var vals []string
// 	keys := map[string]int64{}
// 	for _, v := range params.Orders {
// 		if redisCli != nil {
// 			// 检查缓存，防止恶意重复预定
// 			cacheKey := fmt.Sprintf(common.OrderCacheKey, params.Date, v.Id, v.Block, params.Uid)
// 			keys[cacheKey] = 0
// 			val, redisErr := redisCli.GetOne(cacheKey)
// 			if redisErr == nil {
// 				intVal := tool.StringToInt64(val)
// 				keys[cacheKey] = intVal
// 				if intVal >= app.Kernel().GetConfig().OrderLimit {
// 					return common.ErrOrderLimit
// 				}
// 			}
// 		}

// 		vals = append(vals, fmt.Sprintf("('%v', '%v', '%v', '%v', '%v', '%v', '%v', '%v')", v.Id, params.Date, v.Block, params.Uid, params.Name, params.Avatar, params.Subject, params.Remark))
// 	}
// 	sql = sql + strings.Join(vals, ", ")
// 	if err := models.InsertBySql(sql); err != nil {
// 		return err
// 	}

// 	if redisCli != nil { // 预定成功，记录缓存
// 		for k, _ := range keys {
// 			second, ttlErr := redisCli.TTL(k)
// 			if ttlErr == nil {
// 				if second > 0 {
// 					afterVal, incrErr := redisCli.Incr(k)
// 					if incrErr == nil {
// 						keys[k] = afterVal
// 					}
// 				} else {
// 					setResult, setErr := redisCli.Set(k, "1", app.Kernel().GetConfig().OrderLimitExprie)
// 					fmt.Println("OrderRoom redis Set", setResult, setErr)
// 				}
// 			}
// 		}
// 	}

// 	// 预定成功发射消息到hooks
// 	if app.Kernel().GetConfig().GroupHook != "" {
// 		msgStruct := &open.MsgAttachment{
// 			Type: open.MSG_TYPE_ATTACHMENT,
// 			Message: &open.DataAttachment{
// 				Head: &open.DataAttachmentHead{
// 					Text: "预定信息",
// 				},
// 			},
// 		}
// 		msgFieldInputSubject := &open.DataAttachmentField{
// 			Name: "主题",
// 			Text: params.Subject,
// 		}
// 		msgStruct.Message.Fields = append(msgStruct.Message.Fields, msgFieldInputSubject)

// 		msgFieldInputDate := &open.DataAttachmentField{
// 			Name: "日期",
// 			Text: params.Date,
// 		}
// 		msgStruct.Message.Fields = append(msgStruct.Message.Fields, msgFieldInputDate)

// 		msgFieldInputName := &open.DataAttachmentField{
// 			Name: "预订人",
// 			Text: params.Name,
// 		}
// 		msgStruct.Message.Fields = append(msgStruct.Message.Fields, msgFieldInputName)

// 		classify := common.RoomClassify(params.Orders)
// 		for _, v := range classify {
// 			msgFieldInputTmp := &open.DataAttachmentField{
// 				Name: v["name"],
// 				Input: &open.DataAttachmentFieldInput{
// 					Type:      open.INPUT_TYPE_TEXT,
// 					Childtype: open.INPUT_CHILDTYPE_TEXT_STRING,
// 					Color:     open.COLOR_BLUE,
// 					Text:      v["text"],
// 				},
// 			}
// 			msgStruct.Message.Fields = append(msgStruct.Message.Fields, msgFieldInputTmp)
// 		}

// 		openPlatform := open.New("", "", "")
// 		sendRes, sendErr := openPlatform.MsgService(app.Kernel().GetConfig().GroupHook, msgStruct)
// 		fmt.Println("MsgService sendRes", sendRes, sendErr)
// 	}

// 	return nil
// }
