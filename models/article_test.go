// models.article_test.go

package models

import "testing"

// Test the function that fetches all articles
func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	// Check that the length of the list of articles returned is the
	// same as the length of the global variable holding the list
	if len(alist) != len(ArticleList) {
		t.Fail()
	}

	// Check that each member is identical
	for i, v := range alist {
		if v.Content != ArticleList[i].Content ||
			v.ID != ArticleList[i].ID ||
			v.Title != ArticleList[i].Title {

			t.Fail()
			break
		}
	}
}

// Test the function that fetche an Article by its ID
func TestGetArticleByID(t *testing.T) {
	a, err := GetArticleByID(1)

	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
		t.Fail()
	}
}

// Test the function that creates a new article
func TestCreateNewArticle(t *testing.T) {
	// get the original count of articles
	originalLength := len(GetAllArticles())

	// add another article
	a, err := CreateNewArticle("New test title", "New test content")

	// get the new count of articles
	allArticles := GetAllArticles()
	newLength := len(allArticles)

	if err != nil || newLength != originalLength+1 ||
		a.Title != "New test title" || a.Content != "New test content" {

		t.Fail()
	}
}
