package worker

import (
	"TCP_Packet/utils/global_ips"
	"TCP_Packet/utils/sys_log"
	"TCP_Packet/utils/udp_utils"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"regexp"
)

func WorkerProcessingUDP(hostUDP string, portUDP int, packets <-chan gopacket.Packet) {
	for packet := range packets {
		// Process the packet here

		// Get the UDP layer from the packet
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			setIPs := global_ips.GetIPS()
			// Get the UDP header from the layer
			udpHeader, _ := udpLayer.(*layers.UDP)
			re := regexp.MustCompile(`nat-destination-address="([^"]+)"`)
			//fmt.Println(string(udpHeader.Payload))
			match := re.FindSubmatch(udpHeader.Payload)
			if len(match) > 1 {
				if _, ok := setIPs[string(match[1])]; ok {
					//fmt.Printf("Worker %d:\n", id)
					sys_log := sys_log.ParseSyslog((udpHeader.Payload))
					println(sys_log["nat-destination-address"])
					fmt.Println("Destination address:", string(udpHeader.Payload))
					udp_utils.SendPacket(hostUDP, portUDP, udpHeader.Payload)
				}
			}
			//fmt.Printf("  Payload: %v\n", )
		}
		//fmt.Println()
	}
}
