package commands

import (
	"github.com/Vehsamrak/gomud/console"
	"net"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"github.com/Vehsamrak/gomud/player"
)

const DEFAULT_ENCODING = "utf-8"
const ENCODING_WINDOWS_1251 = "windows-1251"
const ENCODING_KOI8 = "koi8-r"
const ENCODING_WINDOWS_1252 = "windows-1252"

type Sender struct {
	ConnectionPointer *net.Conn
	charset string
}

func (sender *Sender) toClient(message string)  {
	console.Client(sender.ConnectionPointer, sender.encodeToClientCharset(message))
}

func (sender *Sender) toServer(message string)  {
	console.Server(message)
}

func (sender *Sender) toAllClients(players map[string]*player.Player, message string)  {
	for _, player := range players {
		message := sender.encodeToCharset(player.Codepage, "Chat: " + message)
		console.Client(player.ConnectionPointer, message)
	}
}

func (sender *Sender) encodeToClientCharset(message string) string {
	return sender.encodeToCharset(sender.charset, message)
}

func (sender *Sender) encodeToCharset(encoding string, message string) string {
	processedMessageBytes := sender.fixYaLetter(encoding, message)
	charsetTranslator, _ := charset.TranslatorTo(encoding)

	_, translatedMessageBytes, _ := charsetTranslator.Translate(processedMessageBytes, false)

	return string(translatedMessageBytes)
}

func (sender *Sender) encodeToUtf8(message string) string {
	charsetTranslator, _ := charset.TranslatorFrom(sender.charset)
	_, translatedMessageBytes, _ := charsetTranslator.Translate([]byte(message), false)

	return string(translatedMessageBytes)
}

func (sender *Sender) fixYaLetter(encoding string, message string) []byte {
	messageBytes := []byte(message)

	if encoding != ENCODING_WINDOWS_1251 {
		return messageBytes
	}

	var processedMessageBytes []byte
	prebyte := false

	// double "я"-letter (209 & 143 bytes) to fix CP1251 issue
	for _, messageByte := range messageBytes {
		if prebyte && messageByte == 143 {
			processedMessageBytes = append(processedMessageBytes, 143)
			processedMessageBytes = append(processedMessageBytes, 209)
		}

		if messageByte == 209 {
			prebyte = true
		} else {
			prebyte = false
		}

		processedMessageBytes = append(processedMessageBytes, messageByte)
	}

	return processedMessageBytes
}

// remove doubled "я"-letter (209 & 143 bytes) to fix CP1251 issue
func (sender *Sender) removeYaLetterDuplication(message string) string {
	if sender.charset != ENCODING_WINDOWS_1251 {
		return message
	}

	messageBytes := []byte(message)

	firstYaDeleted := false
	deleted := 0
	for i := range messageBytes {
		j := i - deleted

		if messageBytes[j] == 143 {
			if messageBytes[j-1] == 209 {
				if !firstYaDeleted {
					messageBytes = messageBytes[:j+copy(messageBytes[j-1:], messageBytes[j+1:])]
					deleted++
				}

				firstYaDeleted = !firstYaDeleted

				continue
			}
		}
	}

	return string(messageBytes)
}
