package main

import "Structured_Logging/logger"

func main() {
	//log.Println("Usage of Logger!")  ===> o/p:  2021/10/26 15:22:53 Usage of Logger!
	//To avoid this and have structured error log we use this uber zap
	//logger.Log.Info("Use of Logger") ---> returns the caller as structured_logger/main.go
	logger.Info("Use of Logger!!") //"caller":"logger/logger.go
}
