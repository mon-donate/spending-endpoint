# Spending Endpoint
Access the endpoint by sending a request to the endpoint. Eg: `http://urlhere.com/?amount=30&name=Oxfam`

(Replace 30 and Oxfam with the actual amount and charity name)

## How to run
Install go (https://golang.org/doc/install)

Then run the app:
```
go run mondo.go
```

## Serving via Ngrok
Install ngrok (https://ngrok.com)

Run ngrok, whilst the Go app is running:
```
ngrok http 8080
```

Then access the app via the url given (and change urlhere.com to the one given in the terminal)

---
mark@larah.me
