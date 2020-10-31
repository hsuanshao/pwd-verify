package deps

import (
	"flag"

	"github.com/hsuanshao/pwd-verify/app/infra/di"
	"github.com/hsuanshao/pwd-verify/module/password"
	"github.com/sirupsen/logrus"
)

func init() {
	di.Register(&passwordValidatorSer{})
}

var (
	minLength            = flag.Uint("min", 5, "")
	maxLength            = flag.Uint("max", 12, "")
	requireUppercase     = flag.Bool("uppercase", false, "")
	requireLowercase     = flag.Bool("lowercase", true, "")
	requireNumber        = flag.Bool("number", true, "")
	allowedSequence      = flag.Bool("sequence", false, "")
	requireSpecialSymbol = flag.Bool("symbol", false, "")
)

type passwordValidatorSer struct {
	Password *password.Service
}

func (im *passwordValidatorSer) Regist() error {
	pwdSer := password.New(uint8(*minLength), uint8(*maxLength), *requireUppercase, *requireLowercase, *requireNumber, *allowedSequence, *requireSpecialSymbol)
	im.Password = &pwdSer
	return nil
}

func (im *passwordValidatorSer) Destroy() error {
	im.Password = nil
	return nil
}

func (im *passwordValidatorSer) Get() interface{} {
	return im.Password
}

// GetPasswordValidator returns Password.Service
func GetPasswordValidator() password.Service {
	pvs, err := di.Get(&passwordValidatorSer{})
	if err != nil {
		logrus.WithField("err", err.Error()).Panic("get password validator service failure")
	}

	pvsInterface, ok := pvs.(*password.Service)
	if !ok {
		logrus.Panic("transfer password validtor service interface failure")
	}

	return *pvsInterface
}
