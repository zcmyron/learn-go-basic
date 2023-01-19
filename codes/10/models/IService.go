package models

type Iservice interface {
	Get(id int) string
}
