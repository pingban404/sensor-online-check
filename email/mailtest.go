package email

import (
    "gopkg.in/gomail.v2"
    "log"
    "sensor-online-check/config"
)

// SendAlertEmail 发送告警邮件
func SendAlertEmail(conf *config.MailConfig, subject, body string, to []string) error {
    // 创建一个新的邮件消息
    m := gomail.NewMessage()

    // 设置发件人
    m.SetHeader("From", conf.Address)

    // 设置收件人
    m.SetHeader("To", to...)

    // 设置邮件主题
    m.SetHeader("Subject", subject)

    // 设置邮件正文
    m.SetBody("text/plain", body)

    // 设置SMTP服务器配置
    d := gomail.NewDialer(conf.SMTPHost, conf.SMTPPort, conf.Address, conf.Password)

    // 启用SSL
    d.SSL = true

    // 发送邮件
    if err := d.DialAndSend(m); err != nil {
        log.Println("Error sending email:", err)
        return err
    }

    log.Println("Email sent successfully!")
    return nil
}