package news

import (
	"fmt"
	"log"
)

type newsList []News

// Manager ...
type Manager interface {
	All() newsList
	Get(id int) News
}

// newsManager ...
type newsManager struct{}

// getOneNews ...
func getOneNews(title string, text string) News {
	return News{
		Title: title,
		Text:  text,
	}
}

// Get return one news from db by id
func (m newsManager) Get(id int) News {
	return getOneNews("Some news", "Some text")
}

// All - return all news from database
func (m newsManager) All() newsList {
	log.Printf("Gettin all news in newsManager")
	fewNews := make(newsList, 0, 5)
	for i := 0; i < 5; i++ {
		fewNews = append(fewNews, getOneNews(fmt.Sprintf("Some news %d", i+1), "Some text"))
	}
	return fewNews
}
