package worker

import (
	"TCP_Packet/utils/global_ips"
	"TCP_Packet/utils/sys_log"
	"TCP_Packet/utils/udp_utils"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/syndtr/goleveldb/leveldb"
	"regexp"
)

func WorkerProcessingUDP(hostUDP string, portUDP int, packets <-chan gopacket.Packet, db *leveldb.DB) {
	for packet := range packets {
		// Process the packet here
		// Get the UDP layer from the packet
		udpLayer := packet.Layer(layers.LayerTypeUDP)
		if udpLayer != nil {
			//setIPs := global_ips.GetIPS()
			// Get the UDP header from the layer
			udpHeader, _ := udpLayer.(*layers.UDP)
			re := regexp.MustCompile(`nat-destination-address="([^"]+)"`)
			//fmt.Println(string(udpHeader.Payload))
			match := re.FindSubmatch(udpHeader.Payload)
			if len(match) > 1 {
				//Neu key la ip ton tai trong leveldB
				//match[1] is nat-destination-address
				if global_ips.CheckIPSExsist(db, match[1]) {
					msg := sys_log.ParseSyslog((udpHeader.Payload))
					udp_utils.SendPacket(hostUDP, portUDP, []byte(msg))
				}
			}
		}
	}
}
