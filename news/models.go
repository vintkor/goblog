package news

type model interface {
	Objects() Manager
}

// News ...
type News struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

// Objects return manager for model
func (n *News) Objects() Manager {
	return newsManager{}
}
