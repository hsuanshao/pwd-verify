# Password Validator

It is an example to verify password format by given configuration  

## About Validator return parameters

Since you `New` a Password.Service, you need set follows configuration to tell what kind of rule set you want:  
(uint8)  
minLength: the short password request
maxLength: the maximum password length be accepted
(bool)  
requireUppercase: Do you require at least one uppercase character in password?  
requireLowercase: Do you require at least one lowercase character in password?  
requireNumber: Do you require at least one number character in password?  
allowSequence: Do you allowed sequence characters in password?  
requireSymbol: Do you require at least one special symbol in password?  

And while password be varified by `Password.Validtor` it will return multiple values length, uppercase, lowercase, number, symbol, sequence bool, err error).  

If err return is not nil, which mean password does not meets the configuration.  
Others parameters can helps to tell you which rule been violation.  

## Want to apply password validator module in your project

If you want applied this password validator in your project. You can use `go mod` to mangem and get `password validator` module into your project

`go get github.com/hsuanshao/pwd-verify/module/password`  

The you can import it in your project.  

    import (
        "github.com/hsuanshao/pwd-verify/module/password"
    )

## Dependency Injection

There are many dependency injection solition in Go. But I choose to do another example to implement dependency inejction solution by myself in this programming task.  Which is placed in /app/infra/di.  
  
Other famous dependency injection solution:  
  
[Wire](https://github.com/google/wire)  
[Uber-Dig](https://github.com/uber-go/dig)  
[sarulabs-di](https://github.com/sarulabs/di) this almost similar to uber dig, but without use `reflect`, therefor, all deps must have unique name.  

note: reflect affect performance.  

You can get the example to get `Password.Service` by dependency injection in `app/main.go`  

## Run test

It is simple to run test, all you need to do is in root path of this project and type `make test`  
