package yahoo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GetTokenRequest struct {
	Code string `json:"code"`
}

func AuthorizeYahooHandler(svc YahooService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Redirecting user to Authorization page")
		u := svc.GetRedirectURL()
		json.NewEncoder(w).Encode(u)
	}
}

// func AuthorizeYahoo(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Redirecting user to Authorization page")
// 	u := GetRedirectURL()
// 	json.NewEncoder(w).Encode(u)
// }

func GetTokenHandler(svc YahooService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Exchanging Auth code for token")
		// extract the code from request
		decoder := json.NewDecoder(r.Body)
		var body GetTokenRequest
		err := decoder.Decode(&body)
		if err != nil {
			fmt.Println(err)
			return
		}

		token := svc.GetYahooToken(body.Code)
		json.NewEncoder(w).Encode(token)
	}
}
