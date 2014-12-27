package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gernest/jesus"
)

// WorkQeue is a global channel for received sms
var WorkQueue = make(chan []jesus.Inbox, 100)

// InboundMessages its what we collect from inbox
var InboundMessages []jesus.Inbox

// Collector handles http requests
func Collector(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	err := loadSMS()
	if err != nil {
		http.Error(w, "Yummy", http.StatusInternalServerError)
		return
	}
	if len(InboundMessages) == 0 {
		err = reload()
		if err != nil {
			http.Error(w, "Yummy", http.StatusInternalServerError)
			return
		}
	}

	WorkQueue <- InboundMessages
	fmt.Println("Work request queued")
	w.WriteHeader(http.StatusCreated)
	return
}

func loadSMS() error {
	localConn, err := ConnectLocalDB()
	if err != nil {
		// Do something
	}
	InboundMessages = []jesus.Inbox{}
	err = localConn.Find(&InboundMessages).Error
	if err != nil {
		// Do something
	}
	return nil
}

func reload() error {
	fmt.Println("Reloading ...")
	time.Sleep(time.Second)
	err := loadSMS()
	return err
}

// RunMigrations perfoms simple migrations
//
// Note that, its only for testing purposes and I dont intend to maintain this
func RunMigrations() {
	remoteDBConn, _ := RemoteDB()
	p := jesus.UProfile{}
	w := jesus.Withdrawal{}
	d := jesus.Deposit{}
	remoteDBConn.AutoMigrate(&p, &w, &d)
}
