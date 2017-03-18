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
	"time"
	"os"
	"io"
)

type QueueItem struct {
	UrlList     []string
	AccessToken string
	Token       string
	Completed   bool
}

var Queue struct {
	CurrentJobId int
	Items        []QueueItem
	MapTokenId   map[string]int
	Free         bool
}

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

type UserForPhoto struct {
	AccessToken string
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

type Photos struct {
	Photos []string
}

type GetPhotosRequest struct {
	AccessToken string
	Photos      []string
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
	params["v"] = "5.62"
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
	prev := ""
	for len(photos.Response.Items) > 0 {
		log.Println(len(photos.Response.Items))
		if prev != photos.Response.NextFrom {
			for _, item := range photos.Response.Items {
				if item.Attachment.Photo.Photo2560 != "" {
					if strings.Split(item.Attachment.Photo.Photo2560, ".")[len(strings.Split(item.Attachment.Photo.Photo2560, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo2560)
						continue
					}
				}
				if item.Attachment.Photo.Photo1280 != "" {
					if strings.Split(item.Attachment.Photo.Photo1280, ".")[len(strings.Split(item.Attachment.Photo.Photo1280, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo1280)
						continue
					}
				}
				if item.Attachment.Photo.Photo807 != "" {
					if strings.Split(item.Attachment.Photo.Photo807, ".")[len(strings.Split(item.Attachment.Photo.Photo807, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo807)
						continue
					}
				}
				if item.Attachment.Photo.Photo604 != "" {
					if strings.Split(item.Attachment.Photo.Photo604, ".")[len(strings.Split(item.Attachment.Photo.Photo604, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo604)
						continue
					}
				}
				if item.Attachment.Photo.Photo130 != "" {
					if strings.Split(item.Attachment.Photo.Photo130, ".")[len(strings.Split(item.Attachment.Photo.Photo130, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo130)
						continue
					}
				}
				if item.Attachment.Photo.Photo75 != "" {
					if strings.Split(item.Attachment.Photo.Photo75, ".")[len(strings.Split(item.Attachment.Photo.Photo75, "."))-1] != "gif" {
						result = append(result, item.Attachment.Photo.Photo75)
						continue
					}
				} else {
					log.Println("An error was occured", item.Attachment.Photo.Width)
					log.Println("Item trace: ", item)
				}
			}
		} else {
			time.Sleep(time.Millisecond * 300)
		}
		params["start_from"] = photos.Response.NextFrom
		log.Println(photos.Response.NextFrom)
		temp, err = api.Request("messages.getHistoryAttachments", params)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(temp), &photos)
		if err != nil {
			return nil, err
		}
		prev = params["start_from"]
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

	/*temp, _ := getListPhotos(api, "42690043")
	go downloadFiles(temp)*/

	mapJson, err := json.Marshal(tokenAnswer)
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func photosHandler(w http.ResponseWriter, r *http.Request) {
	var userForPhoto UserForPhoto
	var api vk_api.Api
	err := json.NewDecoder(r.Body).Decode(&userForPhoto)
	if err != nil {
		log.Println(err)
		return
	}
	api.AccessToken = userForPhoto.AccessToken
	photos, err := getListPhotos(api, strconv.Itoa(userForPhoto.UserId))
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(len(photos))
	mapJson, err := json.Marshal(Photos{photos})
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func addToQueue()

func getPhotosArchiveHandler(w http.ResponseWriter, r *http.Request) {
	var getPhotoRequest GetPhotosRequest

	err := json.NewDecoder(r.Body).Decode(&getPhotoRequest)
	if err != nil {
		log.Println(err)
		return
	}
	go downloadFiles(getPhotoRequest.Photos)
}

func downloadFiles(files []string) {
	log.Println("Total files: " + strconv.Itoa(len(files)))
	for _, link := range files {
		downloadSingleFile(link)
	}
}

func downloadSingleFile(url string) {
	tokens := strings.Split(url, "/")
	fileName := "userFiles/" + tokens[len(tokens)-1]
	dontDownload := false
	for {
		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			break
		} else {
			dontDownload = true
		}
	}
	if dontDownload {
		return
	}
	output, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
}

func queueWorker() {
	for {
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	log.Println("Application started...")
	Queue.CurrentJobId = -1
	Queue.Items = make([]QueueItem, 0)
	Queue.MapTokenId = make(map[string]int)
	Queue.Free = true
	go queueWorker()
	log.Println("Queue initialized...")
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/token/", tokenHandler)
	http.HandleFunc("/getPhotos/", photosHandler)
	http.HandleFunc("/getArchive/", getPhotosArchiveHandler)
	log.Println("Server started at port :4100")
	err := http.ListenAndServe(":4100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
