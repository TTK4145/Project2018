package main

import "fmt"
import "encoding/json"

const (
	UP   = "up"
	DOWN = "down"
)

type orderBroadcastMsg struct {
	MasterID       uint8  `json:"master_id"`
	OrderFloor     int    `json:"order_floor"`
	OrderDirection string `json:"order_direction"`
}

func main() {

}
