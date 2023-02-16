package main

import (
	"TCP_Packet/routers"
	"TCP_Packet/utils/leveldb_utils"
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
	srcHost1 := settings.GetUDPCollector().SrcUDP1
	srcHost2 := settings.GetUDPCollector().SrcUDP2
	interfaceCollector := settings.GetUDPCollector().Interface
	port := uint16(settings.GetUDPCollector().Port) // specify the port number you want to capture
	sizeCapture := settings.GetUDPCollector().SizeCapture
	numberThread := settings.GetSystem().NumberThread
	udpMappingHost := settings.GetUDPMapping().Host
	udpMappingPort := settings.GetUDPMapping().Port
	//Read ips from file to sharedIPS at starting program
	glog.Info("Readed IPS from file to sharedIPS var at starting program")
	//global_ips.UpdateSetIPS(file.ReadIPSFromFile())
	glog.Info(fmt.Sprintf("Service Processing UDP for Interface %s and from source host: %s and %s Packet Runing ", interfaceCollector, srcHost1, srcHost2))
	handle, err := pcap.OpenLive(interfaceCollector, sizeCapture, true, pcap.BlockForever)
	if err != nil {
		glog.Error(err)
	}
	defer handle.Close()
	// Set a filter to only capture UDP packets on the specified port
	filter := fmt.Sprintf("udp and port %d and (src host %s or src host %s)", port, srcHost1, srcHost2)
	err = handle.SetBPFFilter(filter)
	if err != nil {
		glog.Error(err)
	}
	// Start capturing packets
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	packets := packetSource.Packets()
	//// Start multiple goroutines to process packets
	//setIPs := global_ips.GetIPS()
	levelDB := leveldb_utils.GetLevelDBConnection()
	for i := 0; i < numberThread; i++ {
		go worker.WorkerProcessingUDP(udpMappingHost, udpMappingPort, packets, levelDB)
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
	fmt.Println("Starting Service: " + listenTo)
	go http.ListenAndServe(listenTo, nApi)
	//Scheduled write ips to file daily midnight
	//go file.ScheduleWriteIPSToFile()
	//Wait for the goroutines to finish
	select {}
}
