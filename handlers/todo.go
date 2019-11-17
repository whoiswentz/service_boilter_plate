package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) postTodo(w http.ResponseWriter, r *http.Request) {
	h.logger.Printf("%s - %s\n", r.Method, r.URL.Path)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var task Task
	if err := json.Unmarshal(b, &task); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(task.Title))
	log.Println(task)
}
