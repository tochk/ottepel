package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/yanple/vk_api"
)

const VK_APP_ID = "5930862"

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

type PhotoArchiveRest struct {
	Token string
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

type IsFileExistStruct struct {
	IsExist bool
}

type IsFileExistRequest struct {
	Token string
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	authUrl := "https://oauth.vk.com/authorize?client_id=" + VK_APP_ID + "&redirect_uri=https%3A%2F%2Fapi.vk.com%2Fblank.html&response_type=token&scope=friends%2Cmessages%2Coffline"
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
	log.Println("Sorting photo")
	sortedResult := make([]string, 0, len(result))
	for _, link := range result {
		isNotExist := true
		for _, sortedLink := range sortedResult {
			if sortedLink == link {
				isNotExist = false
				break
			}
		}
		if isNotExist {
			sortedResult = append(sortedResult, link)
		}
	}
	log.Println("All photos loaded from attachments")
	return sortedResult, nil
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
	log.Println("Total photos going to frontend", len(photos))
	mapJson, err := json.Marshal(Photos{photos})
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func createArchive(files []string, dir string) {
	log.Println("Creating archive for", dir)
	downloadFiles(files, dir)
	log.Println("Archive created for", dir)
	log.Println("Moving archive to public directory")
	cmdMove := exec.Command("mv", "tempUserFiles/"+dir+".zip", "userFiles/"+dir+".zip")
	err := cmdMove.Run()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Moving sucesfful")
}

func getPhotosArchiveHandler(w http.ResponseWriter, r *http.Request) {
	var getPhotoRequest GetPhotosRequest

	err := json.NewDecoder(r.Body).Decode(&getPhotoRequest)
	if err != nil {
		log.Println(err)
		return
	}
	token := generateToken()
	go createArchive(getPhotoRequest.Photos, token)
	mapJson, err := json.Marshal(PhotoArchiveRest{Token: token})
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func generateToken() string {
	b := make([]byte, 30)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func isFileExistHandler(w http.ResponseWriter, r *http.Request) {
	var isFileExistRequest IsFileExistRequest

	err := json.NewDecoder(r.Body).Decode(&isFileExistRequest)
	if err != nil {
		log.Println(err)
		return
	}
	isExist := true
	if _, err := os.Stat("userFiles/" + isFileExistRequest.Token + ".zip"); os.IsNotExist(err) {
		isExist = false
	}
	mapJson, err := json.Marshal(IsFileExistStruct{IsExist: isExist})
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func createBashFile(dir string) {
	cmdStr := "#/usr/bin/bash\n" + "zip" + " -m " + "tempUserFiles/" + dir + ".zip " + dir + "/*\n" + "echo \"heh\""
	file, err := os.Create(dir + ".sh")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return
	}
	w := bufio.NewWriter(file)
	buf := []byte(cmdStr)
	_, err = w.Write(buf)
	if err != nil {
		log.Println(err)
		return
	}
	err = w.Flush()
	if err != nil {
		log.Println(err)
		return
	}
}

func downloadFiles(files []string, dir string) {
	log.Println("Total files: " + strconv.Itoa(len(files)))
	err := os.Mkdir(dir, os.FileMode(0744))
	if err != nil {
		log.Println("Error while creating directory", dir, "-", err)
		return
	}
	cmd := exec.Command("zip", "tempUserFiles/"+dir+".zip", dir)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
		return
	}
	for _, link := range files {
		downloadSingleFile(dir, link)
		createBashFile(dir)
		cmd := exec.Command("bash", dir+".sh")
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("Files downloaded for", dir)
}

func downloadSingleFile(dir string, url string) {
	tokens := strings.Split(url, "/")
	fileName := dir + "/" + tokens[len(tokens)-1]
	output, err := os.Create(fileName)
	if err != nil {
		log.Println("Error while creating", fileName, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		log.Println("Error while downloading", url, "-", err)
		return
	}
}

func userFilesHandler(w http.ResponseWriter, r *http.Request) {
	path := "." + r.URL.Path
	if f, err := os.Stat(path); err == nil && !f.IsDir() {
		http.ServeFile(w, r, path)
		return
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/token/", tokenHandler)
	http.HandleFunc("/getPhotos/", photosHandler)
	http.HandleFunc("/getArchive/", getPhotosArchiveHandler)
	http.HandleFunc("/userFiles/", userFilesHandler)
	http.HandleFunc("/isFileExist/", isFileExistHandler)
	log.Println("Server started at port :4100")
	err := http.ListenAndServe(":4100", nil)
	if err != nil {
		log.Fatal(err)
	}
}
