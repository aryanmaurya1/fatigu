cd src
go fmt *.go 
go build -o ../strainbot

cd ..

./strainbot -s -y \
--body '{"name" : "Aryan"}' \
--method POST \
--base https://api.writups.tech \
--ep /user \
--headers '{"h1" : "v1", "h2" : "v2"}' \
--log-file  logs.sb \
--hits 100 \
--hit-start 250 \
--hit-stop 250 \
--hit-step 10

# ./strainbot -s \
# --hits 100 \
# --method GET \
# --base https://jsonplaceholder.typicode.com \
# --ep /todos/1 \
# --headers '{"h1" : "v1", "h2" : "v2"}' \
# --log-file  logs.sb \
# --hits 100 \
# --hit-start 1 \
# --hit-stop 100 \
# --hit-step 10

# ./strainbot -s \
# --method GET \
# --base http://ec2-3-139-63-233.us-east-2.compute.amazonaws.com:4324 \
# --ep /video/bigbuck.mp4 \
# --headers '{"range" : "bytes=0-"}' \
# --hit-start 1000 \
# --hit-stop 1000 \
# --hit-step 5 \
# --log-file  logs.sb \
