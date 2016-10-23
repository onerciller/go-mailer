# go-mailer
:mailbox_with_mail: Mailer for Go 

```go
package main

import (
	"github.com/onerciller/go-mailer"
)

var config = mail.Config{
	Host:     "smtp.yandex.com.tr",
	Username: "onerciller@yandex.com.tr",
	Password: "PASSWORD",
	Port:     "587",
	From:     "onerciller@yandex.com.tr",
}

func main() {
	mail := mail.New(config)
	mail.SetTo("onerciller@gmail.com", "onerciller@yandex.com.tr")
	mail.SetSubject("this is subject")
	mail.SetBody("this is body")
	mail.Send()
}

```
