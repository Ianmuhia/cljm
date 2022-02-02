package controllers

// import (
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
//

// 	"github.com/gin-gonic/gin"

// 	"maranatha_web/utils/errors"
// )

// type CreateDailyVerseReq struct {
// 	Book    string
// 	Chapter string
// 	Verse   string
// }

// func GetDailyScriptures(ctx *gin.Context) {
// 	var req createUserRequest

// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		restErr := errors.NewBadRequestError("invalid json body")
// 		ctx.JSON(restErr.Status, restErr)
// 		ctx.Abort()
// 		return
// 	}

// 	resp, err := http.Get("https://bible-api.com/john%203:16")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer func(Body io.ReadCloser) {
// 		err := Body.Close()
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}(resp.Body)
// 	//We Read the response body on the line below.
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	//Convert the body to type string
// 	sb := string(body)
// 	log.Printf(sb)

// 	ctx.JSON(http.StatusOK, sb)

// }
