package mirai

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Mirai struct {
	host   string
	target int64
}

func New(host string, target int64) *Mirai {
	return &Mirai{
		host:   host,
		target: target,
	}
}

func (m Mirai) Send(content string) error {
	host := m.host
	url := host + "/sendFriendMessage"
	method := "POST"
	data := fmt.Sprintf(`{"target":%d,"messageChain":[{ "type":"Plain", "text":"%s" }]}`, m.target, content)
	// fmt.Println(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(data))

	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))
	return nil
}

func (Mirai) Name() string {
	return "mirai"
}

func (Mirai) EditTitle(title string) {

}

func (Mirai) GetTitle() string {
	return "mirai"
}
