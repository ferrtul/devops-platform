package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
)

func setupTestDB(t *testing.T) {
	t.Helper()

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		t.Skip("DATABASE_URL is not set")
	}

	ctx := context.Background()

	var err error

	db, err = pgx.Connect(ctx, databaseURL)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	if err := initDB(ctx); err != nil {
		t.Fatalf("failed to initialize test database: %v", err)
	}

	_, err = db.Exec(ctx, "TRUNCATE TABLE users RESTART IDENTITY")
	if err != nil {
		t.Fatalf("failed to clean users table: %v", err)
	}
}

func teardownTestDB(t *testing.T) {
	t.Helper()

	if db == nil {
		return
	}

	ctx := context.Background()

	if err := db.Close(ctx); err != nil {
		t.Errorf("failed to close test database: %v", err)
	}

	db = nil
}

func TestCreateUser(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	body := `{
		"name": "John",
		"email": "john@example.com"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users",
		strings.NewReader(body),
	)

	rec := httptest.NewRecorder()

	createUser(context.Background(), rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusCreated,
			rec.Code,
		)
	}

	var user User

	if err := json.NewDecoder(rec.Body).Decode(&user); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if user.ID != 1 {
		t.Errorf("expected ID 1, got %d", user.ID)
	}

	if user.Name != "John" {
		t.Errorf("expected name %q, got %q", "John", user.Name)
	}

	if user.Email != "john@example.com" {
		t.Errorf(
			"expected email %q, got %q",
			"john@example.com",
			user.Email,
		)
	}
}

func TestGetUsers(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	_, err := db.Exec(
		context.Background(),
		`
		INSERT INTO users(name, email)
		VALUES ($1, $2)
		`,
		"John",
		"john@example.com",
	)

	if err != nil {
		t.Fatalf("failed to insert test user: %v", err)
	}

	rec := httptest.NewRecorder()

	getUsers(context.Background(), rec)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var users []User

	if err := json.NewDecoder(rec.Body).Decode(&users); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(users) != 1 {
		t.Fatalf(
			"expected 1 user, got %d",
			len(users),
		)
	}

	if users[0].Name != "John" {
		t.Errorf(
			"expected name %q, got %q",
			"John",
			users[0].Name,
		)
	}

	if users[0].Email != "john@example.com" {
		t.Errorf(
			"expected email %q, got %q",
			"john@example.com",
			users[0].Email,
		)
	}
}

func TestReadyHandler(t *testing.T) {
	setupTestDB(t)
	defer teardownTestDB(t)

	req := httptest.NewRequest(
		http.MethodGet,
		"/ready",
		nil,
	)

	rec := httptest.NewRecorder()

	readyHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var response struct {
		Status string `json:"status"`
	}

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.Status != "ready" {
		t.Errorf(
			"expected status %q, got %q",
			"ready",
			response.Status,
		)
	}
}
