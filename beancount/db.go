package beancount

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
)

// Account contains the ChatID the bot relates to and the balance
type Account struct {
	ChatID  int64
	Balance string
	Hist    []Entry
}

// Entry represents a single transaction
type Entry struct {
	Timestamp int64
	Amount    string
}

const (
	balanceTable = "BeanCount"
	//histTable    = "BeanCount_Trans"
)

// GetBalance returns the stored balance of the provided chat
func GetBalance(db *dynamodb.DynamoDB, key int64) (string, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(balanceTable),
		Key: map[string]dynamodb.AttributeValue{
			"ChatID": {
				N: aws.String(strconv.FormatInt(key, 10)),
			},
		},
	}
	req := db.GetItemRequest(params)
	result, err := req.Send()
	if err != nil {
		return "0", err
	}
	var acc Account
	err = dynamodbattribute.UnmarshalMap(result.Item, &acc)
	if err != nil {
		return "0", err
	}
	return acc.Balance, nil
}

// UpdateBalance sums up the amount provided to the balance
// and record down in history
func UpdateBalance(db *dynamodb.DynamoDB, key int64, amount string) (string, error) {
	entry := Entry{
		Amount:    amount,
		Timestamp: time.Now().Unix(),
	}
	entryMap, err := dynamodbattribute.MarshalMap(&entry)
	if err != nil {
		return "", err
	}
	param := &dynamodb.UpdateItemInput{
		TableName: aws.String(balanceTable),
		Key: map[string]dynamodb.AttributeValue{
			"ChatID": {
				N: aws.String(strconv.FormatInt(key, 10)),
			},
		},
		UpdateExpression: aws.String("SET Balance = Balance + :amount, Hist = list_append(Hist, :entry)"),
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":amount": {
				N: aws.String(amount),
			},
			":entry": {
				L: []dynamodb.AttributeValue{
					dynamodb.AttributeValue{
						M: entryMap,
					},
				},
			},
		},
		ReturnValues: dynamodb.ReturnValueUpdatedNew,
	}
	req := db.UpdateItemRequest(param)
	res, err := req.Send()

	resVal := *res.Attributes["Balance"].N
	if err != nil {
		return "", err
	}
	// Trim tx history into 10
	param = &dynamodb.UpdateItemInput{
		TableName: aws.String(balanceTable),
		Key: map[string]dynamodb.AttributeValue{
			"ChatID": {
				N: aws.String(strconv.FormatInt(key, 10)),
			},
		},
		UpdateExpression:    aws.String("REMOVE Hist[0]"),
		ConditionExpression: aws.String("size (Hist) > :maxLen"),
		ExpressionAttributeValues: map[string]dynamodb.AttributeValue{
			":maxLen": {
				N: aws.String("10"),
			},
		},
		ReturnValues: dynamodb.ReturnValueNone,
	}
	req = db.UpdateItemRequest(param)
	res, _ = req.Send()

	return resVal, err
}

// ResetBalance reset the balance to the newBal
func ResetBalance(db *dynamodb.DynamoDB, key int64, newBal string) error {
	param := &dynamodb.PutItemInput{
		TableName: aws.String(balanceTable),
		Item: map[string]dynamodb.AttributeValue{
			"ChatID": {
				N: aws.String(strconv.FormatInt(key, 10)),
			},
			"Balance": {
				N: aws.String(newBal),
			},
			"Hist": {
				L: []dynamodb.AttributeValue{},
			},
		},
		ReturnValues: dynamodb.ReturnValueNone,
	}
	req := db.PutItemRequest(param)
	_, err := req.Send()
	if err != nil {
		return err
	}
	return nil
}

// GetTxHist returns the transaction history for the record, the number of
// records returns is controlled by the provided count parameter
func GetTxHist(db *dynamodb.DynamoDB, key int64, count int) ([]Entry, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String(balanceTable),
		Key: map[string]dynamodb.AttributeValue{
			"ChatID": {
				N: aws.String(strconv.FormatInt(key, 10)),
			},
		},
		ProjectionExpression: aws.String("Hist"),
	}

	req := db.GetItemRequest(params)
	result, err := req.Send()
	if err != nil {
		return nil, err
	}
	entries := []Entry{}
	res := result.Item["Hist"]
	err = dynamodbattribute.Unmarshal(&res, &entries)
	return entries, err
}
