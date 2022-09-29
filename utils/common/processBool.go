package common

import (
	"ethaddr-update/utils/telegram"
	"fmt"
)

// If shouldSend, send alarm.
func ProcessBool(shouldSend bool, text ...any) {
	if shouldSend {
		fmt.Println(text...)
		telegram.GetAlarmSender().Send(2091573542, text...)
	}
}
