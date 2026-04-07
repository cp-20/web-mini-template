package handler

import (
	"context"
	"fmt"

	"github.com/cp-20/web-mini-template/server/internal/gen/openapi"
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	db *sqlx.DB
}

var _ openapi.Handler = (*Handler)(nil)

func New(db *sqlx.DB) *Handler {
	return &Handler{db: db}
}

type feedRow struct {
	ProjectID   int64  `db:"project_id" json:"project_id"`
	ProjectName string `db:"project_name" json:"project_name"`
	OwnerName   string `db:"owner_name" json:"owner_name"`
	TaskID      int64  `db:"task_id" json:"task_id"`
	TaskTitle   string `db:"task_title" json:"task_title"`
	TaskStatus  string `db:"task_status" json:"task_status"`
}

type memberRow struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type createTaskRequest struct {
	Title    string `json:"title"`
	MemberID int64  `json:"member_id"`
}

func (h *Handler) Initialize(ctx context.Context) error {
	tx, err := h.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer tx.Rollback()

	queries := []string{
		`SET FOREIGN_KEY_CHECKS = 0;`,
		`DROP TABLE IF EXISTS tasks;`,
		`DROP TABLE IF EXISTS projects;`,
		`DROP TABLE IF EXISTS members;`,
		`SET FOREIGN_KEY_CHECKS = 1;`,
		`CREATE TABLE members (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		`CREATE TABLE projects (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			owner_member_id BIGINT NOT NULL,
			FOREIGN KEY (owner_member_id) REFERENCES members(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		`CREATE TABLE tasks (
			id BIGINT PRIMARY KEY AUTO_INCREMENT,
			project_id BIGINT NOT NULL,
			assignee_member_id BIGINT NOT NULL,
			title VARCHAR(255) NOT NULL,
			status VARCHAR(64) NOT NULL,
			FOREIGN KEY (project_id) REFERENCES projects(id),
			FOREIGN KEY (assignee_member_id) REFERENCES members(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		`INSERT INTO members (id, name) VALUES
			(1, 'Sakura'),
			(2, 'Haru'),
			(3, 'Mio');`,
		`INSERT INTO projects (id, name, owner_member_id) VALUES
			(1, 'Landing Page', 1),
			(2, 'Admin API', 2);`,
		`INSERT INTO tasks (id, project_id, assignee_member_id, title, status) VALUES
			(1, 1, 2, 'Design hero section', 'todo'),
			(2, 1, 3, 'Implement CTA animation', 'doing'),
			(3, 2, 1, 'Build initialize endpoint', 'done');`,
	}

	for _, q := range queries {
		if _, err := tx.ExecContext(ctx, q); err != nil {
			return fmt.Errorf("initialize failed: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (h *Handler) GetFeed(ctx context.Context) (*openapi.FeedResponse, error) {
	query := `
		SELECT
			p.id AS project_id,
			p.name AS project_name,
			owner.name AS owner_name,
			t.id AS task_id,
			t.title AS task_title,
			t.status AS task_status
		FROM tasks t
		INNER JOIN projects p ON p.id = t.project_id
		INNER JOIN members owner ON owner.id = p.owner_member_id
		INNER JOIN members assignee ON assignee.id = t.assignee_member_id
		ORDER BY p.id, t.id
	`

	rows := []feedRow{}
	if err := h.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("select feed: %w", err)
	}

	out := make([]openapi.FeedItem, 0, len(rows))
	for _, row := range rows {
		out = append(out, openapi.FeedItem{
			ProjectID:   row.ProjectID,
			ProjectName: row.ProjectName,
			OwnerName:   row.OwnerName,
			TaskID:      row.TaskID,
			TaskTitle:   row.TaskTitle,
			TaskStatus:  row.TaskStatus,
		})
	}

	return &openapi.FeedResponse{Data: out}, nil
}

func (h *Handler) GetMembers(ctx context.Context) (*openapi.MembersResponse, error) {
	query := `SELECT id, name FROM members ORDER BY id`
	rows := []memberRow{}
	if err := h.db.SelectContext(ctx, &rows, query); err != nil {
		return nil, fmt.Errorf("select members: %w", err)
	}

	out := make([]openapi.Member, 0, len(rows))
	for _, row := range rows {
		out = append(out, openapi.Member{ID: row.ID, Name: row.Name})
	}

	return &openapi.MembersResponse{Data: out}, nil
}

func (h *Handler) CreateTask(ctx context.Context, req *openapi.CreateTaskRequest) error {
	if req == nil || req.Title == "" || req.MemberID == 0 {
		return fmt.Errorf("title and member_id are required")
	}

	query := `INSERT INTO tasks (project_id, assignee_member_id, title, status) VALUES (?, ?, ?, 'todo')`
	if _, err := h.db.ExecContext(ctx, query, 1, req.MemberID, req.Title); err != nil {
		return fmt.Errorf("insert task: %w", err)
	}

	return nil
}
