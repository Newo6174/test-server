# test-server 
A simple server for the BGL pre-interview task, bulit with Go

## Requirements
- Go > 1.7 
- [go-chi](https://github.com/go-chi/chi)


## How to run this
**1. Clone This Repo**

    go get https://github.com/Newo6174/test-server
    cd test-server
    
**2. Run the server on port 8080**

    go run server.go
    
**3. Check if the server is running**

    curl 127.0.0.1:8080

**4. Query the endpoints**

#### ADD

    curl -H "Content-Type: application/json" -d "{\"key\": \"Hello\", \"value\": \"World\"}" http://127.0.0.1:8080/add

#### LIST

    curl http://127.0.0.1:8080/list

## Special Thanks
>I just want to say thank you to the staffs in BGL who developed this task. 
>I am a beginner in API and I've learnt a lot from this task.
>I am very appreciated for this great opportunity :)
