package db

import (
	"database/sql"
	"fmt"
	"log"
	"slices"
	"testing"

	"github.com/ahmetalpbalkan/go-linq"
	"github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func connect() *sql.DB {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "ula123$",
		Net:    "tcp",
		Addr:   "localhost:13306",
		DBName: "test",
	}

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func listAlbums(title string, name string) ([]Album, error) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM album a WHERE a.title = ? and a.artist = ?", title, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}

func listAlbumsByArtist(artist string) ([]Album, error) {
	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM album a WHERE a.artist = ?", artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, err
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return albums, nil
}

func TestListAlbums01(t *testing.T) {
	albums, err := listAlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}
	var orderedAlbums []Album
	linq.From(albums).OrderBy(func(i interface{}) interface{} {
		return i.(Album).Price
	}).ToSlice(&orderedAlbums)
	fmt.Printf("Albums found: %v\n", orderedAlbums)
}

func TestListAlbums02(t *testing.T) {
	albums, err := listAlbumsByArtist("John Coltrane")
	if err != nil {
		log.Fatal(err)
	}

	slices.SortFunc(albums, func(l, r Album) int {
		if l.Price < r.Price {
			return -1
		}
		if l.Price > r.Price {
			return 1
		}
		return 0
	})

	result := lo.Filter(albums, func(a Album, index int) bool {
		return a.Price > 50
	})

	fmt.Printf("Albums found: %v\n", result)
}

func TestTextTemplate01(t *testing.T) {
	// tmpl, err := template.New("test").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = tmpl.ExecuteTemplate(os.Stdout, "T", "<script>alert('you have been pwned')</script>")

	// err = tmpl.Execute(log.Writer(), struct{ Name string }{"Alice"})
	// if err != nil {
	// 	log.Fatal(err)
	// }

	s := fmt.Sprintf("Hello, %v!", "Alice")

	fmt.Println(s)
}
