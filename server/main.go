package main

import (
	"log"
	"net/http"
	"github.com/yanple/vk_api"
	"fmt"
	"encoding/json"
	"strings"
)

type tokenUrl struct {
	Url string
}

type AuthAnswer struct {
	Code int
	Url  string
}

type TokenAnswer struct {
	Code   int
	Status string
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

func tokenHandler(w http.ResponseWriter, r *http.Request) {
	var t tokenUrl
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Println(err)
		return
	}
	splittedUrl := strings.Split(strings.Split(t.Url, "#")[1], "&")
	for _, urlPart := range splittedUrl {
		tempUrlForMap := strings.Split(urlPart, "=")
		log.Println(tempUrlForMap)
	}
	mapJson, err := json.Marshal(TokenAnswer{Code: 200, Status: "Ok", })
	if err != nil {
		log.Printf("Error marshal: %s", err)
		return
	}
	fmt.Fprint(w, string(mapJson))
}

func main() {
	http.HandleFunc("/auth/", authHandler)
	http.HandleFunc("/token/", tokenHandler)

	log.Println("Server started at port :4100")
	err := http.ListenAndServe(":4100", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
var api vk_api.Api

func prepareMartini() *martini.ClassicMartini {
	m := martini.Classic()


	m.Get("/vk/token", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")

		err := api.OAuth(
			"http://vk.com/away.php?to=http://localhost:3000/vk/token", // redirect uri
			"mGfyk5OHtxMOMX9ftgi0",
			"5930862",
			code)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", http.StatusFound)
	})

	m.Get("/", func(w http.ResponseWriter, r *http.Request) string {
		if api.AccessToken == "" {
			return "<a href='/vk/auth'>Авторизоваться</a>"
		}

		// Api have: AccessToken, UserId, ExpiresIn
		log.Println("[LOG] martini.go:48 ->", api.AccessToken)

		// Make query
		params := make(map[string]string)
		params["domain"] = "yanple"
		params["count"] = "1"

		strResp, err := api.Request("wall.get", params)
		if err != nil {
			panic(err)
		}
		return strResp
	})
	return m
}

func main() {
	prepareMartini().Run()
}*/
