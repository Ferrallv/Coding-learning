
package controllers

import (
	"encoding/json"
	"fmt"
	"ExercisesAndProjectsInC/GO/go_webdev/hands-on-exc/008/starting-code/models"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/google/uuid"
	"io/ioutil"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(m map[string]models.User) *UserController {
	return &UserController{m}
}

func (uc UserController) GetUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// composite literal
	u, ok := uc.session[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(req.Body).Decode(&u)

	u.Id = uuid.Must(uuid.NewRandom()).String()
	
	// store the user in mongodb
	uc.session[u.Id] = u

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)

	
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	// Delete user
	delete(uc.session, id)

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", id, "\n")
}