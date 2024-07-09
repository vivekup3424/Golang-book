// Code generated by bpf2go; DO NOT EDIT.
//go:build mips || mips64 || ppc64 || s390x

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

// loadPacketFilter returns the embedded CollectionSpec for packetFilter.
func loadPacketFilter() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_PacketFilterBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load packetFilter: %w", err)
	}

	return spec, err
}

// loadPacketFilterObjects loads packetFilter and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*packetFilterObjects
//	*packetFilterPrograms
//	*packetFilterMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func loadPacketFilterObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := loadPacketFilter()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// packetFilterSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetFilterSpecs struct {
	packetFilterProgramSpecs
	packetFilterMapSpecs
}

// packetFilterSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetFilterProgramSpecs struct {
	BlockProcessPorts *ebpf.ProgramSpec `ebpf:"block_process_ports"`
}

// packetFilterMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type packetFilterMapSpecs struct {
	DropCounter    *ebpf.MapSpec `ebpf:"drop_counter"`
	PortMap        *ebpf.MapSpec `ebpf:"port_map"`
	ProcessNameMap *ebpf.MapSpec `ebpf:"process_name_map"`
}

// packetFilterObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to loadPacketFilterObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetFilterObjects struct {
	packetFilterPrograms
	packetFilterMaps
}

func (o *packetFilterObjects) Close() error {
	return _PacketFilterClose(
		&o.packetFilterPrograms,
		&o.packetFilterMaps,
	)
}

// packetFilterMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to loadPacketFilterObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetFilterMaps struct {
	DropCounter    *ebpf.Map `ebpf:"drop_counter"`
	PortMap        *ebpf.Map `ebpf:"port_map"`
	ProcessNameMap *ebpf.Map `ebpf:"process_name_map"`
}

func (m *packetFilterMaps) Close() error {
	return _PacketFilterClose(
		m.DropCounter,
		m.PortMap,
		m.ProcessNameMap,
	)
}

// packetFilterPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to loadPacketFilterObjects or ebpf.CollectionSpec.LoadAndAssign.
type packetFilterPrograms struct {
	BlockProcessPorts *ebpf.Program `ebpf:"block_process_ports"`
}

func (p *packetFilterPrograms) Close() error {
	return _PacketFilterClose(
		p.BlockProcessPorts,
	)
}

func _PacketFilterClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed packetfilter_bpfeb.o
var _PacketFilterBytes []byte