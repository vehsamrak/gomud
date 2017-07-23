package commands

import (
	"github.com/Vehsamrak/gomud/console"
	"net"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
)

const DEFAULT_ENCODING = "utf-8"

type Sender struct {
	ConnectionPointer *net.Conn
	charset string
}

func (sender *Sender) toClient(message string)  {
	console.Client(sender.ConnectionPointer, sender.translateToClientCharset(message))
}

func (sender *Sender) toServer(message string)  {
	console.Server(sender.translateToUtf8(message))
}

func (sender *Sender) translateToClientCharset(message string) string {
	charsetTranslator, _ := charset.TranslatorTo(sender.charset)
	_, translatedMessageBytes, _ := charsetTranslator.Translate([]byte(message), false)

	return string(translatedMessageBytes)
}

func (sender *Sender) translateToUtf8(message string) string {
	charsetTranslator, _ := charset.TranslatorFrom(sender.charset)
	_, translatedMessageBytes, _ := charsetTranslator.Translate([]byte(message), false)

	return string(translatedMessageBytes)
}