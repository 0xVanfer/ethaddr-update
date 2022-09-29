package common

import (
	"strings"

	"github.com/0xVanfer/types"
)

func FindInList[T types.AddressTypes](address T, list map[string]string) bool {
	for _, s := range list {
		if strings.EqualFold(types.ToString(address), s) {
			return true
		}
	}
	return false
}

func FindAndAlarm[T types.AddressTypes](address T, list map[string]string, alarmText ...any) {
	found := FindInList(address, list)
	ProcessBool(!found, alarmText...)
}
