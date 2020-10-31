# Password Validator

It is an example to verify password format by given configuration  

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

## Run test

It is simple to run test, all you need to do is in root path of this project and type `make test`  
