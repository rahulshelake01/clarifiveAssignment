package router

import (
	"clarifiveassignment/api/v1/services"
	"clarifiveassignment/api/v1/utils"
	"clarifiveassignment/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type RouterStruct struct {
	Family services.FamilyServiceInterface
}

func RouteHandler(familyService services.FamilyServiceInterface) *mux.Router {
	handler := &RouterStruct{Family: familyService}

	router := mux.NewRouter()
	router.HandleFunc("/add", handler.AddMember).Methods("POST")
	router.HandleFunc("/search", handler.GetMembers).Methods("GET")

	return router
}

func (ph RouterStruct) AddMember(w http.ResponseWriter, r *http.Request) {

	var request models.AddMemberRequest
	json.NewDecoder(r.Body).Decode(&request)
	response := ph.Family.Add(&request)
	utils.ResponseJson(w, response)

}

func (ph RouterStruct) GetMembers(w http.ResponseWriter, r *http.Request) {

	filter := new(models.GetMemberRequest)
	r.ParseForm()
	schema.NewDecoder().Decode(filter, r.Form)
	response := ph.Family.Get(filter)
	utils.ResponseJson(w, response)

}
