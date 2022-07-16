package main

import (
	"github.com/RetardedRedhaired/GoRshokBot/internal"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var token string

func main(){
	token = ScanToken()

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalln("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(internal.Ready)

	err = dg.Open()
	if err != nil {
		log.Fatalln("Error opening Discord session: ", err)
	}

	log.Println("Gorshok is now alive. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
