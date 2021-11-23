# KeyValue API Assignment

## Installation
App runs on localhost by ``docker-compose up -d`` command. If you want to view logs or develop, run ``docker-compose up --build`` to make fresh build of the app.

Code documentation can be seen from **/lib/godoc/index.html** link while running app. You can install godoc locally by ``go install golang.org/x/tools/cmd/godoc@latest`` and run ``godoc -url "http://localhost:6060/pkg/container/deliveryhero/" > /lib/godoc/index.html`` to regenerate docs.
