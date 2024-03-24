package database

import (
	"database/sql"
	"fmt"
	_ "log"
	_ "sync"
)

// This file was from this tutorial https://earthly.dev/blog/golang-sqlite/

var ErrIDNotFound = fmt.Errorf("ID not found")

type Activity struct {
	ID          int
	Time        string
	Description string
}

const create string = `
  CREATE TABLE IF NOT EXISTS activities (
  id INTEGER NOT NULL PRIMARY KEY,
  time DATETIME NOT NULL,
  description TEXT
  );`

const file string = "test.db"

func (s *service) NewActivities() (sql.Result, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	result, err := db.Exec(create)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// func (c *Activities) Insert(activity Activity) (int, error) {
// 	res, err := c.db.Exec("INSERT INTO activities VALUES(NULL,?,?);", activity.Time, activity.Description)
// 	if err != nil {
// 		return 0, err
// 	}

// 	var id int64
// 	if id, err = res.LastInsertId(); err != nil {
// 		return 0, err
// 	}
// 	return int(id), nil
// }

// func (c *Activities) Retrieve(id int) (Activity, error) {
// 	log.Printf("Getting %d", id)

// 	// Query DB row based on ID
// 	row := c.db.QueryRow("SELECT id, time, description FROM activities WHERE id=?", id)

// 	// Parse row into Activity struct
// 	activity := Activity{}
// 	var err error
// 	if err = row.Scan(&activity.ID, &activity.Time, &activity.Description); err == sql.ErrNoRows {
// 		log.Printf("Id not found")
// 		return Activity{}, ErrIDNotFound
// 	}
// 	return activity, err
// }

// func (c *Activities) List(offset int) ([]Activity, error) {
// 	rows, err := c.db.Query("SELECT * FROM activities WHERE ID > ? ORDER BY id DESC LIMIT 100", offset)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	data := []Activity{}
// 	for rows.Next() {
// 		i := Activity{}
// 		err = rows.Scan(&i.ID, &i.Time, &i.Description)
// 		if err != nil {
// 			return nil, err
// 		}
// 		data = append(data, i)
// 	}
// 	return data, nil
// }
