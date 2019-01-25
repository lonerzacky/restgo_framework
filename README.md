# RestGo Framework
A Golang framework for restful API. This repository demonstrates the usage of Restful API within a Golang application.

## Installation
```bash
$ go get -u github.com/gin-gonic/gin
$ go get -u github.com/jinzhu/gorm
$ go get -u github.com/joho/godotenv
$ go get -u github.com/dgrijalva/jwt-go
```
## Configuration Database
Edit _.env_ file and setup an application config
<pre>
 APP_HOST=localhost
 APP_PORT=5000
 APP_VERSION=1.0.0
 DB_CONNECTION=mysql
 DB_USERNAME=root
 DB_DATABASE=restgo_framework
 DB_PASSWORD=
</pre>
  
## Starting App
<pre>
 go run main.go
</pre>
This will start the application  Just open **http://localhost:5000**.

## Request Headers
<table>
    <thead>
        <tr>
            <th><strong>Required</strong></th>
            <th><strong>Key</strong></th>
            <th><strong>Value</strong></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Yes</td>
            <td>Content-Type</td>
            <td>application/x-www-form-urlencoded</td>
        </tr>
    </tbody>
</table>

## Features
- env configuration
- Controller based
- Route based
- Model based
- Helper JSON format

## Database Sample
![alt tag](https://github.com/lonerzacky/restgo_framework/blob/master/diagram.png)

