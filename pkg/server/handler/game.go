package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/pkg/dcontext"
	"20dojo-online/pkg/http/response"
)

type gameFinishRequest struct {
	Score int32 `json:"score"`
}

type gameFinishResponse struct {
	Coin int32 `json:"coin"`
}

// HandleFinishGame インゲーム終了API
func (uh Handler) HandleFinishGame() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// リクエストBodyから更新後情報を取得
		var requestBody gameFinishRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// リクエストが負の時のバリデーション
		if requestBody.Score < 0 {
			log.Println("score is not positive")
			response.BadRequest(writer, "your score-parameter is not positive")
			return
		}

		// Contextから認証済みのユーザIDを取得
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		coin, err := uh.GameUseCase.FinishGame(userID, requestBody.Score)
		if err != nil {
			log.Println("score is not positive")
			response.BadRequest(writer, "your score-parameter is not positive")
			return
		}
		log.Println("game is finished")
		response.Success(writer, &gameFinishResponse{Coin: coin})
	}
}
