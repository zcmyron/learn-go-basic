package services

type NewsService struct {
}

func (ns *NewsService) Get(id int) string {
	return "news1"
}
