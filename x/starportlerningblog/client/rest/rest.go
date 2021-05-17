package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers starportlerningblog-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/starportlerningblog/comments/{id}", getCommentHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/starportlerningblog/comments", listCommentHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/starportlerningblog/comments", createCommentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/starportlerningblog/comments/{id}", updateCommentHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/starportlerningblog/comments/{id}", deleteCommentHandler(clientCtx)).Methods("POST")

}
