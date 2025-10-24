package main

import (
	"aler93.com/adressongo/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	_ "reflect"
	"regexp"
	"sync"
)

var client *redis.Client
var result model.Cep

// var service int // 0=viacep, 1=opencep, 2=brasilapi
var printed bool = false

func main() {
	//fmt.Println("starting...")
	var wg sync.WaitGroup

	loadConf()

	var search string
	wg.Add(2)
	go (func() {
		client = redis.NewClient(&redis.Options{
			Addr:     "127.0.0.1:6389",
			Password: "secretPass",
			DB:       0,
			Protocol: 2,
		})

		if client == nil {
			panic("Conexão com Redis falhou")
		}

		/*ctx := context.Background()
		curPoint, err := client.Get(ctx, "service").Result()
		if err != nil {
			//fmt.Println("CEP: ", key, "não encontrado localmente")
			service = 0
			client.Set(ctx, "service", 1, 0)
		} else {
			service, _ = strconv.Atoi(curPoint)
		}*/

		wg.Done()
	})()

	go (func() {
		reg := regexp.MustCompile(`^[0-9]+$`)
		search = os.Args[1]

		if App.Length != len(search) {
			panic("Informe o cep")
		}

		if !reg.MatchString(search) {
			panic("CEP inválido")
		}

		wg.Done()
	})()

	wg.Wait()

	get(search)
	if len(result.Cep) > 0 {
		printed = true
		print, _ := json.Marshal(result)
		fmt.Println(string(print))

		os.Exit(0)
	}

	doReq(search)

	// Exibe resultado do request
	if printed == false {
		print, _ := json.Marshal(result)
		fmt.Println(string(print))

		ctx := context.Background()
		client.Set(ctx, search, print, 0)
	}

	os.Exit(0)
}

func doReq(cep string) {
	//service = 0
	//fmt.Println(service)

	var group sync.WaitGroup
	ctx := context.Background()

	group.Add(3)
	go (func() {
		viaCep(cep)
		if printed == false {
			print, _ := json.Marshal(result)
			fmt.Println(string(print))
			printed = true

			client.Set(ctx, cep, print, 0)
		}

		os.Exit(0)
		group.Done()
	})()
	go (func() {
		openCep(cep)
		if printed == false {
			print, _ := json.Marshal(result)
			fmt.Println(string(print))
			printed = true

			client.Set(ctx, cep, print, 0)
		}

		os.Exit(0)
		group.Done()
	})()
	go (func() {
		brasilApi(cep)
		if printed == false {
			print, _ := json.Marshal(result)
			fmt.Println(string(print))
			printed = true

			client.Set(ctx, cep, print, 0)
		}

		os.Exit(0)
		group.Done()
	})()

	group.Wait()
}

func get(key string) bool {
	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		//fmt.Println("CEP: ", key, "não encontrado localmente")
		return false
	}

	json.Unmarshal([]byte(val), &result)

	result.Local = true

	return true
}
