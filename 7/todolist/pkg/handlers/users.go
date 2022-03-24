package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/rRepZ/ProgrammingCourse/7/todolist/users"
)

type UserHandler struct {
	repo *users.UserRepo
}

func (uh *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ioutil.ReadAll(r.Body)
}
