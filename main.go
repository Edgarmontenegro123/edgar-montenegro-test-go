package main

import (
	"fmt"
	"github.com/Edgarmontenegro123/think-us-challenge/encrypt"
)

func main() {
	key := "dcj"
	message := "I love prOgrAmming!"
	encryptedMessage := encrypt.EncryptMessage(key, message)

	fmt.Println(encryptedMessage)

	/*key := ""
	message := ""
	encryptedMessage := encrypt.EncryptMessage(key, message)

	fmt.Println(encryptedMessage)*/
}
