package provider

import (
	"database/sql"
	"fmt"
	"log"
	"todolist/internal/entities"
)

type Provider struct {
	conn *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{conn: conn}
}

func (p *Provider) InsertTask(task entities.Task) (*entities.Task, error) {
	var id int
	err := p.conn.QueryRow(`INSERT INTO tasks (author_name, assignee_name, status) VALUES ($1, $2, $3) RETURNING id`,
		task.AuthorName, task.AssigneeName, task.Status).Scan(&id)
	if err != nil {
		return nil, err
	}
	task.ID = id
	return &task, nil
}

func (p *Provider) SelectAllTasks() ([]*entities.Task, error) {
	var tasks []*entities.Task
	rows, err := p.conn.Query(`SELECT id, author_name, assignee_name, created_date, resolved_date, status FROM tasks`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task entities.Task
		var resolvedDate sql.NullTime
		if err := rows.Scan(&task.ID, &task.AuthorName, &task.AssigneeName, &task.CreatedDate, &resolvedDate, &task.Status); err != nil {
			return nil, err
		}
		task.ResolvedDate = resolvedDate
		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (p *Provider) UpdateTaskStatus(id int, status string) (*entities.Task, error) {
	var task entities.Task
	err := p.conn.QueryRow(`UPDATE tasks SET status = $1, resolved_date = CURRENT_TIMESTAMP WHERE id = $2 RETURNING id, author_name, assignee_name, created_date, resolved_date, status`,
		status, id).Scan(&task.ID, &task.AuthorName, &task.AssigneeName, &task.CreatedDate, &task.ResolvedDate, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (p *Provider) SelectTaskByID(id int) (*entities.Task, error) {
	var task entities.Task
	err := p.conn.QueryRow(`SELECT id, author_name, assignee_name, created_date, resolved_date, status FROM tasks WHERE id = $1`, id).Scan(&task.ID, &task.AuthorName, &task.AssigneeName, &task.CreatedDate, &task.ResolvedDate, &task.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, entities.ErrTaskNotFound
		}
		return nil, err
	}
	return &task, nil
}
