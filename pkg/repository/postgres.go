package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/jmoiron/sqlx"
)

type SearchResults struct {
	ready   bool
	Query   string
	Results []Result
}

type Result struct {
	Name, Description, URL string
}

func (sr *SearchResults) UnmarshalJSON(bs []byte) error {
	array := []interface{}{}
	if err := json.Unmarshal(bs, &array); err != nil {
		return err
	}
	sr.Query = array[0].(string)
	for i := range array[1].([]interface{}) {
		sr.Results = append(sr.Results, Result{
			array[1].([]interface{})[i].(string),
			array[2].([]interface{})[i].(string),
			array[3].([]interface{})[i].(string),
		})
	}
	return nil
}

func WikipediaAPI(request string) (answer []string) {
	s := make([]string, 3)
	if response, err := http.Get(request); err != nil {
		s[0] = "Error"
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		sr := SearchResults{}
		if err := json.Unmarshal(contents, &sr); err != nil {
			s[0] = "Error"
		}
		if !sr.ready {
			s[0] = "Error"
		}
		for i := range sr.Results {
			s[i] = sr.Results[i].URL
		}
	}
	return s
}

func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

type Config struct {
	Host    string
	Port    string
	User    string
	Pass    string
	DBName  string
	SSLmode string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.DBName, cfg.Pass, cfg.SSLmode))
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Файлы миграции + бд в docker
