package messagequeue

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Cyndir/tucows-app/internal/model"
)

//go:generate mockgen -destination=../mocks/messageProcessor.go -package=mocks -source=messageProcessor.go
type MessageProcessor interface {
	ProcessMessage([]byte)
}

type messageProcessorImpl struct {
}

func NewProcessor() MessageProcessor {
	return &messageProcessorImpl{}
}
func (m messageProcessorImpl) ProcessMessage(body []byte) {
	var order model.Order
	json.Unmarshal(body, &order)
	if order.Total > 1000 {
		order.Status = "payment failed"
	} else {
		order.Status = "payment successful"
	}

	sendResponse(order)
}

func sendResponse(order model.Order) {
	data, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPatch, "http://localhost:8081/order", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}
