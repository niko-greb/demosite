package pkg

import (
	"math/rand"
	"strings"
)

func RandomUrl() string {
	symbSet := "abcdefghijklmnopqrstyvwxyzABCDEFGHIJKLMNOPQRSTYVWXYZ1234567890_"
	var shorty strings.Builder
	for i := 0; i < 10; i++ {
		random := rand.Intn(len(symbSet))
		randomSymb := symbSet[random]
		shorty.WriteString(string(randomSymb))
	}
	shrt := shorty.String()
	var finalShrt string
	if strings.Contains(shrt, "_"){
		finalShrt = shrt
	} else {
		random := rand.Intn(len(shrt))
		finalShrt = strings.Replace(shrt, string(shrt[random]), "_", 1 )

	}
	return finalShrt
}

