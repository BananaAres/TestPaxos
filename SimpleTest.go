package main

import (
	"fmt"
	"math/rand"
	"time"
)

func begin(){
	fmt.Printf("本程序模拟Paxos共识机制，可以选择节点数量，设置模拟情景等\n")
	fmt.Printf("请输入需要的节点数量:\n")
}

//func setMember(memberMap map[int] Member) bool{
//	for {
//		fmt.Printf("请输入需要修改的节点号\n")
//
//	}
//}
//第一阶段

func setRandomMembers(members []Member){
	for i:=0; i<len(members); i++{
		time.Sleep(time.Duration(200) * time.Millisecond)
		rand.Seed(time.Now().Unix())
		number := rand.Intn(30)
		members[i].setMember(number,"",number,0)
	}
}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}


func propose(members []Member) {
	var n int
	var content string
	var maxAcceptorN int
	var agreeSum int
	var flag int
	//var boolean bool
	var tmp string
	var contentOfMaxN string
	_ = contentOfMaxN
	flag = 1
	for1:
	for {
		maxAcceptorN = 0
		agreeSum = 0
		//fmt.Printf("map大小等于=%d\n",len(*membersMap))
		fmt.Printf("请输入想要发起共识请求的节点编号和共识内容\n")
		fmt.Scanf("%d %s", &n, &content)
		if n > len(members) {
			fmt.Printf("请输入正确的节点编号和共识内容\n")
		}else{
			//fmt.Printf("map的地址%p\n",membersMap)
			for2:
			for {
			//	boolean = true
				time.Sleep(time.Duration(1) * time.Second)
				fmt.Printf("第一阶段的第%d轮共识请求尝试\n", flag)
				proMessage := members[n].makePropose(content)
				rand.Seed(time.Now().Unix())
				//生成大于节点总数1/2的随机数，用来确定发给多少个节点
				messageSum := rand.Intn(len(members)-(len(members)/2+1))+(len(members)/2+1)
				//生成不重复的随机数数组，用来确定发给哪个节点
				arr := generateRandomNumber(0,len(members)-1,messageSum)
				fmt.Printf("第%d次请求发给了%d个节点，他们是",flag,messageSum,arr)
				fmt.Printf("\n")
			//for3:
				for i := 0; i < messageSum; i++ {
					accMessage := members[arr[i]].checkPropose(proMessage)
					//if accMessage.result == false {
					//	members[n].addNowSeq()
					//	fmt.Printf("有节点拒绝请求，增加序号值，重新发送共识请求\n")
					//	maxAcceptorN = 0
					//	agreeSum = 0
					//	boolean = false
					//	break for3
					//}
						if accMessage.result == true {
						agreeSum++
						//fmt.Printf("agreesum = %d\n", agreeSum)
						if accMessage.seq > maxAcceptorN {
							contentOfMaxN = accMessage.content
							maxAcceptorN = accMessage.seq
						}
					}
				}
			//	if boolean == true{
					//fmt.Printf("没有节点报告false，进入下一步判断\n")
				if agreeSum < len(members)/2 {
					fmt.Printf("同意信息未达到目标数值，增加序号值，重新发送共识请求\n")
					members[n].addNowSeq()
					maxAcceptorN = 0
					agreeSum = 0
				} else {
					fmt.Printf("同意信息达到目标值，进入第二阶段\n")
					acceptMessage := acceptMessage{members[n].nowSeq, maxAcceptorN, content}
					result := Accept(members, acceptMessage, n, messageSum, arr)
					if result == true {
						fmt.Printf("第二阶段成功，共识达成\n")
						fmt.Printf("输出此时的节点信息\n")
						members[n].setProposer(members[n].getSeq(),content,members[n].getSeq())
						for i:=0; i<len(members); i++{
							fmt.Printf("输出%d个节点",i+1,members[i])
							fmt.Printf("\n")
						}
						break for2
					} else {
						fmt.Printf("第二阶段失败，共识提议被拒绝，继续增加序号值，重新开始第一阶段\n")
						members[n].addNowSeq()
					}
				}
			//}
				flag++
			}
			flag = 1
			fmt.Printf("是否继续进行共识？Y/N\n")
			fmt.Scanf("%s",&tmp)
			if tmp == "N" {
				break for1
			}

		}
	}

}

//第二阶段
func Accept(members []Member,message acceptMessage,n int, sumMessage int, arr []int) bool{
	var agreeSum int
	fmt.Printf("进入第二阶段\n")
	for i:=0; i< sumMessage; i++{
		if members[arr[i]].Accept(message) == true {
			agreeSum++
		}
	}
	if agreeSum > len(members)/2{
		return true
	}else{
		return false
	}

}




func main()  {

	begin()
	var n int
	var flag string
	//var membersMap map[int] Member
	//membersMap = make(map[int] Member)
	fmt.Scanf("%d",&n)
	members := make([]Member,n)
	for i:=0; i<n; i++{
		members[i] = newMember()
	}

	fmt.Printf("%d个初始节点创建成功\n",n)
	//TODO
	fmt.Printf("是否需要设置节点的初试值为随机数？Y/N\n")
	fmt.Scanf("%s", &flag)
	if flag == "Y"{
		setRandomMembers(members)
		fmt.Printf("设置成功！\n")
		for i:=0; i<len(members); i++{
			fmt.Printf("输出%d个节点",i+1,members[i])
			fmt.Printf("\n")
		}
	}
	fmt.Printf("")
	propose(members)




}