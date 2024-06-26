package models

import (
	"database/sql"
	"errors"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type NoteModel struct {
	DB *sql.DB
}

func (m *NoteModel) Insert(title, content string, expires int) (int, error) {
	stmt := `INSERT INTO notes (title, content, created, expires)
	VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *NoteModel) Get(id int) (*Note, error) {
	stmt := `SELECT id, title, content, created, expires FROM notes
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	n := &Note{}

	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return n, nil
}

func (m *NoteModel) Latest() ([]*Note, error) {
	stmt := `SELECT id, title, content, created, expires FROM notes
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	notes := []*Note{}

	for rows.Next() {
		n := &Note{}

		err = rows.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
		if err != nil {
			return nil, err
		}

		notes = append(notes, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
