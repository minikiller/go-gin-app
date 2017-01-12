package models

var tmpUserList [] User
var tmpArticleList [] Article

// This function is used to store the main lists into the temporary one
// for testing
func SaveLists() {
	tmpUserList =  UserList
	tmpArticleList =  ArticleList
}

// This function is used to restore the main lists from the temporary one
func RestoreLists() {
	 UserList = tmpUserList
	 ArticleList = tmpArticleList
}


