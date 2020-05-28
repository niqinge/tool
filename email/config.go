package email

import "github.com/pkg/errors"

var (
	ErrToIsNil = errors.New("To is nil")
)

type Config struct {
	Host     string `json:"host" name:"host"`
	Post     int    `json:"post" name:"端口"`
	From     string `json:"from" name:"发件人"`
	Password string `json:"password" name:"密码"`
}

type Option struct {
	To      []string `json:"to" name:"收件人邮箱组"`
	Cc      []string `json:"cc" name:"抄送人邮箱组"`
	Subject string   `json:"subject" name:"主题"`
	Content string   `json:"content" name:"内容"`
}

func (o *Option) IsValid() error {
	switch {
	case len(o.To) == 0:
		return ErrToIsNil
	default:
		return nil
	}
}
