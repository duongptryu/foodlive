package common

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

var (
	RecordNotFound = errors.New("record not found")
)

func AppRecovery() {
	if err := recover(); err != nil {
		log.Error("Recovery error -", err)
	}
}
