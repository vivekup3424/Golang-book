package main

import (
	"log"
	"net"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

func main() {
	//remove memory limits for kernels<5.11
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal("Removing memlock: ", err)
	}

	//get the compiled eBPF ELF and load it into the kernel
	var obj xdpPassKernelObjects
	if err := loadXdpPassKernelObjects(&obj, nil); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer obj.Close()

	ifname := "wlp0s20f3" //interface name on my machine
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatalf("Getting interface %s : %s\n", ifname, err)
	}

	//attach my ELF objects to the interface
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   obj.XdpProgSimple,
		Interface: iface.Index,
	})
	if err != nil {
		log.Fatalf("Attaching XDP to NIC: %s\n", err)
	}
	//close the link when after using it
	defer link.Close()

	log.Printf("Monitoring the packets coming on: %s...\n", ifname)
	
}
