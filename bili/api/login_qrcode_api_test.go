package api

import (
	"fmt"
	"github.com/alice52/archive/common/global"
	"github.com/skip2/go-qrcode"
	"github.com/wordpress-plus/kit-logger/viperx"
	"os"
	"testing"
)

func TestQrCodeImage(t *testing.T) {
	defer func(fn string) {
		os.Remove(fn)
	}("qr.png")

	err := qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "qr.png")
	if err != nil {
		return
	}
}

func TestGenerateAndEmail(t *testing.T) {
	global.VIPER = viperx.Viper(&global.CONFIG, "../../config.yaml")

	err := GenerateAndEmail("zack", qrcode.Low, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
}
