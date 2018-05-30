package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	spiderinfo spiderInfo
	client     *http.Client
}

type spiderInfo struct {
	ID     int         `json:"id"`
	Status int         `json:"status"`
	Url    string      `json:"url"`
	Worker string      `json:"worker"`
	Result interface{} `json:"result"`
}

var spiderData Data

func (this *Data) Push() {
	spinfo := spiderInfo{10, 5, "http://blog.leixin.wang", "192.168.1.5", "abcdddaaabbbccc"}

	var client http.Client
	encode, err := json.Marshal(spinfo)
	if err != nil {
		fmt.Println("abc", err.Error)
	}

	request, err := http.NewRequest("POST", "http://127.0.0.1:8080/spiderapi", bytes.NewBuffer(encode))
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:52.0) Gecko/20100101 Firefox/52.0")

	res, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer

	io.Copy(&buf, res.Body)

	fmt.Println(buf.String())
}
