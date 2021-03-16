package handler

import (
	"log"
	"net/http"

	"20dojo-online/pkg/dcontext"
	"20dojo-online/pkg/http/response"
)

// HandleCollectionList コレクションリスト情報の取得ハンドラ
func (h Handler) HandleCollectionList() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		userID := dcontext.GetUserIDFromContext(ctx)
		collectionResponse, err := h.CollectionItemUseCase.SelectALLUserColectionItems(userID)
		if err != nil {
			log.Printf("cannot get userCollectionItems: %v", err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		response.Success(writer, collectionResponse)
	}
}
