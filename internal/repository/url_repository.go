package repository

import "url-shortener/pkg/db"

func Save(original string) (int, error) {
	var id int
	err := db.DB.QueryRow(
		"INSERT INTO urls(original_url) VALUES($1) RETURNING id",
		original,
	).Scan(&id)

	return id, err
}

func UpdateCode(id int, code string) error {
	_, err := db.DB.Exec(
		"UPDATE urls SET short_code=$1 WHERE id=$2",
		code, id,
	)
	return err
}

func Get(code string) (string, error) {
	var url string
	err := db.DB.QueryRow(
		"SELECT original_url FROM urls WHERE short_code=$1",
		code,
	).Scan(&url)

	return url, err
}

func SaveURL(shortCode string, originalURL string) error {
	query := "INSERT INTO urls (short_code, original_url) VALUES ($1, $2)"
	_, err := db.DB.Exec(query, shortCode, originalURL)
	return err
}
