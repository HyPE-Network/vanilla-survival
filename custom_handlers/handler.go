package custom_handlers

import (
	"github.com/HyPE-Network/vanilla-proxy/proxy/block"
	"github.com/HyPE-Network/vanilla-proxy/proxy/player/human"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type PlaceBlockHandler struct {
}

func (PlaceBlockHandler) Handle(pk packet.Packet, player human.Human) (bool, packet.Packet, error) {
	dataPacket := pk.(*packet.InventoryTransaction)

	switch td := dataPacket.TransactionData.(type) {
	case *protocol.UseItemTransactionData:
		if td.ActionType == protocol.UseItemActionClickBlock {
			if td.HeldItem.Stack.BlockRuntimeID != 0 {
				state := block.Blocks[td.HeldItem.Stack.BlockRuntimeID]

				if state.Name == "minecraft:obsidian" && player.InNether() {
					player.SendMessage("§cYou can't place obsidian in this world!")
					return false, pk, nil
				}
			}
		}
	}

	return true, pk, nil
}
