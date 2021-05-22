cd src
go fmt *.go 
go build -o ../fatigu
# ./fatigu -s \
# --body '{"name" : "Aryan"}' \
# --base https://api.writups.tech \
# --hits 1000 --method POST --ep /user 
cd ..
./fatigu -s \
--hits 10 --method GET --ep /todos/1 \
--base https://jsonplaceholder.typicode.com