package models

import "github.com/pquerna/ffjson/ffjson"

type NewsModel struct {
	NewsID                 int
	NewsTitle, NewsContent string
}

func (news NewsModel) ToJSON() string {
	res, err := ffjson.Marshal(news)
	if err != nil {
		return err.Error()
	} else {
		return string(res)
	}
}

func (news *NewsModel) Get(id int) string {
	return "News1"
}
