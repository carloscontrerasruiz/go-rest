package routes

import (
	"encoding/json"
	"net/http"

	"github.com/carloscontrerasruiz/gorest/models"
	"github.com/carloscontrerasruiz/gorest/repository"
	"github.com/carloscontrerasruiz/gorest/server"
)

type SingUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

type SingUpResponse struct {
	Email string `json:"email"`
	Id    int64  `json:"id"`
}

func SignUpHander(s server.Server, repo repository.CrudRepository[models.User]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req = SingUpRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var user = models.User{
			Email:    req.Email,
			Password: req.Password,
		}

		response, err := repo.Create(r.Context(), user)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SingUpResponse{
			Id:    response.Id,
			Email: response.Email,
		})

	}

}
