package webhook

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"
// 	"strings"

// 	"github.com/line/line-bot-sdk-go/v7/linebot"
// )

// var bot *linebot.Client

// func main() {
// 	var err error
// 	// ChannelSecret := os.Getenv("ChannelSecret")
// 	ChannelSecret := "d8d8ea9a7ec6968984e4cc95f994e48c"

// 	ChannelAccessToken := "sGcSiTRiZGfvTfqzuoephLDrKdUwsnZ+aeVYTVV7vUPO1/T/o1IGyWT5UosPC3NcO9iEpKYXZy5ReyLUvKGdD1+EzooACw7nQhcNiXUpP7m/85MGZWG9+G5g1vT6BY6vD+V3A4FrRWNpSjCS/1dImwdB04t89/1O/w1cDnyilFU="
// 	fmt.Println(ChannelSecret)
// 	// ChannelAccessToken := os.Getenv("ChannelAccessToken")
// 	// Port := os.Getenv("Port")
// 	Port := "8080"
// 	bot, err = linebot.New(ChannelSecret, ChannelAccessToken)
// 	log.Println("Bot:", bot, " err:", err)
// 	http.HandleFunc("/callback", callbackHandler)
// 	addr := fmt.Sprintf(":%s", Port)
// 	http.ListenAndServe(addr, nil)
// }

// func callbackHandler(w http.ResponseWriter, r *http.Request) {
// 	events, err := bot.ParseRequest(r)

// 	if err != nil {
// 		if err == linebot.ErrInvalidSignature {
// 			w.WriteHeader(400)
// 		} else {
// 			w.WriteHeader(500)
// 		}
// 		return
// 	}

// 	for _, event := range events {
// 		if event.Type == linebot.EventTypeMessage {
// 			switch message := event.Message.(type) {
// 			// Handle only on text message
// 			case *linebot.TextMessage:
// 				// GetMessageQuota: Get how many remain free tier push message quota you still have this month. (maximum 500)
// 				quota, err := bot.GetMessageQuota().Do()
// 				if err != nil {
// 					log.Println("Quota err:", err)
// 				}
// 				// message.ID: Msg unique ID
// 				// message.Text: Msg text
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("msg ID:"+message.ID+":"+"Get:"+message.Text+" , \n OK! remain message:"+strconv.FormatInt(quota.Value, 10))).Do(); err != nil {
// 					log.Print(err)
// 				}

// 			// Handle only on Sticker message
// 			case *linebot.StickerMessage:
// 				var kw string
// 				for _, k := range message.Keywords {
// 					kw = kw + "," + k
// 				}

// 				outStickerResult := fmt.Sprintf("收到貼圖訊息: %s, pkg: %s kw: %s  text: %s", message.StickerID, message.PackageID, kw, message.Text)
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(outStickerResult)).Do(); err != nil {
// 					log.Print(err)
// 				}
// 			case *linebot.FileMessage:
// 				resp, err := bot.GetMessageContent(message.ID).Do()
// 				if err != nil {
// 					log.Print(err)
// 				}
// 				r1 := resp.Content

// 				filename := message.FileName
// 				fileNameSlice := strings.Split(filename, ".")
// 				var fileNameExtension string
// 				if len(fileNameSlice) > 1 {
// 					fileNameExtension = fileNameSlice[len(fileNameSlice)-1]
// 				}
// 				fmt.Println("file name extension: ", fileNameExtension)
// 				f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
// 				if err != nil {
// 					log.Print(err)
// 				}
// 				defer f.Close()
// 				io.Copy(f, r1)
// 				if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("we get your data")).Do(); err != nil {
// 					log.Print(err)
// 				}
// 			}
// 		}
// 	}
// }

import (
	"fmt"
	"log"
	"net/http"

	// cloudkms "cloud.google.com/go/kms/apiv1"
	// kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	ChannelSecret := "d8d8ea9a7ec6968984e4cc95f994e48c"
	ChannelAccessToken := "sGcSiTRiZGfvTfqzuoephLDrKdUwsnZ+aeVYTVV7vUPO1/T/o1IGyWT5UosPC3NcO9iEpKYXZy5ReyLUvKGdD1+EzooACw7nQhcNiXUpP7m/85MGZWG9+G5g1vT6BY6vD+V3A4FrRWNpSjCS/1dImwdB04t89/1O/w1cDnyilFU="
	client, err := linebot.New(ChannelSecret, ChannelAccessToken)
	if err != nil {
		http.Error(w, "Error init client", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	events, err := client.ParseRequest(r)
	if err != nil {
		http.Error(w, "Error parse request", http.StatusBadRequest)
		log.Fatal(err)
		return
	}
	for _, e := range events {
		switch e.Type {
		case linebot.EventTypeMessage:
			message := linebot.NewTextMessage("Test")
			// invoke add notion service
			_, err := client.ReplyMessage(e.ReplyToken, message).Do()
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
	fmt.Fprint(w, "ok")
}
