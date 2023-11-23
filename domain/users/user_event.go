package users

import (
	"blog/infra/e"
	"blog/infra/event"
	"blog/infra/logging"
	"blog/infra/mail"
	"blog/utils"
	"context"
	"fmt"
)

const EmailSubject = "Please Confirm Your email"

var (
	bus    = event.GetBus(event.DomainServiceBus)
	logger = logging.New("user-event")
)

func init() {
	var sendEmailEventChan = make(event.DataChan)
	bus.Subscribe(event.SendMail, sendEmailEventChan)
	ctx, cc := context.WithCancel(context.Background())
	utils.WaitForShutdown().Add(func() {
		cc()
	})
	go func(c context.Context) {
		for {
			select {
			case ed := <-sendEmailEventChan:
				go func() {
					if err := SendCaptcpaEmail(ed.Data); err != nil {
						logger.Warnf("%s:%s:%d:%v", ed.Bus, ed.Topic, ed.Time, ed.Data)
					}
				}()
			case <-c.Done():
				fmt.Println("事件循环监听已关闭")
				return
			}

		}
	}(ctx)
}

func SendCaptcpaEmail(data any) error {
	ed, ok := data.(event.SendEmailED)
	if !ok {
		return e.NewError(e.DomainEventDataTypeErr, nil)
	}
	html := mail.NewTemplateForVerifyCode(ed.Email, ed.Captcha)
	m := mail.PostMan.NewMail()
	m.SetHeader("Subject", EmailSubject)
	m.SetHeader("To", ed.Email)
	m.SetBody("text/html", html.ToString())
	dialer := mail.PostMan.NewDialer()
	if err := dialer.DialAndSend(m); err != nil {
		return e.NewError(e.EmailSendErr, err)
	}
	return nil
}
