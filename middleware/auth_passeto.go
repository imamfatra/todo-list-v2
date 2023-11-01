package middleware

// import (
// 	"net/http"
// 	"todo-api/helper"
// 	"todo-api/model"
// )

// type AuthMiddleware struct {
// 	Handler http.Handler
// }

// func NewAuthMiddlerware(handler http.Handler) *AuthMiddleware {
// 	return &AuthMiddleware{
// 		Handler: handler,
// 	}
// }

// func (auth *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	config, err := model.LoadConfig("../")
// 	if err != nil {
// 		webResponse := model.WebResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "UNAUTHORIZED",
// 			Data:   err,
// 		}
// 		helper.WriteToResponse(w, webResponse)
// 	}

// 	tokenString := r.Header.Get("authorization")
// 	if err != nil {
// 		webResponse := model.WebResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "UNAUTHORIZED",
// 			Data:   err,
// 		}
// 		helper.WriteToResponse(w, webResponse)
// 	}
// 	maker, err := NewPasetoMaker(config.TokenSymmetricKey)
// 	if err != nil {
// 		webResponse := model.WebResponse{
// 			Code:   http.StatusInternalServerError,
// 			Status: "Interneal Server Error",
// 			Data:   err,
// 		}
// 		helper.WriteToResponse(w, webResponse)
// 	}

// 	_, err = maker.VarifyToken(tokenString)
// 	if err != nil {
// 		webResponse := model.WebResponse{
// 			Code:   http.StatusUnauthorized,
// 			Status: "UNAUTHORIZED",
// 			Data:   err,
// 		}
// 		helper.WriteToResponse(w, webResponse)
// 	}

// 	auth.Handler.ServeHTTP(w, r)
// }

// func Auth(next http.HandlerFunc) http.HandlerFunc {
// 	return http.HandleFunc(func(w http.ResponseWriter, r *http.Request) {
// 			config, err := model.LoadConfig("../")
// 			if err != nil {
// 				webResponse := model.WebResponse{
// 					Code:   http.StatusUnauthorized,
// 					Status: "UNAUTHORIZED",
// 					Data:   err,
// 				}
// 				helper.WriteToResponse(w, webResponse)
// 			}

// 			tokenString := r.Header.Get("authorization")
// 			if err != nil {
// 				webResponse := model.WebResponse{
// 					Code:   http.StatusUnauthorized,
// 					Status: "UNAUTHORIZED",
// 					Data:   err,
// 				}
// 				helper.WriteToResponse(w, webResponse)
// 			}
// 			maker, err := NewPasetoMaker(config.TokenSymmetricKey)
// 			if err != nil {
// 				webResponse := model.WebResponse{
// 					Code:   http.StatusInternalServerError,
// 					Status: "Interneal Server Error",
// 					Data:   err,
// 				}
// 				helper.WriteToResponse(w, webResponse)
// 			}

// 			_, err = maker.VarifyToken(tokenString)
// 			if err != nil {
// 				webResponse := model.WebResponse{
// 					Code:   http.StatusUnauthorized,
// 					Status: "UNAUTHORIZED",
// 					Data:   err,
// 				}
// 				helper.WriteToResponse(w, webResponse)
// 			}

// 			// http.Handler.ServeHTTP(w, r)
// 		}
// 	)
// }
