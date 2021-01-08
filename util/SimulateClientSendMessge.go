package util

import (
	"casino_common/proto/ddproto"
	commonNewPorot "casino_common/proto/funcsInit"
	"casino_common/utils/security"
	"casino_hall/msg"
	"crypto/md5"
	"fmt"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/websocket"
)

//发消息测试,模拟客户端操作
func WebSocketTest(url string,message ...interface{})  {
	//var url = "ws://47.110.72.104:3801"
	//var url = "ws://120.77.201.111:3801"
	var origin = "http://localhost:443//"
	//var protocol = "protobuf"//这个不知道有啥用,没搞懂
	var protocol = ""
	//连接ws  访问地址  协议  请求头
	ws,err := websocket.Dial(url,protocol,origin)
	if err != nil {
		fmt.Println(err)
	}
	//获取请求
	req := GetMsgReq()//默认请求
	if len(message) != 0 {//传了消息过来则用传过来的结构
		req = message[0]
	}
	//序列化 返回序列化数组,由消息号跟消息内容组成
	//这里内部实际用的还是proto.Marshal() ,只不过多返回一个内部注册的消息号数组(约定好的)
	//具体可以看下面注释掉的客户端代码实现
	data,err := msg.Processor.Marshal(req)//取到消息id和真实数据
	if err != nil {
		fmt.Println("Marshal失败:",err)
		return
	}
	//拼装消息号和内容成一个数组
	//data2 => []byte{msgId,msg...}
	data2 := append(data[0],data[1]...)//拼装成一个
	//做md5加密,服务器会做校验
	var md5data []byte
	copy(md5data,data2)
	md5data = append(data2,security.SECRET_KEY[0],security.SECRET_KEY[1],security.SECRET_KEY[2],security.SECRET_KEY[3])
	h := md5.New()
	h.Write(md5data)
	resultByte := h.Sum(nil)
	//封装校验数据到消息
	//data2 => []byte{msgId,msg...,md5[4],md5[6],md5[8],md5[10]}
	data2 = append(data2,resultByte[4],resultByte[6],resultByte[8],resultByte[10])

	//发送消息
	if _,err = ws.Write(data2);err != nil {
		fmt.Println(err)
	}
	//暂时没有接受消息模块

	/*//客户端代码,消息加密代码
	var SECRET_KEY = new Uint8Array([0x93, 0x46, 0x78, 0x20] ); //MD5签名密钥
	var buf = new ArrayBuffer(2 + msgBuf.length + 4);
	var bytes = new Uint8Array(buf)
	//填充header
	var dataLen = buf.byteLength;
	var msgId = id;
	bytes[0] = (msgId & 0xff00) >> 8;
	bytes[1] = (msgId & 0x00ff);

	//填充proto message
	for (var i = 0; i < msgBuf.length; i++) {
		bytes[2 + i] = msgBuf[i];
	}

	//添上MD5_Key
	for (var i = 0; i < SECRET_KEY.byteLength; i++) {
		bytes[2 + msgBuf.length + i] = SECRET_KEY[i];
	}

	//计算MD5签名
	var sign = SparkMD5.ArrayBuffer.hash(buf); //SparkMD5.hashBinary(string);
	//AppLog.log(" =========  md5:" + sign);

	var md5 = [];
	for (var x = 0; x < sign.length - 1; x += 2) {
		md5.push(parseInt(sign.substr(x, 2), 16));
	}

	bytes[2 + msgBuf.length + 0] = md5[4];
	bytes[2 + msgBuf.length + 1] = md5[6];
	bytes[2 + msgBuf.length + 2] = md5[8];
	bytes[2 + msgBuf.length + 3] = md5[10];*/
}
//获取消息
func GetMsgReq() interface{} {
	req := &ddproto.GroupReqSetClose{Header: commonNewPorot.NewHeader(),GroupId: proto.Int32(416199),Closing: proto.Bool(true)}
	req.Header.UserId = proto.Uint32(19924)
	return req
}
