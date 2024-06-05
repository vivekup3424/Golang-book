package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

// Load and attach the eBPF program to the appropriate hook (e.g., XDP)
func main() {
	// Allow the current process to lock memory for eBPF maps
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to set rlimit: %v", err)
	}

	// Load eBPF objects
	var objs process_packet_filterObjects
	if err := loadProcess_packet_filterObjects(&objs, nil); err != nil {
		log.Fatalf("Loading eBPF objects failed: %v", err)
	}
	defer objs.Close()

	// Attach the eBPF program to a network interface
	// Change "eth0" to your network interface name
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   objs.BlockProcessPorts,
		Interface: 0, // replace with the actual network interface index
		Flags:     link.XDPGenericMode,
	})
	if err != nil {
		log.Fatalf("Failed to attach XDP program: %v", err)
	}
	defer link.Close()

	log.Println("eBPF program successfully loaded and attached.")

	// Wait for a signal (e.g., Ctrl+C) to exit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("Received signal, exiting program.")
}
