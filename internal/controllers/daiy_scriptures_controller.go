package controllers

import (
	"maranatha_web/internal/services"
	"maranatha_web/internal/utils/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDailyVerseReq struct {
	Book    string `json:"book,omitempty"`
	Chapter string `json:"chapter,omitempty"`
	Verse   string `json:"verse,omitempty"`
}

func (r *Repository) GetDailyScriptures(ctx *gin.Context) {
	var req CreateDailyVerseReq

	if err := ctx.ShouldBindJSON(&req); err != nil {

		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.Status, restErr)
		ctx.Abort()
		return
	}
	type dvr services.CreateDailyVerseReq
	data := dvr{
		Book:    req.Book,
		Chapter: req.Chapter,
		Verse:   req.Verse,
	}

	err := r.dailyVerse.ScheduleDailyVerse(services.CreateDailyVerseReq(data))
	if err != nil {
		return
	}
	//err := services.MailService.SendMsg(services.Mail(m))
	////log.Println(&dc)
	//
	//if err != nil {
	//	log.Println(err)
	//	//logger.Info("could not send email ")
	//	return
	//}
	////sp := "%20"
	//url := fmt.Sprintf("https://bible-api.com/%v%v%v:%v", req.Book, sp, req.Chapter, req.Verse)
	//log.Println(url)
	//resp, err := http.Get(url)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		log.Println(err)
	//	}
	//}(resp.Body)
	////We Read the response body on the line below.
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	////Convert the body to type string
	//sb := string(body)
	//log.Printf(sb)

	ctx.JSON(http.StatusOK, "ok")

}
