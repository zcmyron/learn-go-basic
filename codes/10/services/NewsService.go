package services

type NewsService struct {
	s string
}

func (ns *NewsService) Get(id int) string {
	return "news1"
}
