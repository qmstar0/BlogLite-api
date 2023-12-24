package users

import (
	"blog/infra/e"
	"blog/infra/event"
	"blog/infra/shutdown"
	"blog/utils/logging"
	mail2 "blog/utils/mail"
	"context"
	"fmt"
)

const EmailSubject = "Please Confirm Your email"

var (
	bus    = event.GetBus(event.DomainServiceBus)
	logger = logging.New("users-commonEvent")
)

func init() {
	var sendEmailEventChan = make(event.DataChan)
	bus.Subscribe(event.SendMail, sendEmailEventChan)
	ctx, cc := context.WithCancel(context.Background())
	shutdown.WaitForShutdown().Add(func() {
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
	html := mail2.NewTemplateForVerifyCode(ed.Email, ed.Captcha)
	m := mail2.PostMan.NewMail()
	m.SetHeader("Subject", EmailSubject)
	m.SetHeader("To", ed.Email)
	m.SetBody("text/html", html.ToString())
	dialer := mail2.PostMan.NewDialer()
	if err := dialer.DialAndSend(m); err != nil {
		return e.NewError(e.EmailSendErr, err)
	}
	return nil
}
