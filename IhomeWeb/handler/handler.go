package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	example "github.com/micro/examples/template/srv/proto/example"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
)

func ExampleCall(w http.ResponseWriter, r *http.Request,params httprouter.Params) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	cli:=grpc.NewService()
	cli.Init()
	// call the backend service
	exampleClient := example.NewExampleService("go.micro.srv.template", cli.Client())
	rsp, err := exampleClient.Call(context.TODO(), &example.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
