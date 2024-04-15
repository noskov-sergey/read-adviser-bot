package main

import (
	"flag"
	"log"
	tgClient "read-aviser-bot/clients/telegram"
	event_consumer "read-aviser-bot/consumer/event-consumer"
	"read-aviser-bot/events/telegram"
	"read-aviser-bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage/memory"
	bathSize    = 100
)

func main() {
	var eventsProcessor = telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Printf("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, bathSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}

func mustToken() string {
	t := flag.String("tg-token", "", "token for access telegram")
	flag.Parse()

	if *t == "" {
		log.Fatal("can't find access token, shutdown")
	}

	return *t
}
