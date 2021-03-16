package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"20dojo-online/pkg/dcontext"
	"20dojo-online/pkg/http/response"
	"20dojo-online/pkg/usecase/interactor"
)

type gachaDrawRequest struct {
	Times int32 `json:"times"`
}

type gachaDrawResponse struct {
	Results []*interactor.GachaResult `json:"results"`
}

type gachaResult struct {
	CollectionID string `json:"item_id"`
	Name         string `json:"name"`
	Rarity       int32  `json:"rarity"`
	IsNew        bool   `json:"is_new"`
}

// HandleGachaDraw ガチャ実行処理
func (h Handler) HandleGachaDraw() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestBody gachaDrawRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Printf("failed json decode : %v", err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		// リクエストが負の時のバリデーション
		if requestBody.Times <= 0 {
			log.Println("times is not positive")
			response.BadRequest(writer, "your times-parameter is not positive")
			return
		}

		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)

		// 実行、保存処理
		gachaResults, err := h.GachaUseCase.GachaDraw(ctx, userID, requestBody.Times)
		if err != nil {
			log.Printf("GachaDraw is failed: %v", err)
			response.InternalServerError(writer, "server error: save collection item is failed")
			return
		}

		response.Success(writer, &gachaDrawResponse{
			Results: gachaResults,
		})
	}
}
