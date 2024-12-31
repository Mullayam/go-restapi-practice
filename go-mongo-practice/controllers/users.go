package controllers

import (
	"encoding/json"
	"net/http"

	_ "enjoys.in/mongo-go-test/config"
	"enjoys.in/mongo-go-test/models"
	"enjoys.in/mongo-go-test/routes"
	types "enjoys.in/mongo-go-test/types"
	utils "enjoys.in/mongo-go-test/utils"
	"gopkg.in/mgo.v2/bson"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Handler"))
	utils.WriteJson(w, http.StatusOK, "User Handler")
}
func AddUsers(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	json.NewDecoder(r.Body).Decode(&user)
	user.Id = bson.NewObjectId()
	// config.DB.Session.DB("test").C("users").Insert(u)

	utils.WriteJson(w, http.StatusOK, user)
}
func UpdateUsers(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&models.User{})
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User Handler"))
	utils.WriteJson(w, http.StatusOK, "User Handler")
}
func (h *routes.Controller) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var payload = &types.JsonResponse{}

	if !bson.IsObjectIdHex(id) {
		payload.Success = false
		payload.Message = "Invalid User ID"
		payload.Result = nil
		utils.WriteJson(w, http.StatusNotFound, payload)
		return
	}

	var oid = bson.ObjectIdHex(id)
	u := models.User{}
	h.Session.DB("test").C("users").FindId(oid).One(&u)
	utils.WriteJson(w, http.StatusOK, payload)

}
