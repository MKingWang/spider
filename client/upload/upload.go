package upload

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type spiderInfo struct {
	ID     int         `json:"id"`
	Status int         `json:"status"`
	Url    string      `json:"url"`
	Worker string      `json:"worker"`
	Result interface{} `json:"result"`
}

func pushData(encode []byte) {
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

func GetData() {

	var conf map[string]string
	fd, err := os.Open("./conf/spider.conf")
	if err != nil {
		panic(err)
	}
	defer fd.Close()

	buf := bufio.NewReader(fd)
	for {
		l, _, err := buf.ReadLine()
		line := string(l)

		if err != nil {
			if len(line) == 0 {
				break
			}
			panic(err)
		}
		if strings.IndexAny(line, "=") == -1 {
			continue
		}

		confline := strings.Split(line, "=")
		conf[confline[0]] = confline[1]
		fmt.Println(confline[0])
	}

}

func Run() {
	fmt.Println("fff")
}
