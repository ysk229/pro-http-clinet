package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/ysk229/pro-http-clinet/pb/proto/api/weather/v1"
	"log"
)

func main() {
	client := resty.New()

	var req = weather.Request{Key: "SCYrvkytJze9qyzOh", Location: "be"}
	err := req.Validate()
	if err != nil {
		fmt.Println("err", err)
	}

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"key":      req.GetKey(),
			"location": req.GetLocation(),
		}).
		SetHeader("Accept", "application/json").
		Get("https://api.seniverse.com/v3/weather/daily.json?language=zh-Hans&unit=c")
	rst := &weather.Reply{}
	err = json.Unmarshal(resp.Body(), rst)
	//err = proto.Unmarshal(resp.Body(),rst )
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	for _, item := range rst.Results {
		//log.Println(item)
		log.Println(item.Location.Id)
		log.Println(item.Location.Name)
	}
	//
	//// Explore response object
	//fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())
	//fmt.Println("  Status     :", resp.Status())
	//fmt.Println("  Proto      :", resp.Proto())
	//fmt.Println("  Time       :", resp.Time())
	//fmt.Println("  Received At:", resp.ReceivedAt())
	//fmt.Println("  Body       :\n", resp)
	//fmt.Println()
	//
	//// Explore trace info
	//fmt.Println("Request Trace Info:")
	//ti := resp.Request.TraceInfo()
	//fmt.Println("  DNSLookup     :", ti.DNSLookup)
	//fmt.Println("  ConnTime      :", ti.ConnTime)
	//fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	//fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	//fmt.Println("  ServerTime    :", ti.ServerTime)
	//fmt.Println("  ResponseTime  :", ti.ResponseTime)
	//fmt.Println("  TotalTime     :", ti.TotalTime)
	//fmt.Println("  IsConnReused  :", ti.IsConnReused)
	//fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	//fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	//fmt.Println("  RequestAttempt:", ti.RequestAttempt)

}
