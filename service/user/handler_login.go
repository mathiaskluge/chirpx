package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mathiaskluge/chirpx/config"
	"github.com/mathiaskluge/chirpx/service/auth"
	"github.com/mathiaskluge/chirpx/types"
	"github.com/mathiaskluge/chirpx/utils"
)

func (h *Handler) handlerLoginUser(w http.ResponseWriter, req *http.Request) {
	// Parse payload
	var payload types.LoginUserPayload
	if err := utils.ParseJSON(req, &payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err)
		return
	}

	// Validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	//Check if user exists and get it
	user, err := h.store.GetUserByEmail(payload.Email)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid email or password"))
		return
	}

	//Compare password hash to payload
	if !auth.ComparePasswords(user.PwHash, []byte(payload.Password)) {
		utils.RespondWithError(w, http.StatusBadRequest, fmt.Errorf("user not found, invalid email or password"))
		return
	}

	//Generate access token with a 1 hour lifespan
	token, err := auth.CreateJWT(config.Env.JWTSecret, user.ID, 60*60)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	//Generate a new SessionID
	sToken, err := auth.GenerateSessionID()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	// Create the session in the DB with a 60 day lifespan
	err = h.store.CreateSession(sToken, user.ID, 60*60*24*60)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err)
	}

	//Respond with token
	utils.RespondWithJSON(w, http.StatusOK, types.LoginUserResponse{
		ID:      user.ID,
		Email:   user.Email,
		Token:   token,
		Session: sToken,
	})
}
