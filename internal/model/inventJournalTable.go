package model

import (
	"Axsprav/internal/config"
	"errors"
	"log"
	"time"
)

type StoryToken struct {
	GUID          string    `db:"GUIDTOKEN"`
	GUIDDATETIME  time.Time `db:"GUIDDATETIME"`
	JOURNALID     string    `db:"JOURNALID"`
	TYPEOPERATION string    `db:"TYPEOPERATION"`
	APPROVED      int       `db:"APPROVED"`
}

type DayOfMail struct {
	DayOfMail int `db:"DAYOFMAIL"`
}

func GetCREATEDDATETIMEByToken(token string) (StoryToken, error) {

	config.Config.LOGGER.Info(token)
	storyToken := StoryToken{}
	err := config.Config.DB.Get(&storyToken, "SELECT GUIDTOKEN, GUIDDATETIME, JOURNALID, TYPEOPERATION, APPROVED, JournalIdCompTrans  from [AxSprav].[dbo].ZSIGNATUREHISTORYTOKEN where GUIDTOKEN=@p1", token)

	if err != nil {
		return StoryToken{}, err
	}

	return storyToken, nil
}

func GetDayOfMail() (DayOfMail, error) {

	dayOfMail := DayOfMail{}
	err := config.Config.DB.Get(&dayOfMail, "select DAYOFMAIL from [AxSprav].[dbo].InventParameters where DATAAREAID = 'ref'")

	if err != nil {
		return DayOfMail{}, err
	}

	return dayOfMail, nil
}

func UpdateInventJournal(storyToken StoryToken) (int, error) {

	log.Println(storyToken.TYPEOPERATION)

	if storyToken.TYPEOPERATION != "receive" && storyToken.TYPEOPERATION != "transfer" {
		return 0, errors.New("undefined type in TYPEOPERATION")
	}

	if storyToken.TYPEOPERATION == "receive" {
		_, err := config.Config.DB.Exec("UPDATE  [AxSprav].[dbo].[INVENTJOURNALTABLE] SET   [JOURNALCHEKMOLTO] = 1 where [AxSprav].[dbo].[INVENTJOURNALTABLE].JOURNALID =@p1", storyToken.JOURNALID)

		if err != nil {
			return 0, err
		}

	}

	if storyToken.TYPEOPERATION == "transfer" {
		_, err := config.Config.DB.Exec("UPDATE  [AxSprav].[dbo].[INVENTJOURNALTABLE] SET   [JOURNALCHEKMOLFROM] = 1 where [AxSprav].[dbo].[INVENTJOURNALTABLE].JOURNALID =@p1", storyToken.JOURNALID)

		if err != nil {
			return 0, err
		}
	}

	_, err := config.Config.DB.Exec("UPDATE  [AxSprav].[dbo].ZSIGNATUREHISTORYTOKEN SET  [APPROVEDDATETIME] = @p1, [APPROVED] = 1, [OperationStatus] = 'Успешно' where GUIDTOKEN=@p2", time.Now(), storyToken.GUID)

	if err != nil {
		return 0, err
	}

	return 200, nil
}
