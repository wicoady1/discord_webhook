package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var connString string

func Init(conn string) {
	connString = conn
}

func GetPendingArticleList() ([]ArticleData, error) {
	result := []ArticleData{}

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("[Open DB Error]", err)
	}
	defer db.Close()

	query := `
		SELECT a.post_author, b.user_login, a.post_status, a.post_date_gmt, a.post_title, a.post_modified_gmt
		FROM wp_posts a
		INNER JOIN wp_users b ON a.post_author = b.ID
		WHERE a.post_status IN ('pending','draft')
		AND a.post_type =  'post'
	`

	rows, err := db.Query(query)
	if err != nil {
		log.Println("[GetPendingArticleList]", err)
	}
	defer rows.Close()

	for rows.Next() {
		var articleData ArticleData
		err := rows.Scan(&articleData.AuthorID, &articleData.AuthorName, &articleData.Status, &articleData.PostTime, &articleData.Title, &articleData.UpdatedTime)
		if err != nil {
			log.Println("[GetPendingArticleList]", err)
			return []ArticleData{}, err
		}
		result = append(result, articleData)
	}
	return result, nil
}

func GetUser(userID int64) (string, string) {

	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("[GetUser]", err)
		return "", ""
	}
	defer db.Close()

	query := fmt.Sprintf(`
		SELECT user_login, display_name
		FROM wp_users
		WHERE ID = %d
		LIMIT 1
	`, userID)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("[GetUser]", err)
		return "", ""
	}
	defer rows.Close()

	var userLogin, userDisplay string
	for rows.Next() {
		err := rows.Scan(&userLogin, &userDisplay)
		if err != nil {
			log.Println("[GetUser]", err)
		}
	}

	return userLogin, userDisplay
}
