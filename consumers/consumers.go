package consumers

import (
	"time"

	"github.com/sensorsdata/sa-sdk-go/structs"
	"github.com/sensorsdata/sa-sdk-go/utils"
)

type Consumer interface {
	Send(data structs.EventData) error
	Flush() error
	Close() error
	ItemSend(item structs.Item) error
}

func send(url string, data string, to time.Duration, list bool) error {
	pdata := ""

	if list {
		rdata, err := utils.GeneratePostDataList(data)
		if err != nil {
			return err
		}
		pdata = rdata
	} else {
		rdata, err := utils.GeneratePostData(data)
		if err != nil {
			return err
		}
		pdata = rdata
	}

	err := utils.DoRequest(url, pdata, to)
	if err != nil {
		return err
	}

	return nil
}
