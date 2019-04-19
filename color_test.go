package nbcolor

import (
	"log"
	"os"
	"testing"
)

func TestColor(t *testing.T) {
	f, err := os.OpenFile("log.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Fatal(err)
	}
	logger := log.New(f, "nbcolor", log.Ltime|log.Lshortfile)
	text := "hello world"
	logger.Println(Black(text))
	logger.Println(Red(text))
	logger.Println(Green(text))
	logger.Println(Brown(text))
	logger.Println(Blue(text))
	logger.Println(Purple(text))
	logger.Println(Cyan(text))
	logger.Println(LightGray(text))
	logger.Println(DarkGray(text))
	logger.Println(LightRed(text))
	logger.Println(LightGreen(text))
	logger.Println(Yellow(text))
	logger.Println(LightBlue(text))
	logger.Println(LightPurple(text))
	logger.Println(LightCyan(text))
	logger.Println(White(text))
}
