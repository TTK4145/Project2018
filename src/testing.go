package main

import (
	"encoding/json"
	"fmt"
	"net"
	"network"
	"time"
)

const server_ip = "129.241.187.38"

type Direction int

const (
	Up Direction = iota
	Down
)

type ElevatorState int

const (
	MovingUp ElevatorState = iota
	MovingDown
	StopUp
	StopDown
)

type Order struct {
	OrderID   int       `json:"order_id"`
	Floor     int       `json:"floor"`
	Direction Direction `json:"direction"`
}

type OrderPlacedMsg struct {
	SourceID int    `json:"source_id"`
	MsgType  string `json:"msg_type"`
	Order    Order  `json:"order"`
	Priority int    `json:"priority"`
}

type TakeOrderMsg struct {
	Order Order `json:"order"`
	CmdID int   `json:"cmd_id"` // specify the elevator that should take the order
}

func main() {
	err := network.SendBytes([]byte("Message sending test"), server_ip+":20010")
	if err != nil {
		fmt.Println(err)
	}

	addr, err := net.ResolveUDPAddr("udp", ":20010")
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	buffer := make([]byte, 1024)
	conn.SetReadDeadline(time.Now().Add(1 * time.Second))
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(buffer[:n]))

	msg := OrderPlacedMsg{SourceID: "source", OrderID: 1234, MsgType: "testing type", Order: Order{Floor: 1, Direction: -1}, Priority: 1}

	data, err := json.MarshalIndent(msg, "", " ")
	if err != nil {
		fmt.Printf("JSON failed")
	}
	fmt.Printf("%s\n", data)

	var un_msg OrderPlacedMsg
	json.Unmarshal(data, &un_msg)

	fmt.Println(un_msg)
}
