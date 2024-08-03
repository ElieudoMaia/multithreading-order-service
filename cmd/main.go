package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/elieudomaia/multithreading-order-service/internal/order/infra/database"
	"github.com/elieudomaia/multithreading-order-service/internal/order/usecase"
	"github.com/elieudomaia/multithreading-order-service/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)

	http.HandleFunc("/total", func(w http.ResponseWriter, r *http.Request) {
		uc := usecase.NewGetTotalUseCase(repository)
		output, err := uc.Execute()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(output)
	})

	go http.ListenAndServe(":8181", nil)

	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	ch.Qos(100, 0, false)

	deliveryMessage := make(chan amqp.Delivery)

	wg := sync.WaitGroup{}

	go rabbitmq.Consume(ch, deliveryMessage)

	maxWorkers := 10

	wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		defer wg.Done()
		go worker(deliveryMessage, uc, i)
	}

	wg.Wait()

	// output, err := uc.Execute(usecase.OrderInputDTO{
	// 	ID:    "1",
	// 	Price: 100,
	// 	Tax:   10,
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// println(output.FinalPrice)
}

func worker(deliveredMessage chan amqp.Delivery, uc *usecase.CalculateFinalPriceUseCase, workerId int) {
	for msg := range deliveredMessage {
		inputDTO := usecase.OrderInputDTO{}
		err := json.Unmarshal(msg.Body, &inputDTO)
		if err != nil {
			fmt.Println(err)
		}

		println("workerId", workerId, "Message", inputDTO.ID, inputDTO.Price)

		inputDTO.Tax = 10.0

		_, err = uc.Execute(inputDTO)
		if err != nil {
			fmt.Println(err)
		}

		msg.Ack(false)
	}
}
