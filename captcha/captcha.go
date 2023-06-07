package captcha

import (
	"github.com/mojocn/base64Captcha"
)

var driver = base64Captcha.DefaultDriverDigit

func Generate() (id, b64s string, answer string, err error) {
	id, content, answer := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	if err != nil {
		return "", "", "", err
	}
	b64s = item.EncodeB64string()
	return
}

//type Driver interface {
//	//DrawCaptcha draws binary item
//	DrawCaptcha(content string) (item Item, err error)
//	//GenerateIdQuestionAnswer creates rand id, content and answer
//	GenerateIdQuestionAnswer() (id, q, a string)
//}
