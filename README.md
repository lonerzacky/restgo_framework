# RestGo Framework
A Golang framework for restful API. This repository demonstrates the usage of Restful API within a Golang application.

## Installation
##### Without Docker
1. go dep 
```bash
$ go get -u github.com/golang/dep/cmd/dep
$ dep init -v
$ dep ensure -v
```
2. Create .env file
```bash
$ make env-init
```
> It will create .env file
```
APP_HOST=yourapphost
APP_DB_HOST=yourdbhost
APP_PORT=yourapphost
APP_VERSION=ypurappversion
DB_CONNECTION=yourdbdriver
DB_USERNAME=yourdatabaseuser
DB_DATABASE=yourdatabasename
DB_PASSWORD=yourdatabasepassword
DB_PORT=yourdatabaseport
```

## Starting App
<pre>
 go run main.go
</pre>

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

