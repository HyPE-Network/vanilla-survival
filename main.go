package main

import (
	"vanilla-survival/custom_handlers"

	"github.com/HyPE-Network/vanilla-proxy/handler"
	"github.com/HyPE-Network/vanilla-proxy/handler/handlers"
	"github.com/HyPE-Network/vanilla-proxy/log"
	"github.com/HyPE-Network/vanilla-proxy/proxy"
	"github.com/HyPE-Network/vanilla-proxy/proxy/player/manager"
	"github.com/HyPE-Network/vanilla-proxy/utils"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

func main() {
	log.Logger = log.New()
	log.Logger.Debugln("Logger has been started")

	config := utils.ReadConfig()
	proxy.ProxyInstance = proxy.New(config, manager.NewPlayerManager())

	err := proxy.ProxyInstance.Start(loadHandlers())
	if err != nil {
		log.Logger.Errorln("Error while starting server: ", err)
		panic(err)
	}
}

func loadHandlers() handler.HandlerManager {
	h := handlers.New()
	h.RegisterHandler(packet.IDInventoryTransaction, custom_handlers.PlaceBlockHandler{})

	return h
}
