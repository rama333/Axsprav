package controller

import (
	"Axsprav/internal/config"
	"Axsprav/internal/model"
	"Axsprav/internal/restapi/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"time"
)

func (c *Controller) UpdateInventJournalTable(ctx *gin.Context) {


	slogger := config.Config.LOGGER

	_, err := uuid.Parse(ctx.Query("token"))
	if err != nil{
		response.ResponseBadRequest("Неправильный формат токена", ctx)
		return
	}

	val, err := model.GetCREATEDDATETIMEByToken(ctx.Query("token"))

	if err != nil{
		slogger.Error(err)
		response.ResponseInternalServerError("Данный токен не существует", ctx)
		return
	}

	if val.GUID == "" {
		slogger.Info(err)
		response.ResponseInternalServerError("Данный токен не найден", ctx)
		return
	}

	t, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil{
		slogger.Error(err)
		response.ResponseInternalServerError("Внутренняя ошибка сервера", ctx)
		return
	}


	if val.APPROVED==1 {
		slogger.Info("Данный % журнал был уже утвержден ранее.")
		response.ResponseOkRequest("Журнал №"+ val.JOURNALID + " был уже утвержден ранее.", ctx)
		return
	}

	if t.Sub(val.GUIDDATETIME).Hours() > 24  {
		slogger.Info("Время действия ссылки для проставления подписи журнала №"+ val.JOURNALID + " истекло. Запросите новое подтверждение подписи.")
		response.ResponseOkRequest("Время действия ссылки для проставления подписи журнала №"+ val.JOURNALID + " истекло. Запросите новое подтверждение подписи.", ctx)
		return
	}

	code, err := model.UpdateInventJournal(val)

	if err != nil {
		slogger.Error(err)
		response.ResponseInternalServerError("Внутренняя ошибка сервера", ctx)
		return
	}

	if code == 200{
		response.ResponseOkRequest("Журнал №"+ val.JOURNALID + " успешно подтвержден", ctx)
	}

	//if err != nil {
	//	slogger.Error(err, des)
	//	restapi.ResponseInternalServerError("internal server error", ctx)
	//	return
	//}
	//
	//if code == 0 {
	//	restapi.ResponseOkRequest("ok", ctx)
	//
	//} else {
	//	restapi.ResponseOkRequest("session not active", ctx)
	//}
}