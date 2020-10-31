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
	_, _, _, _, _, _, err := pv.Validator("a3dh9sa3s")

	if err != nil {
		logrus.WithField("err", err.Error())
	}

}
