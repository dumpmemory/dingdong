package app

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	"dingdong/internal/app/api"
	"dingdong/internal/app/config"
	"dingdong/internal/app/service"
)

func Run() {
	go Monitor()
	go service.SnapUp()
	go service.PickUp()
	go service.Notify()

	http.HandleFunc("/", api.SayWelcome)
	http.HandleFunc("/set", api.SetConfig)
	http.HandleFunc("/address", api.GetAddress)

	conf := config.Get()
	err := http.ListenAndServe(conf.Addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe => ", err)
	}
}

func isPeak() bool {
	now := time.Now()
	if now.Hour() >= 0 && now.Hour() < 9 {
		return true
	}
	return false
}

// Monitor 监视器 监听运力
func Monitor() {
	cartMap := service.MockCartMap()
	executedCount := 0
	for {
		conf := config.Get()
		duration := conf.MonitorIntervalMin + rand.Intn(conf.MonitorIntervalMax-conf.MonitorIntervalMin)
		if !conf.MonitorNeeded {
			<-time.After(time.Duration(duration) * time.Second)
			continue
		}
		if isPeak() {
			log.Println("当前高峰期或暂未营业")
			<-time.After(time.Duration(duration) * time.Second)
			continue
		}
		executedCount++
		cartMap = service.MockCartMap()
		service.GetMultiReserveTimeAndNotify(cartMap)
		<-time.After(time.Duration(duration) * time.Second)
	}
}
