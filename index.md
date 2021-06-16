# Strainbot

<img src="./Strainbot_logo.png">


## About
Strainbot is an API Performace Testing and Analysis Tool. Using Strainbot you can perform a number of concurrent request to an endpoint and it will show you detailed analysis of time taken to complete all the requests.
Basic functionalities are ready to use but it is a work under progress, so some functionalities are not yet implemented.

**Batch mode is still under development**

Feel free to fork this repository, create pull requests, feature request and report bugs.
## Local Compilation
- Golang compiler must be installed
- No extra dependencies other than standard library is required.
- Run **$**`./compile.sh`

## Usage 
To see list of all arguments and parameters required.

&nbsp;&nbsp;&nbsp;&nbsp;**$** `strainbot -h`  

&nbsp;&nbsp;&nbsp;&nbsp;**$** `strainbot --help`

### Flags and Arguments

&nbsp;&nbsp;&nbsp;&nbsp; `-l` If set, shows full logs.

&nbsp;&nbsp;&nbsp;&nbsp; `-log-file` File to write logs.

&nbsp;&nbsp;&nbsp;&nbsp; `-base` Base URL of the API.

&nbsp;&nbsp;&nbsp;&nbsp; `-body` Body to send in every request.

&nbsp;&nbsp;&nbsp;&nbsp; `-body-file`  Path to json file to use as body content.

&nbsp;&nbsp;&nbsp;&nbsp; `-ep` Endpoint of API to hit. `[URL = BASE + ENDPOINT]`

&nbsp;&nbsp;&nbsp;&nbsp; `-headers`  Headers for request in form of key-value pair. `(Valid JSON) (default "{}")`

&nbsp;&nbsp;&nbsp;&nbsp; `-hit-start` Starting value of hit range. `([START, STOP, STEP]) (default -1)`

&nbsp;&nbsp;&nbsp;&nbsp; `-hit-step` Step size to use for excuting range. `([START, STOP, STEP]) (default 10)`

&nbsp;&nbsp;&nbsp;&nbsp; `-hit-stop` Stoping value of hit range. `([START, STOP, STEP]) (default -1)`

&nbsp;&nbsp;&nbsp;&nbsp; `hits`  Number of concurrent hits to perform. (default 10)

&nbsp;&nbsp;&nbsp;&nbsp; `-method` Comma separated list of methods to use. *Do not include space in list*

### Example 

**$** `./strainbot 
-s -l 
--method GET 
--base https://jsonplaceholder.typicode.com 
--ep /todos/1 
--headers '{"h1" : "v1", "h2" : "v2"}' 
--log-file  logs.sb 
--hit-start 10 
--hit-stop 100
--hit-step 10`
