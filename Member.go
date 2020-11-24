package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Member struct{
	nowN int
	nowContent string
	maxN int
	nowSeq int
}

type proposeMessage struct{
	seq int
	content string
}

type acceptorMessage struct {
	seq int
	content string
	result bool
}

type acceptMessage struct {
	n int
	value int
	content string
}


func newMember() Member{
	mem := Member{0,"",0,0}
	return mem
}

func(mem *Member) setSeq(nowSeq int){
	mem.nowSeq = nowSeq
}
func(mem *Member) setMember(nowN int, nowContent string, maxN int, nowSeq int){
	mem.nowN = nowN
	mem.nowContent = nowContent
	mem.maxN = maxN
	mem.nowSeq = nowSeq
}
func(mem *Member) setProposer(nowN int, nowContent string, maxN int){
	mem.nowN = nowN
	mem.nowContent = nowContent
	mem.maxN = maxN
}
func(mem *Member) getSeq() int{
	return mem.nowSeq
}
func(mem *Member) addNowSeq() bool{
	mem.nowSeq += 5
	fmt.Printf("调用addNowSeq方法，现在nowSeq=%d\n", mem.nowSeq)
	return true
}

func(mem *Member) makePropose(wantContent string) proposeMessage{
	message := proposeMessage{mem.nowSeq, wantContent}
	return message
}

func(mem *Member) checkPropose(message proposeMessage) acceptorMessage {
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	accMessage := acceptorMessage{0,"",true}
	if mem.maxN >= message.seq {
		accMessage.result = false
	}else{
		accMessage.result = true
		accMessage.seq = mem.nowN
		accMessage.content = mem.nowContent
		mem.maxN = message.seq
	}
	//fmt.Printf("checkPropose = ", accMessage)
	//fmt.Printf("\n")
	return accMessage
}

func(mem *Member) Accept(message acceptMessage) bool{
	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	if mem.maxN > message.n {
		return false
	}else{
		mem.nowN = message.n
		mem.nowContent = message.content
		//fmt.Printf("现在的content = %s\n", message.content)
		return true
	}
}

