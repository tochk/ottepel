# SVK

Tool for download photo from vk converstaions

Installing:

1. Download app
```sh
git clone https://github.com/tochk/ottepel.git
```


2. Install and run backend (served on localhost:4100)
Required go 1.8 or higher 
```sh
$ cd server
$ go get github.com/yanple/vk_api
$ go run main.go
```

3. Run frontend (served on localhost:4200)
Required node.js 7.4.0 with npm 4.0.5 or higher
```sh
$ cd client
$ npm install -g @angular/cli
$ npm install
$ npm run serve
```
