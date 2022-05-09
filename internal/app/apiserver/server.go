package apiserver

import (
	"Diplom/internal/app/apifilesystem"
	"Diplom/internal/app/model"
	"Diplom/internal/app/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type server struct {
	router           *mux.Router
	store            store.Store
	fileSystemConfig *apifilesystem.ConfigDirectories
}

func newServer(store store.Store, configFileSystem *apifilesystem.ConfigDirectories) *server {
	s := &server{
		router:           mux.NewRouter(),
		store:            store,
		fileSystemConfig: configFileSystem,
	}

	s.configureRouter()

	return s
}

func (s *server) serverHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/create_users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods("POST")
	s.router.HandleFunc("/create_group", s.handleCreateGroup()).Methods("POST")
	s.router.HandleFunc("/delete_group", s.handleDeleteGroup()).Methods("POST")
	s.router.HandleFunc("/create_file", s.handleCreateFile()).Methods("POST")
	s.router.HandleFunc("/delete_file", s.handleDeleteFile()).Methods("POST")
	s.router.HandleFunc("/get_file", s.handleGetFile()).Methods("GET")
	s.router.HandleFunc("/add_user_in_group", s.handleAddUserInGroup()).Methods("POST")
}

func (s *server) handleAddUserInGroup() http.HandlerFunc {
	type request struct {
		UserID    int    `json:"user_id"`
		GroupName string `json:"group_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		groupModel, err := s.store.Group().FindGroupByName(req.GroupName)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.UserInGroup().AddUserInGroup(req.UserID, groupModel.ID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleGetFile() http.HandlerFunc {
	type request struct {
		UserID    int    `json:"user_id"`
		FileQuery string `json:"file_query"`
	}

	type response struct {
		ID        int
		FileQuery string
		FileName  string
		FileData  []byte
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		fileModel, err := s.store.File().FindByQuery(req.FileQuery)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		data, err := apifilesystem.ReadFile(fileModel.FilePath)

		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
		}

		responseUser := &response{
			FileData:  data,
			FileQuery: req.FileQuery,
			ID:        req.UserID,
			FileName:  fileModel.FileName,
		}
		s.respond(w, r, http.StatusOK, responseUser)
	}
}

func (s *server) handleDeleteFile() http.HandlerFunc {
	type request struct {
		UserID    int    `json:"user_id"`
		FileQuery string `json:"file_query"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.File().UpdateAvailableFile(req.FileQuery, false); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := apifilesystem.DeleteFile(req.FileQuery); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.File().Delete(req.FileQuery, req.UserID); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleCreateFile() http.HandlerFunc {
	type request struct {
		UserID   int    `json:"user_id"`
		FileName string `json:"file_name"`
		FileData []byte `json:"file_data"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		currentDir := s.fileSystemConfig.CurrentDirectoryPath
		if err := apifilesystem.SaveFileInDirectory(req.FileData); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		fileModel := &model.File{
			FileOwner:     req.UserID,
			FileName:      req.FileName,
			FilePath:      currentDir,
			FileQuery:     currentDir,
			FileAvailable: true,
		}

		if err := s.store.File().Create(fileModel); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleDeleteGroup() http.HandlerFunc {
	type request struct {
		UserID    int    `json:"user_id"`
		GroupName string `json:"group_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		if err := s.store.Group().Delete(req.UserID, req.GroupName); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) handleCreateGroup() http.HandlerFunc {
	type request struct {
		GroupOwner int    `json:"group_owner"`
		GroupName  string `json:"group_name"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		gr := &model.Group{
			GroupOwner: req.GroupOwner,
			GroupName:  req.GroupName,
		}

		if err := s.store.Group().Create(gr); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, gr)
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u := &model.User{
			Email:    req.Email,
			Password: req.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleSessionsCreate() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		u, err := s.store.User().FindByEmail(req.Email)
		if err != nil || !u.ComparePassword(req.Password) {
			s.error(w, r, http.StatusUnauthorized, store.ErrIncorrectEmailOrPassword)
			return
		}
		s.respond(w, r, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
