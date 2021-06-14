package utils

import (
	"fmt"
	"testing"
	"time"

	"github.com/betelgeuse-7/words/utils"
)

func TestMakeNew(t *testing.T) {
	ACCESS_SECRET := []byte("Secret123")
	//REFRESH_SECRET := []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

	atStruct := utils.AccessToken{
		SecretKey: ACCESS_SECRET,
		Payload: struct{ UserId int }{
			UserId: 1,
		},
		Expiry: time.Now().Add(time.Minute * 2).Unix(),
	}
	got, err := atStruct.MakeNew()
	if err != nil {
		fmt.Println(err)
		t.Errorf("got err %v", err)
	}

	fmt.Println(got)
}
