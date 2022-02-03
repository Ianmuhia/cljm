package services

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"io"
	"io/ioutil"
	"log"
	tasks_client "maranatha_web/tasks/client"
	"net/http"
	"time"
)

type DailVerseService interface {
	ScheduleDailyVerse(verse CreateDailyVerseReq) error
}

type dailyVerseService struct{}

func NewDailyVerseService() DailVerseService {
	return &dailyVerseService{}
}

type CreateDailyVerseReq struct {
	Book    string `json:"book,omitempty"`
	Chapter string `json:"chapter,omitempty"`
	Verse   string `json:"verse,omitempty"`
}

func (bs *dailyVerseService) ScheduleDailyVerse(verse CreateDailyVerseReq) error {
	log.Println("send message has been called.")
	marshal, err := json.Marshal(verse)
	if err != nil {
		return err
	}
	//return

	t1 := asynq.NewTask(TypeDailyVerse, marshal)
	info, err := tasks_client.TasksClient.Enqueue(t1, asynq.ProcessIn(2*time.Minute))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(" [*] Successfully enqueued task: %+v", info)
	return nil
}

func HandleDailyVerseTask(ctx context.Context, t *asynq.Task) error {
	var v CreateDailyVerseReq

	if err := json.Unmarshal(t.Payload(), &v); err != nil {
		return err
	}
	sp := "%20"
	url := fmt.Sprintf("https://bible-api.com/%v%v%v:%v", v.Book, sp, v.Chapter, v.Verse)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	//Convert the body to type string
	sb := string(body)
	log.Printf(sb)

	if err != nil {
		log.Println("failing here.")
		log.Println(err)
		return err
	}
	log.Printf(" [*] Schedule daily verse %s", v.Book)
	return nil
}
