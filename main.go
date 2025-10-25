package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	_ "reflect"
	"sync"

	"aler93.com/adressongo/model"
	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var result model.Cep

type errMsg struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func main() {
	loadConf()

	//var search string
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:" + App.RedisPort,
		Password: App.RedisPassword,
		DB:       0,
		Protocol: 2,
	})

	if client == nil {
		panic("Conexão com Redis falhou")
	}

	fmt.Println("Iniciando servidor na porta " + App.ServerPort)

	http.HandleFunc("/", handler)

	err := http.ListenAndServe(":"+App.ServerPort, nil)
	if err != nil {
		panic(err)
	}
}

func doReq(cep string, w http.ResponseWriter, r *http.Request) {
	var group sync.WaitGroup
	ctx := context.Background()

	//fmt.Println("Iniciando requisições para:", cep)
	printed := false
	result = model.Cep{}
	group.Add(3)
	go (func() {
		viaCep(cep)
		if printed == false {
			printed = true
			print, _ := json.Marshal(result)

			if len(result.Cep) > 0 {
				client.Set(ctx, cep, print, 0)
				renderJson(JsonResp{w, r, result, 200})
			}

			//fmt.Println("ViaCep - Retornando resposta")
		}

		//fmt.Println("ViaCep concluído")
		group.Done()
	})()
	go (func() {
		openCep(cep)
		if printed == false {
			printed = true
			print, _ := json.Marshal(result)

			if len(result.Cep) > 0 {
				client.Set(ctx, cep, print, 0)
				renderJson(JsonResp{w, r, result, 200})
			}
			//fmt.Println("OpenCep - Retornando resposta")
		}

		//fmt.Println("OpenCep concluído")
		group.Done()
	})()
	go (func() {
		brasilApi(cep)
		if printed == false {
			printed = true
			print, _ := json.Marshal(result)

			if len(result.Cep) > 0 {
				client.Set(ctx, cep, print, 0)
				renderJson(JsonResp{w, r, result, 200})
			}
			//fmt.Println("BrasilApi - Retornando resposta")
		}

		//fmt.Println("BrasilApi concluído")
		group.Done()

		if len(result.Cep) <= 0 {
			renderJson(JsonResp{w, r, errMsg{true, "CEP não encontrado"}, 404})
		}
	})()

	//fmt.Println("Aguardando conclusões das threads")
	group.Wait()
	//fmt.Println("Encerrando doReq()")
}

func get(key string) bool {
	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		//fmt.Println("CEP: ", key, "não encontrado localmente")
		return false
	}

	json.Unmarshal([]byte(val), &result)
	//fmt.Println("CEP: ", key, "Local")
	result.Local = true

	return true
}
