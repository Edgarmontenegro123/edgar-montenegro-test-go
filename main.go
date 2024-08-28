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

	arr := []int{3, 4, -7, 5, -6, 2, 5, -1, 8}
	// arr := []int{1, 2, -3, 3, 1, -1}
	// arr := []int{2, 3, -5, -2, 7}
	result := encrypt.RemoveZeroSumSubArrays(arr)
	fmt.Println(result)
}
