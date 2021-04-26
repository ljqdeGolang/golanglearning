// func main() {
// 	var o options
// 	iniName :=SetConfigFile("ec20config.ini")
// 	iniName(&o)
// 	cfg ,err :=ini.Load(o.ConfigFile)
// 	if err !=nil {
// 		fmt.Printf("Fail to read configfile,the error is:%v",err)
// 	}
// 	filepath :=SetLogFilePath(cfg.Section("log").Key("filePath").String())
// 	fileName :=SetLogFileName(cfg.Section("log").Key("fileName").String())
// 	filepath(&o)
// 	fileName(&o)

// 	//设置全局日志
// 	log,err :=IniLog(o.LogFilePath,o.LogFileName)
// 	if err !=nil {
// 		panic(err)
// 	}
// 	config ,err :=config.LoadINI(o.ConfigFile, log)
// 	if err !=nil {
// 		log.WithFields(logger.Fields{
// 			"system":"init",
// 		}).Panic("Load config ini file:",err)
// 	}

// 	link,err := IsCanPingOpen(config.Net.Dns)
// 	if err != nil {
// 		//将错误记录在日志中
// 		log.WithFields(logger.Fields{
// 			"network":"ping",
// 		}).Info("ping was wrong:",err)
// 	}
// 	if link {
// 		//将结果记录在日志中
// 		log.WithFields(logger.Fields{
// 			"network":"ping",
// 		}).Info("network is normal")
// 		return
// 	} else {
// 		//重新连接网络，调用Reconnection函数
// 		//循环调用Reconnection函数，直到isOK为真
// 		var i bool = false
// 		var n int
// 		for ;!i;n++ {
// 			isOK,err :=Reconnection(log)
// 			if err != nil {
// 				log.WithFields(logger.Fields{
// 					"network":"reconection",
// 				}).Error("reconnection is wrong:",err)
				
// 			}
// 			i = isOK

// 		}
// 		//这里再ping，不会成功，可能因为会执行脚本的时间较长，还没联网成功
// 		Isnet,err :=IsCanPingOpen(config.Net.Dns)
// 		if Isnet {
// 			log.WithFields(logger.Fields{
// 				"network":"re-ping",
// 			}).Info("reconnection is sucessful")
// 		} else {
// 			log.WithFields(logger.Fields{
// 				"network":"re-ping",
// 			}).Error("reconection failed",err)
// 		}
// 		return
// 	}
// }


// //Reconcetion 函数是重新联网的方法
// func Reconnection(log *logger.Logger) (bool,error) {
// 	//要判断.sh文件相关进程是否在后台运行，若在就需要终止运行进程
// 	excmd :=exec.Command("/bin/bash","-c","ps -aux | grep pppd | grep -v grep")
// 	std_out, _ := excmd.Output()
// 	runtimingProcess := strings.Split(string(std_out),"\n")
// 	for _, str :=range runtimingProcess {
// 		splitString :=strings.Fields(str)
// 		var isEnd bool =false
// 		for _,fragment :=range splitString {
// 			//这里检验是否包含pppd的后台进程
// 			if strings.Contains(fragment,"pppd") {
// 				killCommand :="sudo kill -9 "+splitString[2]
// 				cmd :=exec.Command("/bin/bash","-c",killCommand)
// 				out_bytes,err :=cmd.Output()
// 				if err !=nil {
// 					log.WithFields(logger.Fields{
// 						"execution":"kill",
// 					}).Error("file execution failed:",err)			
// 					return false,nil
// 				}
// 				out_string :=string(out_bytes)
// 				isEnd =strings.Contains(out_string,"已杀死")
// 				if isEnd {
// 					log.WithFields(logger.Fields{
// 						"network":"kill",
// 					}).Info(out_string)
// 					break
// 				}
// 			}
// 		}
// 		if isEnd { break }
// 	}
// 	//先授予.sh文件执行权限，再执行.sh文件
// 	cmdd :=exec.Command("/bin/bash","-c","chmod u+x quectel-pppd.sh")
// 	stdoutbyte,err :=cmdd.Output()
// 	if err !=nil {
// 		log.WithFields(logger.Fields{
// 			"execution":"chmod",
// 		}).Error("chmod output failed:",err)	
// 		return false,err		
// 	} else {
// 		log.WithFields(logger.Fields{
// 			"execution":"chmod",
// 		}).Info("chmod succeeded:"+string(stdoutbyte))		
// 	}
// 	isNote,err :=ExecCommand("sudo ./quectel-pppd.sh",log)
// 	if err !=nil {
// 		log.WithFields(logger.Fields{
// 			"execution":"sh",
// 		}).Error("file execution failed:",err)
// 		return isNote,err
// 	}
// 	return isNote,nil
// }