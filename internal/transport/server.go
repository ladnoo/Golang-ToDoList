package transport

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	handler *HTTPHandlers
}

func NewServer(handler *HTTPHandlers) *Server {
	return &Server{handler: handler}
}

func (s *Server) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.handler.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.handler.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.handler.HandleGetUncompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.handler.HandleGetAllTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.handler.HandleCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.handler.HandleDeleteTask)

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}

	return nil
}
