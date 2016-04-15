package godatayes

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	"github.com/lhboy1984/godatayes/helper"
)

var Token string = "93927cca1ee5fece9187f990e0187f208e9a4dc954253cd6f6d8751a406cedc2"
var client *http.Client

func init() {
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: false},
			DisableCompression: true,
		},
	}
}

func GetData(arg *Argument) string {
	req, err := http.NewRequest("GET", "https://api.wmcloud.com:443/data/v1"+arg.URL(), nil)
	helper.IfErrPanic(err)
	req.Header.Add("Authorization", "Bearer "+Token)
	rsp, err := client.Do(req)
	helper.IfErrPanic(err)
	helper.IfFailPanic(rsp.StatusCode == 200, "service invalid")

	data, _ := ioutil.ReadAll(rsp.Body)
	return string(data)
}
