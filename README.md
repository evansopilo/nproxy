# nproxy
Simple reverse proxy that forwards and receives request to and from target server.
Hides end server IP address.

## Setup
To setup this project:

1. Install [GO](https://golang.org/doc/install)
2. Clone repo `git clone https://github.com/evansopilo/nproxy.git`

## Usage
```
// incomming request will be forwaded to "https://www.example.com" and response send back to the 
// client.
http.ListenAndServe(":8080", ReverseProxy("https://www.example.com"))
```