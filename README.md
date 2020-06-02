# video-analysis-worker


##
* set GOOGLE_APPLICATION_CREDENTIALS=magisterij-6d3594ec69ea.json
* export GOOGLE_APPLICATION_CREDENTIALS="magisterij-6d3594ec69ea.json"


## Worker example message

```
{
"mediaId": 36
}
```

##PROTOCOL BUFFER

```.env
protoc proto\helloworld.proto --go_out=plugins=grpc:.
```