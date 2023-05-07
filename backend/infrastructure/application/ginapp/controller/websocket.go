package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/fcfcqloow/first-todo-list/backend/infrastructure/application/ginapp/model"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WebsocketHandle(handlerMap map[string]gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.RemoteAddr)
		conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			ctx.Errors = append(ctx.Errors, ctx.Error(err))
			return
		}

		defer conn.Close()

		var request model.WebsocketJson
		for {
			_, byts, err := conn.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				return
			}
			if err != nil {
				ctx.Errors = append(ctx.Errors, ctx.Error(err))
				return
			}

			if err := json.Unmarshal(byts, &request); err != nil {
				ctx.Errors = append(ctx.Errors, ctx.Error(err))
				return
			}

			ctx.Request.Body = ioutil.NopCloser(bytes.NewReader(byts))

			handler, ok := handlerMap[request.Type]
			if !ok {
				ctx.Errors = append(ctx.Errors, ctx.Error(errors.New("not implement handler "+request.Type)))
				return
			}

			handler(ctx)
			if len(ctx.Errors) > 0 {
				return
			}
		}
	}
}
