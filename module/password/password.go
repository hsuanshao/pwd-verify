package password

/**
 * Date: 2020/10/29
 * Author: William Chang
 * Email: hsuanshao@gmail.com
 * Description:
 *    Password service provide various application as follows service:
 *        Validator: validator to verify password format meets the requirement to password string
 */

// Service describe interfaces to password service
type Service interface {
	// Validator to verify passsword format is correct or not by given rules, if error is not nil, previous variables will show which rule is incorrect
	Validator(pwd string) (lengthValid, uppercaseValid, lowercaseValid, numberValid, symbolValid, sequenceValid bool, err error)
}
