package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	var filename, url string
	fmt.Println("QR Code oluşturmak istediğiniz bağlantıyı girin:")
	fmt.Scan(&url)
	fmt.Println("QR Code kaydetmek istediğiniz ismi girin:")
	fmt.Scan(&filename)
	filename = filename + ".png"
	qrcode.WriteFile(url, qrcode.Medium, 256, filename)

}
