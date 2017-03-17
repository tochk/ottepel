package main

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"

	"github.com/yanple/vk_api"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

type tokenUrl struct {
	Url string
}

type AuthAnswer struct {
	Code int
	Url  string
}

type TokenAnswer struct {
	Code        int
	Status      string
	AccessToken string
	ExpiresIn   int
	UserId      int
	UserName    string
}

type UserData struct {
	AccessToken string
	ExpiresIn   int
	UserId      int
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	var api vk_api.Api
	authUrl, err := api.GetAuthUrl(
		"https://api.vk.com/blank.html",
		"5930862",
		"friends,messages,offline")
	if err != nil {
		log.Println(err)
		return
	}
	mapJson, err := json.Marshal(AuthAnswer{Code: 200, Url: authUrl, })
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func getConversationsHandler(w http.ResponseWriter, r *http.Request) {
	v := make(map[string]interface{})
	var t UserData
	var api vk_api.Api
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}
	api.ExpiresIn = t.ExpiresIn
	api.AccessToken = t.AccessToken
	api.UserId = t.UserId

	params := make(map[string]string)
	params["count"] = "200"
	api.Request("messages.getDialogs", params)
	temp, err := api.Request("messages.getHistoryAttachments", params)
	if err != nil {
		log.Printf("Error query: %s", err)
		return
	}
	log.Println(temp)
	err = json.Unmarshal([]byte(temp), &v)
	if err != nil {
		log.Printf("Error Unmarshall: %s", err)
		return
	}
	log.Println(v)
}

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	var t tokenUrl
	var api vk_api.Api
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}
	tokenAnswer := TokenAnswer{Code: 200, Status: "Ok", }
	splitUrl := strings.Split(strings.Split(t.Url, "#")[1], "&")
	for _, urlPart := range splitUrl {
		tempUrlForMap := strings.Split(urlPart, "=")
		if tempUrlForMap[0] == "access_token" {
			tokenAnswer.AccessToken = tempUrlForMap[1]
			api.AccessToken = tempUrlForMap[1]
		} else if tempUrlForMap[0] == "expires_in" {
			tempInt, err := strconv.Atoi(tempUrlForMap[1])
			if err != nil {
				log.Printf("Error parsing data: %s", err)
				return
			}
			tokenAnswer.ExpiresIn = tempInt
			api.ExpiresIn = tempInt
		} else if tempUrlForMap[0] == "user_id" {
			tempInt, err := strconv.Atoi(tempUrlForMap[1])
			if err != nil {
				log.Printf("Error parsing data: %s", err)
				return
			}
			tokenAnswer.UserId = tempInt
			api.UserId = tempInt
		}
	}
	params := make(map[string]string)
	params["peer_id"] = "42690043"
	params["media_type"] = "photo"
	params["count"] = "1"

	temp, err := api.Request("messages.getHistoryAttachments", params)
	if err != nil {
		log.Printf("Error query: %s", err)
		return
	}
	log.Println(temp)
	mapJson, err := json.Marshal(tokenAnswer)
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func getPhotos(api vk_api.Api) {

}

func main() {
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/token/", tokenHandler)
	http.HandleFunc("/getConversations/", getConversationsHandler)

	log.Println("Server started at port :4100")
	err := http.ListenAndServe(":4100", nil)
	if err != nil {
		log.Fatal(err)
	}
}