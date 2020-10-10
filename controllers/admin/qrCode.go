package admin

import (
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
)

type QRCodeController struct {
	BaseController
}


func (q *QRCodeController) QR() {
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 256, "static/upload/qr.png")
	if err != nil{
		q.Ctx.WriteString("QR error")
		fmt.Println(err)
	}
	q.Ctx.WriteString("QR success")
}