package pushplus

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type PushPlus struct {
	Token string
	Title string
}

func New(token, title string) *PushPlus {
	if title == "" {
		title = "PushPlus"
	}
	return &PushPlus{
		Token: token,
		Title: title,
	}
}

func (p PushPlus) Send(content string) error {
	urlstr := "http://www.pushplus.plus/send/" + p.Token + "?content=" + url.QueryEscape(content) + "&title=" + url.QueryEscape(p.Title)
	resp, err := http.Get(urlstr)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s", resp.Status)
	}
	r, _ := io.ReadAll(resp.Body)
	log.Println(string(r))
	return nil
}

func (PushPlus) Name() string {
	return "pushplus"
}

func (p *PushPlus) EditTitle(title string) {
	p.Title = title
}

func (p PushPlus) GetTitle() string {
	return p.Title
}
