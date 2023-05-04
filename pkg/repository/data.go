package repository

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

type DataPostgres struct {
	db *sqlx.DB
}

func NewDataPostgres(db *sqlx.DB) *DataPostgres {
	return &DataPostgres{db: db}
}

func (r *DataPostgres) CollectData(username string, chatid int64, message string, answer []string) error {

	//Конвертируем срез с ответом в строку
	answ := strings.Join(answer, ", ")

	//Создаем SQL запрос
	query := `INSERT INTO users(username, chat_id, message, answer) VALUES($1, $2, $3, $4)`

	//Выполняем наш SQL запрос
	if _, err := r.db.Exec(query, `@`+username, chatid, message, answ); err != nil {
		return err
	}

	return nil
}
func (r *DataPostgres) GetNumberOfUsers() (int64, error) {
	var count int64
	//Отправляем запрос в БД для подсчета числа уникальных пользователей
	row := r.db.QueryRow("SELECT COUNT(DISTINCT username) FROM users")
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
