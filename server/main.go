package main

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	"strings"
	"strconv"

	"github.com/yanple/vk_api"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
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

type PhotoResp struct {
	Response struct {
		Items []struct {
			MessageId int `json:"message_id"`
			Attachment struct {
				Type string `json:"type"`
				Photo struct {
					Id        int `json:"id"`
					AlbumId   int `json:"album_id"`
					OwnerId   int `json:"owner_id"`
					Photo75   string `json:"photo_75"`
					Photo130  string `json:"photo_130"`
					Photo604  string `json:"photo_604"`
					Photo807  string `json:"photo_807"`
					Photo1280 string `json:"photo_1280"`
					Photo2560 string `json:"photo_2560"`
					Width     int `json:"width"`
					Height    int `json:"height"`
					Text      string `json:"text"`
					Date      int `json:"date"`
					AccessKey string `json:"access_key"`
				} `json:"photo"`
			} `json:"attachment"`
		} `json:"items"`
		NextFrom string `json:"next_from"`
	} `json:"response"`
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
		log.Println(err)
		return
	}

	fmt.Fprint(w, string(mapJson))
}

func getListPhotos(api vk_api.Api, userId string) ([]string, error) {
	result := make([]string, 0, 2000)
	params := make(map[string]string)
	params["peer_id"] = userId
	params["media_type"] = "photo"
	params["count"] = "200"
	params["photos"] = "5.62"
	params["offset"] = "0"

	var photos PhotoResp
	temp, err := api.Request("messages.getHistoryAttachments", params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(temp), &photos)
	if err != nil {
		return nil, err
	}
	for _, item := range photos.Response.Items {
		log.Println("item")
		if item.Attachment.Photo.Width > 1280 {
			result = append(result, item.Attachment.Photo.Photo1280)
		} else if item.Attachment.Photo.Width > 807 {
			result = append(result, item.Attachment.Photo.Photo807)
		} else if item.Attachment.Photo.Width > 604 {
			result = append(result, item.Attachment.Photo.Photo604)
		} else if item.Attachment.Photo.Width > 130 {
			result = append(result, item.Attachment.Photo.Photo130)
		} else if item.Attachment.Photo.Width > 75 {
			result = append(result, item.Attachment.Photo.Photo75)
		} else {
			log.Println("too small image", item.Attachment.Photo.Width)
		}
	}
	return result, nil
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

	photos, err := getListPhotos(api, "42690043")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(photos)

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
	//http.HandleFunc("/getConversations/", getListPhotos)

	log.Println("Server started at port :4100")
	err := http.ListenAndServe(":4100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
