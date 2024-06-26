package htmx

import (
	"net/http"

	"github.com/jstamariz/go-htmx/internal/core/domain"
	"github.com/jstamariz/go-htmx/internal/core/services/todosrv"
)

func (hx *HTMXHandler) AddHandleFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusInternalServerError)
		return
	}

	item := domain.TodoItem{
		Title:       r.Form.Get("title"),
		Description: r.Form.Get("description"),
	}

	err = hx.srv.Create(&item)
	if err != nil {
		if err == todosrv.ErrorTitleRequired {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	err = hx.tmpl.ExecuteTemplate(w, "list_item.html", ToView(item))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
