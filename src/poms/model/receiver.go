package model

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Receiver struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ReceiverDto struct {
	ID       int    `json:"id"`
	Location string `json:"location"`
	Type     string `json:"type"`
}

func GetReceivers() []*Receiver {
	var result []*Receiver

	resp, _ := http.Get("http://localhost:4000/api/orgUnits?type=all")

	data := make([]byte, resp.ContentLength)
	resp.Body.Read(data)
	var receiverDtos []*ReceiverDto
	json.Unmarshal(data, &receiverDtos)
	result = convertToReceivers(receiverDtos)

	return result
}

func convertToReceivers(receiverDtos []*ReceiverDto) []*Receiver {
	result := make([]*Receiver, len(receiverDtos))

	for i, r := range receiverDtos {
		result[i] = &Receiver{
			ID:   r.ID,
			Name: fmt.Sprintf("%s %d", r.Type, r.ID)}
	}

	return result
}
