package main

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/louvre2489/go_todo_app/clock"
	"github.com/louvre2489/go_todo_app/config"
	"github.com/louvre2489/go_todo_app/handler"
	"github.com/louvre2489/go_todo_app/service"
	"github.com/louvre2489/go_todo_app/store"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 静的解析のエラーを回避するために明示的に戻り値を捨てる
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	db, cleanup, err := store.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}

	r := store.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Service:   &service.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTTP)
	lt := &handler.ListTask{
		Service: &service.ListTask{
			DB:   db,
			Repo: &r,
		},
	}
	mux.Get("/tasks", lt.ServeHTTTP)

	return mux, cleanup, nil
}
