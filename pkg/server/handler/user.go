package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/pkg/dcontext"
	"20dojo-online/pkg/http/response"
)

// HandleCreate ユーザを作成するHandler
func (h Handler) HandleCreate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody userCreateRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// データベースにユーザデータを登録する
		authToken, err := h.UserUseCase.Create(requestBody.Name)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// 生成した認証トークンを返却
		response.Success(writer, &userCreateResponse{Token: authToken})
	}
}

// HandleGet ユーザー取得処理
func (h Handler) HandleGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		user, err := h.UserUseCase.Get(userID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// レスポンスに必要な情報を詰めて返却
		response.Success(writer, &userGetResponse{
			ID:        user.ID,
			Name:      user.Name,
			HighScore: user.HighScore,
			Coin:      user.Coin,
		})
	}
}

// HandleUpdate ユーザー更新処理
func (h Handler) HandleUpdateName() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody userUpdateNameRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		if err := h.UserUseCase.UpdateName(userID, requestBody.Name); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		response.Success(writer, nil)
	}
}

// HandleUpdate ユーザー更新処理
func (h Handler) HandleUpdateStatus() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody userUpdateStatusRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		if err := h.UserUseCase.UpdateStatus(userID, requestBody.WeaponID, requestBody.SkinID); err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		response.Success(writer, nil)
	}
}

// HandleUpdate ユーザー更新処理
func (h Handler) HandleGetStatus() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		user_status, err := h.UserUseCase.GetStatus(userID)
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}
		response.Success(writer, user_status)
	}
}

type userGetResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	HighScore int32  `json:"highScore"`
	Coin      int32  `json:"coin"`
}

type userCreateRequest struct {
	Name string `json:"name"`
}

type userCreateResponse struct {
	Token string `json:"token"`
}

type userUpdateNameRequest struct {
	Name string `json:"name"`
}

type userUpdateStatusRequest struct {
	WeaponID string `json:"weapon_id"`
	SkinID   string `json:"skin_id"`
}


