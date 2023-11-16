package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

type Message struct {
	Text string `json:"text"`
}

func main() {
	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/":
			data := map[string]string{"message": "Hello, fast-Http World!"}
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(fasthttp.StatusOK)
			if err := json.NewEncoder(ctx).Encode(data); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			}
		case "/post":
			var message Message
			if err := json.Unmarshal(ctx.PostBody(), &message); err != nil {
				log.Printf("Invalid request payload: %v", err)
				ctx.Error("Invalid request payload", fasthttp.StatusBadRequest)
				return
			}

			// Process the received data (in this case, just echoing it back)
			response := map[string]string{"received": message.Text}
			ctx.Response.Header.Set("Content-Type", "application/json")
			ctx.SetStatusCode(fasthttp.StatusOK)
			if err := json.NewEncoder(ctx).Encode(response); err != nil {
				log.Printf("Error encoding JSON: %v", err)
				ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			}
		default:
			ctx.NotFound()
		}
	}

	log.Println("Server is running on :1335...")
	if err := fasthttp.ListenAndServe(":1334", requestHandler); err != nil {
		log.Fatal(err)
	}
}
