package mobile

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/newinfoOffical/core"
	"github.com/newinfoOffical/core/mobile/networkInterface"
	"github.com/newinfoOffical/core/webapi"
	"net/http"
	"runtime"
	"time"
)

func MobileMain(path string, Networkinterface string) {
	// LockOSThread wires the calling goroutine to its current operating system thread.
	runtime.LockOSThread()

	// Parsing the network interface got as a string from the app
	if Networkinterface == "" {
		print("networkInterface string empty \n")
		return
	}
	err := networkInterface.InitInterfaces(Networkinterface)

	if err != nil {
		print("networkInterface", "parsing error: %s\n", err.Error())
		return
	}

	var config core.Config

	// Load the config file
	core.LoadConfig(path+"Config.yaml", &config)

	//Setting modified paths in the config file
	config.SearchIndex = path + "data/search_Index/"
	config.BlockchainGlobal = path + "data/blockchain/"
	config.BlockchainMain = path + "data/blockchain_main/"
	config.WarehouseMain = path + "data/warehouse/"
	config.GeoIPDatabase = path + "data/GeoLite2-City.mmdb"
	config.LogFile = path + "data/log.txt"

	// save modified config changes
	core.SaveConfig(path+"Config.yaml", &config)

	backendInit, status, err := core.Init("Your application/1.0", path+"Config.yaml", nil, nil)
	if status != core.ExitSuccess {
		fmt.Printf("Error %d initializing config: %s\n", status, err.Error())
		return
	}

	// start config api server
	webapi.Start(backendInit, []string{"127.0.0.1:5125"}, false, "", "", 10*time.Second, 10*time.Second, uuid.Nil)

	backendInit.Connect()

	// Checks if the go code can access the internet
	if !connected() {
		fmt.Print("Not connected to the internet ")
	} else {
		fmt.Print("Connected")
	}
	/*http.HandleFunc("/test", ReceiveFile)
	log.Fatal(http.ListenAndServe(":8080", nil))*/
	//KotlinServer.RequestFile("")

}

func connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
