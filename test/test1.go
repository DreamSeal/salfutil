package test

import (
	"casino_common/proto/ddproto"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func slice() {
	var array = [3]int32{1, 2, 3}
	//fmt.Print(reflect.TypeOf(array).String())
	arr := array[0:]
	_ = append(arr, arr...)
	var a *ddproto.GroupMemberInfo = nil
	_ = a.GetUid()
}

func main() {
	for i := 0; i < 10; i++ {
		RedPackage(10, 500)
		fmt.Println("")
	}
}
// 随机红包// remainCount: 剩余红包数// remainMoney: 剩余红包金额（单位：分)
func randomMoney(remainCount, remainMoney int) int {
	if remainCount == 1 {
		return remainMoney
	}

	//rand.Seed(time.Now().UnixNano())
	var min = 1
	//保证每人都有1分
	if remainMoney == remainCount {
		return min
	}
	remainMoney = remainMoney - remainCount * min
	//最多能拿到  剩下每人平均数的2倍
	max := remainMoney / remainCount * 2
	money := rand.Intn(max) + min
	return money
}
// 发红包// count: 红包数量// money: 红包金额（单位：分)
func RedPackage(count, money int) {
	for i := 0; i < count; i++ {
		m := randomMoney(count-i, money)
		fmt.Printf("%d  ", m)
		if m == 0 {
			fmt.Println("buxing")
		}
		money -= m
	}
	fmt.Println(" ")
}


//红包
type RedPacket struct {
	Rid				int32	//红包id，区分不同的红包,用redis做自增
	userid			uint32	//发起人
	GroupId			int32	//群id
	CreateTime		int64	//创建时间
	LiveTime		int64	//红包存活时间，防止后续要求时间可选
	Count			int32 	//红包总数
	Score			int32	//红包总分
	RemainCount		int32	//剩余个数
	RemainScore		int32	//剩余分，未抢光时,该字段同时可以表示返还的分数
	BestScore 		int32	//手气最佳分数  防止要先加上
	BestUserid		uint32	//手气最佳用户  防止要先加上
	ScoreList		[]*GrabRedPacketInfo //每个人的领取信息
	Status			int32	//状态(1.开抢中  2.已抢光  3.时间结束未抢光,剩余积分返还)
}


//抢红包信息的回复
type GrabRedPacketInfo struct {
	Rid				int32	//红包id，区分不同的红包
	userid			uint32	//用户
	GroupId			int32	//群id
	Time			string	//时间
	Score			int64	//抢到的分数
}