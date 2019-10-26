package ctrl

import (
	"encoding/json"
	"net/http"
	"poms/model"
	"testing"
)

var captureData []byte

func TestGetCurrencies(t *testing.T) {
	//arrange

	var controller CurrencyController

	var responseWriter mockResponseWriter

	responseWriter.header = make(map[string][]string)

	//act
	controller.GetCurrencies(responseWriter, &http.Request{})

	//assert
	if responseWriter.header.Get("Content-Type") != "application/json" {
		t.Error("Missing or incorrect Content-Type header")
	}

	//sets content-type header

	//writes correct data to client

	currencies := model.GetCurrencies()

	currencyData, _ := json.Marshal(currencies)

	if string(captureData) != string(currencyData) {
		t.Log(string(captureData))
		t.Log(string(currencyData))
		t.Error("incorrect data send to client")
	}

}

type mockResponseWriter struct {
	header http.Header
}

func (mrw mockResponseWriter) Header() http.Header {
	return mrw.header
}

func (mrw mockResponseWriter) Write(data []byte) (int, error) {
	captureData = data[:]
	return len(data), nil
}

func (mrw mockResponseWriter) WriteHeader(code int) {

}
