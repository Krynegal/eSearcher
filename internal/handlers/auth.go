package handlers

import (
	"eSearcher/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func (r *Router) registration(w http.ResponseWriter, req *http.Request) {
	ct := req.Header.Get("Content-Type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong data format"))
		return
	}
	var user models.User
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = r.readingUserData(w, req, user)
	if err != nil {
		return
	}

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	fmt.Printf("user: %+v", user)
	var userID int
	userID, err = r.Services.AuthService.CreateUser(user.Login, user.Password, user.Role)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.writeToken(w, userID, user.Role)
}

func (r *Router) authentication(w http.ResponseWriter, req *http.Request) {
	ct := req.Header.Get("Content-Type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Wrong data format"))
		return
	}
	var u models.User
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(body, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = r.readingUserData(w, req, u)
	if err != nil {
		return
	}

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	user, err := r.Services.AuthService.AuthUser(u.Login, u.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	r.writeToken(w, user.ID, user.Role)
}

func (r *Router) readingUserData(w http.ResponseWriter, req *http.Request, user models.User) error {
	if user.Login == "" || user.Password == "" {
		http.Error(w, "empty login or password", http.StatusBadRequest)
		return errors.New("empty login or password")
	}
	return nil
}

func (r *Router) writeToken(w http.ResponseWriter, uid int, role int) {
	token, err := r.Services.AuthService.GenerateToken(uid, role)
	if err != nil {
		http.Error(w, "Can't get token", http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "token",
		Value: token,
	})
}

func (r *Router) getUserIDFromToken(w http.ResponseWriter, req *http.Request) (int, error) {
	ID := w.Header().Get("user_id")
	userID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "can't get user ID", http.StatusInternalServerError)
		return 0, err
	}
	return userID, nil
}

func (r *Router) getUserRoleFromToken(w http.ResponseWriter, req *http.Request) (int, error) {
	ID := w.Header().Get("user_role")
	userID, err := strconv.Atoi(ID)
	if err != nil {
		http.Error(w, "can't get user role", http.StatusInternalServerError)
		return 0, err
	}
	return userID, nil
}
