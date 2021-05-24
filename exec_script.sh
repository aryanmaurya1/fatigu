cd src
go fmt *.go 
go build -o ../strainbot

cd ..

# ./strainbot -s \
# --body '{"name" : "Aryan"}' \
# --method POST \
# --base https://api.writups.tech \
# --ep /user \
# --headers '{"h1" : "v1", "h2" : "v2"}' \
# --log-file  logs.sb \
# --hits 100 \
# --hit-start 10 \
# --hit-stop 500 \
# --hit-step 10

./strainbot -s \
--hits 100 \
--method GET \
--base https://jsonplaceholder.typicode.com \
--ep /todos/1 \
--headers '{"h1" : "v1", "h2" : "v2"}' \
--log-file  logs.sb \
--hits 100 \
--hit-start 10 \
--hit-stop 500 \
--hit-step 10