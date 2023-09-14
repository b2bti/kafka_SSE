package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
)

var producer sarama.SyncProducer

func main() {
	brokerList := []string{"localhost:9091", "localhost:9092", "localhost:9093"}
	topic := "my-topic-three"
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	var err error
	producer, err = sarama.NewSyncProducer(brokerList, config)
	if err != nil {
		fmt.Printf("Erro ao criar produtor Kafka: %v\n", err)
		return
	}
	defer producer.Close()

	priceCh := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go generateCryptoPrice(ctx, priceCh)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	for {
		select {
		case <-sigchan:
			fmt.Println("Recebido sinal de interrupção. Encerrando o produtor Kafka.")
			return
		case price := <-priceCh:
			sendPriceToKafka(topic, price)
		}
	}
}

func generateCryptoPrice(ctx context.Context, priceCh chan<- int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	ticker := time.NewTicker(time.Second)

outerloop:
	for {
		select {
		case <-ctx.Done():
			break outerloop
		case <-ticker.C:
			p := r.Intn(100)
			priceCh <- p
		}
	}

	ticker.Stop()
	close(priceCh)
	fmt.Println("generateCryptoPrice: Finished generating")
}

func sendPriceToKafka(topic string, price int) {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(fmt.Sprintf("%d", price)),
	}

	_, _, err := producer.SendMessage(kafkaMessage)
	if err != nil {
		fmt.Printf("Erro ao enviar mensagem para o Kafka: %v\n", err)
	}
}
