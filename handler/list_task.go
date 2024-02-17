package handler

import (
	"net/http"

	"github.com/louvre2489/go_todo_app/entity"
	"github.com/louvre2489/go_todo_app/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	tasks := lt.Store.All()
	rsp := []task{}

	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
