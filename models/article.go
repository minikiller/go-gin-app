// models.article.go

package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func GetAllArticles() []Article {
	return ArticleList
}

// Fetch an article based on the ID supplied
func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}

// Create a new article with the title and content provided
func CreateNewArticle(title, content string) (*Article, error) {
	// Set the ID of a new article to one more than the number of articles
	a := Article{ID: len(ArticleList) + 1, Title: title, Content: content}

	// Add the article to the list of articles
	ArticleList = append(ArticleList, a)

	return &a, nil
}
