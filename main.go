package main

import (
	"App"
	Web "HillWeb"
	"Libs/Log"
	"Libs/Sync"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func init() {
	settingBytes, err := os.ReadFile("settings.json")
	if err != nil {
		log.Fatalf("[[main.init]] read settings error: %v\n", err)
		return
	}
	err = App.InitMode(settingBytes)
	//err = json.Unmarshal(settingBytes, models.Config)
	if err != nil {
		log.Fatalf("[[main.init]] json unmarshal error: %v\n", err)
		return
	}

	App.InitDb()
	//Log.Error.Print("OnInit App.Ver=", App.Ver, " Config=", App.Config.Host.Version)
}

type testRun struct {
	value int
}

func (this testRun) RunSync() {
	time.Sleep(time.Second * 1)
	Log.Warn.Print("RunSync test val:", this.value)
	// this.ch <- this.value
}

func main() {
	fmt.Println("Hello Hill Api")

	//var channel = make(chan int)

	Sync.Run(func() {
		var val = 1
		time.Sleep(time.Second * 1)
		Log.Debug.Print("Run func val:", val)
	}, func() {
		var val = 2
		time.Sleep(time.Second * 1)
		Log.Debug.Print("Run func val:", val)
	}, func() {
		var val = 3
		time.Sleep(time.Second * 1)
		Log.Debug.Print("Run func val:", val)
	})

	// Sync.Run_Seq(func() {
	// 	var val = 1
	// 	time.Sleep(time.Second * 1)
	// 	Log.Error.Print("Main Run Seq func val:", val)

	// }, func() {
	// 	var val = 2
	// 	time.Sleep(time.Second * 1)
	// 	Log.Error.Print("Main Run Seq func val:", val)

	// }, func() {
	// 	var val = 3
	// 	time.Sleep(time.Second * 1)
	// 	Log.Error.Print("Main Run Seq func val:", val)

	// })

	Sync.RunThread_Seq(testRun{
		value: 1,
	}, testRun{
		value: 2,
	}, testRun{
		value: 3,
	})

	// for {
	// 	var v = <-channel
	// 	Log.Warn.Print("testRun val:", v)
	// 	if v == 2 {
	// 		break
	// 	}
	// }

	// select {
	// case v := <-channel:
	// 	Log.Warn.Print("testRun val:", v)
	// default:
	// 	Log.Error.Print("testRun end")

	// }

	// App.RunMode(App.Mode_Dev, func() {
	// 	Log.D("This is Dev Mode")
	// })
	// App.RunMode(App.Mode_Http, func() {
	// 	Log.W("This is http Mode")
	// })
	// Log.Debug.Print("test Log.Debug.Print ", "v1=", 1, " v2=", true)
	// Log.Warn.Printp("test Log.Warn.Printp ", "v1=", 1, " v2=", true)

	//Ext.TypeEqual[bool](false)
	//Enum.Init()

	//router := mux.NewRouter()

	initWebRouter()

}

func initWebRouter() {

	router := Web.InitFromGorilla()
	cors := Web.InitCors()

	http.ListenAndServe(":"+App.Web.Port, cors.Handler(router))

	Log.Info.Printf("Server(ver: %s) run on port: %s...", App.Ver, App.Web.Port)
}
