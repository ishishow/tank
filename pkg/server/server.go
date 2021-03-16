package server

import (
	"20dojo-online/pkg/registry"
	"log"
	"net/http"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	handler := registry.InitializeHandler()
	auth := registry.InitializeAuth()

	http.HandleFunc("/user/create", post(handler.HandleCreate()))
	http.HandleFunc("/user/get", get(auth.Authenticate(handler.HandleGet())))
	http.HandleFunc("/user/update", put(auth.Authenticate(handler.HandleUpdateName())))
	http.HandleFunc("/user/get_status", get(auth.Authenticate(handler.HandleGetStatus())))
	http.HandleFunc("/user/status", put(auth.Authenticate(handler.HandleUpdateStatus())))
	http.HandleFunc("/game/finish", post(auth.Authenticate(handler.HandleFinishGame())))
	http.HandleFunc("/gacha/draw", post(auth.Authenticate(handler.HandleGachaDraw())))
	http.HandleFunc("/collection", get(auth.Authenticate(handler.HandleCollectionList())))
	http.HandleFunc("/ranking", post(handler.HandleRankingList()))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// post POSTリクエストを処理する
func put(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPut)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {

		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")
		writer.Header().Set("Access-Control-Allow-Headers", "*")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// // // プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			writer.WriteHeader(http.StatusOK)
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			writer.Write([]byte("Method Not Allowed"))
			return
		}
		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
