package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mkishere/beancount/beancount"
)

var (
	apiKey string
	cfg    *aws.Config
	bot    *tgbotapi.BotAPI
	db     *dynamodb.DynamoDB
	tz     *time.Location
)

const ()

func init() {
	apiKey = os.Getenv("TELEGRAM_API_KEY")
	var err error
	tz, err = time.LoadLocation("Asia/Hong_Kong")
	if err != nil {
		log.Println("Cannot load timezone information, revert to use local time")
		tz = time.Local
	}
}

func main() {
	lambda.Start(handleTelegramMsg)
}

func handleTelegramMsg(update tgbotapi.Update) {
	if cfg == nil {
		var err error
		cfgVal, err := external.LoadDefaultAWSConfig()
		if err != nil {
			log.Fatalf("failed to load config, %v", err)
		}
		cfg = &cfgVal
	}

	if bot == nil {
		var err error
		bot, err = tgbotapi.NewBotAPI(apiKey)
		if err != nil {
			log.Fatal(err)
		}
	}

	if db == nil {
		db = dynamodb.New(*cfg)
	}

	if update.CallbackQuery != nil {
		cbQuery := update.CallbackQuery
		data := strings.Split(cbQuery.Data, " ")
		chatID, _ := strconv.ParseInt(data[0], 10, 64)
		newBal, err := beancount.UpdateBalance(db, chatID, data[1])
		if err != nil {
			log.Fatal(err)
		}
		err = SendMsg(chatID, fmt.Sprintf("@%v reverted entry.\nBalance updated. New balance %v", cbQuery.From.UserName, newBal))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	chatID := update.Message.Chat.ID

	if update.Message.ReplyToMessage != nil {
		amt := update.Message.Text
		if strings.ToLower(strings.TrimSpace(amt)) == "cancel" {
			return
		}
		// Trim text after space if there's one
		spaceIdx := strings.Index(amt, " ")
		if spaceIdx >= 0 {
			amt = amt[:spaceIdx]
		}
		if _, err := strconv.ParseFloat(amt, 32); len(amt) == 0 || err != nil {
			msg := tgbotapi.NewMessage(chatID, "Don't know what you're talking. Try again, or say \"cancel\" to cancel.")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			bot.Send(msg)
			return
		}
		newBal, err := beancount.UpdateBalance(db, chatID, amt)
		if err != nil {
			log.Fatal(err)
		}
		err = SendMsg(chatID, fmt.Sprintf("Balance updated. New balance %v", newBal))
		if err != nil {
			log.Fatal(err)
		}
	}
	// Split command and argument
	spaceIdx := strings.Index(update.Message.Text, " ")
	var cmdText string
	var args string
	if spaceIdx >= 0 {
		cmdText = update.Message.Text[:spaceIdx]
		args = update.Message.Text[spaceIdx+1:]
	} else {
		cmdText = update.Message.Text
	}
	// Trim the @botName out
	atIdx := strings.Index(cmdText, "@")
	if atIdx >= 0 {
		cmdText = cmdText[:atIdx]
	}
	switch cmdText {
	case "/add":
		if len(args) == 0 {
			// Force reply from user if no argument detected
			msg := tgbotapi.NewMessage(chatID, "How much?")
			msg.ReplyToMessageID = update.Message.MessageID
			msg.ReplyMarkup = tgbotapi.ForceReply{
				ForceReply: true,
				Selective:  true,
			}
			bot.Send(msg)
			return
		}
		// Trim everything after space
		spaceIdx = strings.Index(args, " ")
		if spaceIdx >= 0 {
			args = args[:spaceIdx]
		}
		if _, err := strconv.ParseFloat(args, 32); err != nil {
			err = SendMsg(chatID, "Don't understand, please try again.")
			if err != nil {
				log.Fatal(err)
			}
		}
		newBal, err := beancount.UpdateBalance(db, chatID, args)
		if err != nil {
			log.Fatal(err)
		}
		err = SendMsg(chatID, fmt.Sprintf("Balance updated. New balance %v", newBal))
		if err != nil {
			log.Fatal(err)
		}
	case "/list":
		cnt, _ := strconv.Atoi(args)
		if cnt == 0 {
			cnt = 10
		}
		hist, err := beancount.GetTxHist(db, chatID, cnt)
		if err != nil {
			log.Fatal(err)
		}
		PrintHist(chatID, hist)
	case "/balance":
		bal, err := beancount.GetBalance(db, chatID)
		if err != nil {
			log.Fatal(err)
		}
		err = SendMsg(chatID, fmt.Sprintf("Current balance: %v", bal))
		if err != nil {
			log.Fatal(err)
		}
	case "/reset":
		newBal := "0"
		if _, err := strconv.ParseFloat(args, 32); err == nil {
			newBal = args
		}
		err := beancount.ResetBalance(db, chatID, newBal)
		if err != nil {
			log.Fatal(err)
		}
		err = SendMsg(chatID, "Balance reset")
		if err != nil {
			log.Fatal(err)
		}
	}
}

// SendMsg sends simple telegram message back to user
func SendMsg(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// PrintHist display transaction history in chat, along
// with the reverse entry button
func PrintHist(chatID int64, hist []beancount.Entry) {
	for _, entry := range hist {
		ts := time.Unix(entry.Timestamp, -1)
		msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("%v @ %v", entry.Amount, ts.In(tz).Format("02/01/2006 15:04:05")))
		kbMarkup := tgbotapi.NewInlineKeyboardMarkup(
			[]tgbotapi.InlineKeyboardButton{
				tgbotapi.NewInlineKeyboardButtonData("⏪", fmt.Sprintf("%v %v", chatID, negate(entry.Amount))),
			})
		msg.ReplyMarkup = kbMarkup
		bot.Send(msg)
	}
}

// negate returns the
func negate(num string) string {
	if num[0] == '-' {
		return num[1:]
	}
	return "-" + num
}

func decrypt(cfg aws.Config, encrypted string) (string, error) {
	kmsClient := kms.New(cfg)

	decodedBytes, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}
	input := &kms.DecryptInput{
		CiphertextBlob: decodedBytes,
	}
	req := kmsClient.DecryptRequest(input)
	rsp, err := req.Send()
	if err != nil {
		return "", err
	}
	// Plaintext is a byte array, so convert to string
	return string(rsp.Plaintext), nil
}
