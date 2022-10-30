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

func DeleteArticle(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	db := mongodb.GetDB()
	defer db.Client().Disconnect(context.Background())
	parsedId, err := strconv.ParseInt(id, 10, 12)

	if err != nil {
		err := structs.Error{
			Status:  400,
			Message: "Erro, insira um número válido.",
		}
		jsonError, _ := json.Marshal(err)
		w.WriteHeader(int(err.Status))
		w.Write(jsonError)
		return
	}

	documentExists := db.Collection("articles").FindOne(context.TODO(), bson.D{{Key: "id", Value: parsedId}})

	if documentExists.Err() != nil {
		err := structs.Error{
			Status:  400,
			Message: "Nenhum registro encontrado com esse id",
		}
		jsonError, _ := json.Marshal(err)
		w.WriteHeader(int(err.Status))
		w.Write(jsonError)
		return
	}

	db.Collection("articles").DeleteOne(context.TODO(), bson.D{{Key: "id", Value: parsedId}})

	sucessMessage := structs.Error{
		Status:  200,
		Message: "Registro excluído com sucesso.",
	}
	jsonSucessMessage, _ := json.Marshal(sucessMessage)

	w.WriteHeader(int(sucessMessage.Status))
	w.Write(jsonSucessMessage)
}
