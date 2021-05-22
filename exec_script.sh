cd src
go fmt *.go 
go build -o ../fatigu
# ./fatigu -s \
# --body '{"name" : "Aryan"}' \
# --base https://api.writups.tech \
# --hits 1000 --method POST --ep /user 
cd ..
./fatigu -s \
--hits 100 \
--method GET \
--base https://jsonplaceholder.typicode.com \
--ep /todos/1 \
--headers '{"h1" : "v1", "h2" : "v2"}'