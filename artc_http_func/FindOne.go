package artc_http_func

import (
	"articles/mongodb"
	"articles/structs"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func FindOne(w http.ResponseWriter, r *http.Request) {
	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())

	id := mux.Vars(r)["id"]

	parsedId, err := strconv.ParseInt(id, 10, 16)

	if err != nil {
		err := structs.Error{
			Status:  int16(400),
			Message: err.Error(),
		}

		errorStruct, _ := json.Marshal(err)
		w.WriteHeader(int(err.Status))
		w.Write(errorStruct)
		return
	}

	var result structs.Article
	documentExists := db.Collection("articles").FindOne(r.Context(), bson.D{{Key: "id", Value: parsedId}})

	if documentExists.Err() != nil {
		err := structs.Error{
			Status:  401,
			Message: "Nenhum documento encontrado com o Id informado",
		}

		errorJson, _ := json.Marshal(err)
		w.WriteHeader(int(err.Status))
		w.Write(errorJson)
		return
	}

	documentExists.Decode(&result)

	responseResult, err := json.Marshal(result)

	w.Write(responseResult)
}
