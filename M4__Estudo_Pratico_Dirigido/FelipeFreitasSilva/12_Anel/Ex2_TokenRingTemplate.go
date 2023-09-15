/*
	* Felipe Freitas Silva
	* 15/09/2023

	* Conforme o arquivo Ex2-TokenRingExplanacao.pdf (disponível em: https://github.com/EngenhariaSoftwarePUCRS/Fundamentos_de_Processamento_Paralelo_e_Distribuido/blob/develop/M4__Estudo_Pratico_Dirigido/12-Anel/Ex2-TokenRingExplanacao.pdf)
	* 1) Implemente o código do nodo
	* R: Código
*/

package main

import (
	"fmt"
	"strconv"
	"time"
)

const (
	LOG = false
	NODES_AMOUNT = 4
	MESSAGES_AMOUNT = 100
)

type Msg struct {
	sender   int
	receiver int
	message  string
}

type Packet struct {
	token bool
	msg   Msg
}

type User struct {
	id			int
	sendChan	chan Msg
	receiveChan	chan Msg
}

// Estação tem um usuários que manda e recebe mensagens
//
// A qualquer momento, o usuário pode postar uma mensagem a enviar
func (u User) sendMessages(amount int) {
	for i := 0; i <= amount; i++ {
		u.sendChan <- Msg{u.id, u.id % NODES_AMOUNT, "msg" + strconv.Itoa(i)}
		time.Sleep(100 * time.Millisecond)
	}
}

func (u User) receiveMessages() {
	for {
		m := <-u.receiveChan
		Print(fmt.Sprintf("User: %d received from: %d - %s\n", u.id, m.sender, m.message))
	}
}

func (u User) work() {
	go u.sendMessages(MESSAGES_AMOUNT)
	go u.receiveMessages()
}

type Node struct {
	user		User
	hasToken	bool
	ringMy		chan Packet
	ringNext	chan Packet
}

func (n Node) run() {
	Print("node ", fmt.Sprint(n.user.id))
	for {
		if n.hasToken {
			select {
			case msg := <-n.user.sendChan:
				n.ringNext <- Packet{false, msg}
				<-n.ringMy
				n.ringNext <- Packet{true, Msg{}}
				n.hasToken = false

			default:
				n.ringNext <- Packet{true, Msg{}}
				n.hasToken = false
			}
		} else {
			packet := <-n.ringMy

			if packet.token {
				n.hasToken = true
			} else if packet.msg.receiver == n.user.id {
				n.user.receiveChan <- packet.msg
				n.ringNext <- packet
			} else {
				n.ringNext <- packet
			}
		}
	}
}
		
// Um token fica circulando pelo anel. Quando uma estação recebe um token, ela pode passar o token adiante ou enviar uma mensagem criada pelo usuário (se existir)
//
// Se uma mensagem é enviada ela circula até o originador, que então retira a mensagem do anel e passa o token adiante
func main() {
	var chanRing [NODES_AMOUNT]chan Packet
	var chanSend [NODES_AMOUNT]chan Msg
	var chanRec [NODES_AMOUNT]chan Msg

	for i := 0; i < NODES_AMOUNT; i++ {
		chanRing[i] = make(chan Packet)
		chanSend[i] = make(chan Msg)
		chanRec[i] = make(chan Msg)
	}

	for i := 0; i < (NODES_AMOUNT - 1); i++ {
		user := User{i, chanSend[i], chanRec[i]}
		node := Node{user, false, chanRing[i], chanRing[i+1]}
		go node.run()
		go user.work()
	}
	user := User{NODES_AMOUNT - 1, chanSend[NODES_AMOUNT - 1], chanRec[NODES_AMOUNT - 1]}
	node := Node{user, true, chanRing[NODES_AMOUNT - 1], chanRing[0]}
	go node.run()
	go user.work()

	<- make(chan struct{})
}

func Print(msg... string) {
	if LOG {
		fmt.Print(msg)
	}
}
