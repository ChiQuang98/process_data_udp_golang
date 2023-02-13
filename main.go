package main

import (
	"TCP_Packet/routers"
	"TCP_Packet/utils/settings"
	"TCP_Packet/worker"
	"flag"
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/golang/glog"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/rs/cors"
	"net/http"
	"os"
	"strconv"
)

func init() {
	//glog
	//create logs folder
	os.Mkdir("./logs", 0777)
	flag.Lookup("stderrthreshold").Value.Set("[INFO|WARN|FATAL]")
	flag.Lookup("logtostderr").Value.Set("false")
	flag.Lookup("alsologtostderr").Value.Set("true")
	flag.Lookup("log_dir").Value.Set("./logs")
	glog.MaxSize = 1024 * 1024 * settings.GetGlogConfig().MaxSize
	flag.Lookup("v").Value.Set(fmt.Sprintf("%d", settings.GetGlogConfig().V))
	flag.Parse()

}
func main() {
	srcHost := settings.GetUDPCollector().SrcUDP
	interfaceCollector := settings.GetUDPCollector().Interface
	port := uint16(settings.GetUDPCollector().Port) // specify the port number you want to capture
	sizeCapture := settings.GetUDPCollector().SizeCapture
	numberThread := settings.GetSystem().NumberThread
	udpMappingHost := settings.GetUDPMapping().Host
	udpMappingPort := settings.GetUDPMapping().Port
	glog.Info(fmt.Sprintf("Service Processing UDP for Interface %s and from source host: %s Packet Runing ", interfaceCollector, srcHost))
	handle, err := pcap.OpenLive(interfaceCollector, sizeCapture, true, pcap.BlockForever)
	if err != nil {
		glog.Error(err)
	}
	defer handle.Close()
	// Set a filter to only capture UDP packets on the specified port
	filter := fmt.Sprintf("udp and port %d and src host %s", port, srcHost)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		glog.Error(err)
	}
	// Start capturing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	//
	//// Start multiple goroutines to process packets
	for i := 0; i < numberThread; i++ {
		go worker.WorkerProcessingUDP(udpMappingHost, udpMappingPort, packets)
	}
	routerApi := routers.InitRoutes()
	nApi := negroni.Classic()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"DELETE", "PUT", "GET", "HEAD", "OPTIONS", "POST"},
	})
	nApi.Use(c)
	nApi.UseHandler(routerApi)
	listenTo := settings.GetRestful().Host + ":" + strconv.Itoa(settings.GetRestful().Port)
	fmt.Println(listenTo)
	go http.ListenAndServe(listenTo, nApi)
	//Wait for the goroutines to finish
	select {}
}
