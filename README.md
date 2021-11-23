# KeyValue API Assignment

## Installation

App runs on localhost by ``docker-compose up -d`` command. If you want to view logs or develop, run ``docker-compose up --build`` to make fresh build of the app.

Code documentation could seen by **/lib/godoc/index.html** url while running app. You can install godoc locally by ``go install golang.org/x/tools/cmd/godoc@latest`` and run ``godoc -url "http://localhost:6060/pkg/deliveryhero/" > ./lib/godoc/index.html`` to regenerate docs.

The api is served by heroku with dockerized web type.

Temp url is https://peaceful-stream-88325.herokuapp.com/

## API Endpoints

-   [GET] /get/{key} => keyValueModel
-   [POST] /set (keyValueModel) => true
-   [POST] /flushDb => true

Model **keyValueModel**
{
    "key": string,
    "value": string
}