package controller

import (
	"MVC/message"
	"MVC/service"
	"encoding/json"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {

	}

	requestData := message.UserRequest{}
	decode := schema.NewDecoder()
	if err := decode.Decode(&requestData, r.Form); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userID, err := strconv.ParseInt(requestData.UserID, 10, 64)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errors.New("bad Parameter userID is number").Error()))
		return
	}

	user, err := service.GetUser(userID)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	outData := message.UserResponse{
		UserID:    user.UserID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	jsonOutData, _ := json.Marshal(outData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonOutData)

}
