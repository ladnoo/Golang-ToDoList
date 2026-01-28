package httpserver

import (
	"ToDoList/todo"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type HTTPHandlers struct {
	toDoList *todo.List
}

func NewHTTPHandlers(list *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		list,
	}
}

func handleErrorNotFound(err error, w http.ResponseWriter) {
	errDTO := NewErrorDTO(err)
	if errors.Is(err, todo.ErrTaskNotFound) {
		http.Error(w, errDTO.ToString(), http.StatusNotFound)
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
	}
}

func handleErrorConflict(err error, w http.ResponseWriter) {
	errDTO := NewErrorDTO(err)
	if errors.Is(err, todo.ErrTaskAlreadyExists) {
		http.Error(w, errDTO.ToString(), http.StatusConflict)
	} else {
		http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
	}
}

/* HandleCreateTask
pattern: /tasks
method: POST
info: JSON in HTTP request body

succeed:
	- status-code - 201
	- response body: JSON represented created todo
failed:
	- status-code: 400, 409 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := taskDTO.ValidateToCreate(); err != nil {
		errDTO := NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	toDoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.toDoList.AddTask(toDoTask); err != nil {
		handleErrorConflict(err, w)
		return
	}

	b, err := json.MarshalIndent(toDoTask, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write HTTP response: ", err)
		return
	}
}

/* HandleGetTask
   pattern: /tasks/{title}
   method: GET
   info: pattern

   succeed:
   	- status-code - 200
   	- response body: JSON represented existing todo
   failed:
   	- status-code: 400, 404 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.toDoList.GetTask(title)
	if err != nil {
		handleErrorNotFound(err, w)
		return
	}

	b, err := json.MarshalIndent(task, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write HTTP response: ", err)
		return
	}
}

/* HandleGetAllTasks
   pattern: /tasks
   method: GET
   info: -

   succeed:
   	- status-code - 200
   	- response body: JSON represent list of existing tasks
   failed:
   	- status-code: 400 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.toDoList.GetTasks()

	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write HTTP response: ", err)
		return
	}
}

/* HandleGetUncompletedTasks
   pattern: /tasks?completed=false
   method: GET
   info: query params

   succeed:
   	- status-code - 200
   	- response body: JSON represent list of found tasks with filter
   failed:
   	- status-code: 400 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleGetUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.toDoList.GetUncompletedTasks()

	b, err := json.MarshalIndent(tasks, "", "	")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write HTTP response: ", err)
		return
	}
}

/* HandleCompleteTask
   pattern: /tasks/{title}
   method: PATCH
   info: pattern and JSON with changed todo's field

   succeed:
   	- status-code - 200
   	- response body: JSON represent edited todo
   failed:
   	- status-code: 400, 409 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		errDTO := NewErrorDTO(err)
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if completeDTO.Complete {
		changedTask, err = h.toDoList.CompleteTask(title)
	} else {
		changedTask, err = h.toDoList.UncompleteTask(title)
	}

	if err != nil {
		handleErrorNotFound(err, w)
		return
	}

	b, err := json.MarshalIndent(changedTask, "", "	")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("Failed to write HTTP response: ", err)
		return
	}
}

/* HandleDeleteTask
   pattern: /tasks/{title}
   method: DELETE
   info: pattern

   succeed:
   	- status-code - 204
   	- response body: empty
   failed:
   	- status-code: 400, 404 or 500
	- response body: JSON with error and time
*/

func (h *HTTPHandlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.toDoList.DeleteTask(title); err != nil {
		handleErrorNotFound(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
