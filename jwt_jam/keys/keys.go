package keys

import (
	"fmt"
	"io/ioutil"
	"os"
)

var PubKey []byte
var PrvKey []byte

func init() {
	var err error
	PrvKey, err = ioutil.ReadFile("keys/private.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	PubKey, err = ioutil.ReadFile("keys/public.pem")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
