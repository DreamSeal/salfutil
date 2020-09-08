package main

import (
	_ "net/http/pprof"
)

//func init() {
//	fmt.Println("初始化package main init()")
//	e := sys.SysInit(
//		conf.Server.ReleaseTag,
//		conf.Server.ProdMode,
//		conf.Server.RedisAddr,
//		conf.Server.DBName,
//		conf.Server.RedisPwd,
//		conf.Server.LogPath,
//		"hall",
//		conf.Server.LogFileSize,
//		conf.Server.LogFileCount,
//		conf.Server.MongoIp,
//		conf.Server.MongoLogIp,
//		conf.Server.DBName,
//		[]string{tableName.DBT_T_PAYBASEDETAILS, tableName.DBT_T_USER})
//	//初始化系统
//	if e != nil {
//		fmt.Printf("初始化的时候出错:%v", e)
//		os.Exit(-1)
//	}
//
//	//初始化微信支付回调
//	service.DoInitWxpay(conf.Server.WXPAY_NOTIFYURL)
//
//	//初始化灰度的白名单
//	model.GetGrayReleaseUsers()
//}

func main() {
	////打印日志
	//lconf.LogLevel = conf.Server.LogLevel
	//lconf.LogPath = conf.Server.LogPath
	//lconf.ConsolePort = conf.Server.ConsolePort
	//
	////微信支付配置
	//service.WXConfig.ApiKey = conf.Server.WxApikey
	//service.WXConfig.MchId = conf.Server.WxMchid
	//service.WXConfig.AppId = conf.Server.WxAppid
	//
	////启动rpc监听
	//_, _ = rpc_handler.LisenAndServeHallRPC(conf.Server.HallRpcAddr)
	////初始化rpc客户端
	//_ = rpcService.ZzHzPool.Init(conf.Server.MjZzHzRpcAddr, 1)
	//_ = rpcService.PdkPool.Init(conf.Server.PdkRpcAddr, 1)
	//_ = rpcService.SjhMjPool.Init(conf.Server.SjhRpcAddr, 1)
	//_ = rpcService.NiuniuPool.Init(conf.Server.NiuRpcAddr, 1)
	//_ = rpcService.SiChuanPool.Init(conf.Server.MjSiChuanRpcAddr, 1)
	//_ = rpcService.ZhadanPool.Init(conf.Server.ZhadanRpcAddr, 1)
	////_ = rpcService.LiuZhouPool.Init(conf.Server.MjLiuZhouRpcAddr, 1)
	//_ = rpcService.PaoHuZiPool.Init(conf.Server.PaoHuZiRpcAddr, 1)
	////_ = rpcService.PaoHuZiPool.Init(conf.Server.PaoHuZiSYRpcAddr, 1)
	//_ = rpcService.DaTongZiPool.Init(conf.Server.DaTongZiRpcAddr,1)
	//
	////启动微信支付回调
	//go func() {
	//	wxpay.WxpayCb()
	//}()
	//
	////启动调试端口
	//go func() {
	//	log.T("尝试开启pprof监听调试端口[23801]")
	//	log.T("ListenAndServe:23801 [%v]", http.ListenAndServe("0.0.0.0:23801", nil))
	//}()
	//
	//leaf.Run(
	//	game.Module,
	//	gate.Module,
	//	login.Module,
	//	//rpc.Module,
	//)
}
