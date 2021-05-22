go fmt *.go 
go run *.go -s \
--body '{"name" : "Aryan"}' \
--base https://api.writups.tech \
--hits 10 --method GET,POST --ep /user \
--body-file ./config.json \
--config-file ./config.json