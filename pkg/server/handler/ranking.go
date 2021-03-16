package handler

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	rm "20dojo-online/pkg/domain/model/ranking"

	"20dojo-online/pkg/http/response"
)

// HandleRankingList ランキング一覧取得処理
func (h Handler) HandleRankingList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var requestBody RankingRequest
		if err := json.NewDecoder(request.Body).Decode(&requestBody); err != nil {
			log.Println(err)
			response.BadRequest(writer, "Content-Type must be application/json")
			return
		}

		fmt.Println(requestBody.Start)
		if requestBody.Start <= 0 {
			log.Println("start must be positive")
			response.BadRequest(writer, "start must be positive")
			return
		}

		rankInfoList, err := h.RankingUseCase.GetRankingList(int64(requestBody.Start))
		if err != nil {
			log.Println("start must be p")
			response.BadRequest(writer, "start must be positive")
			return
		}

		if len(rankInfoList) == 0 {
			log.Println("users is empty")
		}

		response.Success(writer, &RankingResponse{
			Ranks: rankInfoList,
		})
	}
}

type RankingRequest struct {
	Start int32 `json:"start"`
}

type RankingResponse struct {
	Ranks []*rm.RankInfo `json:"ranks"`
}