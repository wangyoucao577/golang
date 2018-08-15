package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

func bytesInUse(username string) int64 { return 0 /* ... */ }

// Email sender configuration
// NOTE: never put passwords in source code!
const sender = "notifications@example.com"
const password = "example"
const hostname = "smtp.example.com"

const template = `warning: you are using %d bytes of storage, %d%% of your quota.`

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return //OK
	}
	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("stmp.SendMail(%s) failed: %s", username, err)
	}
}
