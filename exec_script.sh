cd src
go fmt *.go 
go build -o ../strainbot

cd ..

./strainbot -s -l \
--hits 100 \
--method GET \
--base https://jsonplaceholder.typicode.com \
--ep /todos/1 \
--headers '{"h1" : "v1", "h2" : "v2"}' \
--log-file  logs.sb \
--hit-start 10 \
--hit-stop 100 \
--hit-step 10
