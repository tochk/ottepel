# SVK

Tool for download photo from vk converstaions

Installing:

1. Download app
```sh
git clone https://github.com/tochk/ottepel.git
```


2. Build backend (served on localhost:4100)
Required go 1.8 or higher 
```sh
$ cd server
$ go get github.com/yanple/vk_api
```

3. Run backend:
```sh
$ go run main.go
```

4. Setup frontend (served on localhost:4200)
Required node.js 7.4.0 with npm 4.0.5 or higher
```sh
$ cd client
$ npm install -g @angular/cli
$ npm install
```

5. Change 
``` 
"target": "http://deploy-12.sgu.ru:4100",
``` 
line in `client/proxy.conf.json` on 
``` 
"target": "http://localhost:4100",
```

6. Change 
``` 
  SERVER: "http://deploy-12.sgu.ru:4100",
``` 
line in `client/src/app/constants/urls.ts` on 
``` 
  SERVER: "http://localhost:4100",
```

7. Run frontend:
```sh
$ npm run serve
```

8. Open http://localhost:4200 in your browser
