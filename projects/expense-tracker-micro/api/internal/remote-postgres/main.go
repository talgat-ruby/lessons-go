package remote_postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	_ "github.com/lib/pq"

	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/config"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/types/database"
)

type Postgres struct {
	conf   *config.RemotePostgresConfig
	logger *slog.Logger
}

func New(conf *config.RemotePostgresConfig, logger *slog.Logger) (*Postgres, error) {
	return &Postgres{
		conf:   conf,
		logger: logger,
	}, nil
}

func (p *Postgres) CreateUser(ctx context.Context, req database.CreateUserReq) (database.CreateUserResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) FindUserByEmail(ctx context.Context, req database.FindUserByEmailReq) (database.FindUserByEmailResp, error) {
	url := fmt.Sprintf("http://%s:%d/users/%s", p.conf.Host, p.conf.Port, req.GetEmail())
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("request error", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("body error", err)
	}

	respBody := &FindUserByEmailResp{}
	if err := json.Unmarshal(body, respBody); err != nil {
		log.Fatal("json unmarshal", err)
	}

	return respBody, nil
}

type FindUserByEmailResp struct {
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash"`
	Salt         string    `json:"salt"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (resp *FindUserByEmailResp) GetID() string {
	return resp.ID
}

func (resp *FindUserByEmailResp) GetEmail() string {
	return resp.Email
}

func (resp *FindUserByEmailResp) GetPasswordHash() string {
	return resp.PasswordHash
}

func (resp *FindUserByEmailResp) GetSalt() string {
	return resp.Salt
}

func (resp *FindUserByEmailResp) GetCreatedAt() time.Time {
	return resp.CreatedAt
}

func (resp *FindUserByEmailResp) GetUpdatedAt() time.Time {
	return resp.UpdatedAt
}

func (p *Postgres) ListExpense(ctx context.Context, req database.ListExpenseReq) (database.ListExpenseResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) CreateExpense(ctx context.Context, req database.CreateExpenseReq) (database.CreateExpenseResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) UpdateExpense(ctx context.Context, req database.UpdateExpenseReq) (database.UpdateExpenseResp, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Postgres) DeleteExpense(ctx context.Context, req database.DeleteExpenseReq) (database.DeleteExpenseResp, error) {
	//TODO implement me
	panic("implement me")
}
