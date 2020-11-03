package main

import (
	"flag"

	"github.com/hsuanshao/pwd-verify/app/infra/di/deps"
	"github.com/sirupsen/logrus"
)

/**
 * an example to use dependency injection to get password validator service
 */

func main() {
	flag.Parse()

	pv := deps.GetPasswordValidator()
	lengthV, upperV, lowerV, numberV, symbolV, sequenceV, err := pv.Validator("a12A9sa3s")

	if err != nil {
		logrus.WithFields(logrus.Fields{"length": lengthV, "up": upperV, "lower": lowerV, "number": numberV, "symbol": symbolV, "sequence": sequenceV, "err": err.Error()}).Error("Password Validation failure")
	}

}
