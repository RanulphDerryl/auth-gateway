package auth_gateway

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

func DecodeRequest(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err!= nil {
		return err
	}
	return validate.Struct(v)
}

func GetHeaderInt(r *http.Request, key string) (int, error) {
	header := r.Header.Get(key)
	if header == "" {
		return 0, nil
	}
	val, err := strconv.Atoi(header)
	if err!= nil {
		return 0, err
	}
	return val, nil
}

func GetHeaderString(r *http.Request, key string) (string, error) {
	return r.Header.Get(key), nil
}

func GetQueryInt(r *http.Request, key string) (int, error) {
	val := r.URL.Query().Get(key)
	if val == "" {
		return 0, nil
	}
	val, err := strconv.Atoi(val)
	if err!= nil {
		return 0, err
	}
	return val, nil
}

func GetQueryStrings(r *http.Request, key string) ([]string, error) {
	return strings.Split(r.URL.Query().Get(key), ","), nil
}

func ValidateUUID(s string) bool {
	_, err := time.Parse(time.RFC3339, s)
	return err == nil
}

func ValidatePassword(password string) bool {
	return len(password) >= 8 && strings.Count(password, string(rune(0x20))) > 0
}

func SetResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func GetDefaultHandler(r *http.Request) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		SetResponse(w, http.StatusNotFound, nil)
	}
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func NewRouter() *mux.Router {
	return mux.NewRouter()
}