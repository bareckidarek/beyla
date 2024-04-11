// Code generated by bpf2go; DO NOT EDIT.
//go:build arm64

package ebpf

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"

	"github.com/cilium/ebpf"
)

type NetFlowId NetFlowIdT

type NetFlowIdT struct {
	SrcIp             struct{ In6U struct{ U6Addr8 [16]uint8 } }
	DstIp             struct{ In6U struct{ U6Addr8 [16]uint8 } }
	EthProtocol       uint16
	Direction         uint8
	SrcPort           uint16
	DstPort           uint16
	TransportProtocol uint8
	IfIndex           uint32
}

type NetFlowMetrics NetFlowMetricsT

type NetFlowMetricsT struct {
	Packets         uint32
	Bytes           uint64
	StartMonoTimeNs uint64
	EndMonoTimeNs   uint64
	Flags           uint16
	Errno           uint8
}

type NetFlowRecordT struct {
	Id      NetFlowId
	Metrics NetFlowMetrics
}

// LoadNet returns the embedded CollectionSpec for Net.
func LoadNet() (*ebpf.CollectionSpec, error) {
	reader := bytes.NewReader(_NetBytes)
	spec, err := ebpf.LoadCollectionSpecFromReader(reader)
	if err != nil {
		return nil, fmt.Errorf("can't load Net: %w", err)
	}

	return spec, err
}

// LoadNetObjects loads Net and converts it into a struct.
//
// The following types are suitable as obj argument:
//
//	*NetObjects
//	*NetPrograms
//	*NetMaps
//
// See ebpf.CollectionSpec.LoadAndAssign documentation for details.
func LoadNetObjects(obj interface{}, opts *ebpf.CollectionOptions) error {
	spec, err := LoadNet()
	if err != nil {
		return err
	}

	return spec.LoadAndAssign(obj, opts)
}

// NetSpecs contains maps and programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetSpecs struct {
	NetProgramSpecs
	NetMapSpecs
}

// NetSpecs contains programs before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetProgramSpecs struct {
	EgressFlowParse  *ebpf.ProgramSpec `ebpf:"egress_flow_parse"`
	IngressFlowParse *ebpf.ProgramSpec `ebpf:"ingress_flow_parse"`
}

// NetMapSpecs contains maps before they are loaded into the kernel.
//
// It can be passed ebpf.CollectionSpec.Assign.
type NetMapSpecs struct {
	AggregatedFlows *ebpf.MapSpec `ebpf:"aggregated_flows"`
	DirectFlows     *ebpf.MapSpec `ebpf:"direct_flows"`
}

// NetObjects contains all objects after they have been loaded into the kernel.
//
// It can be passed to LoadNetObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetObjects struct {
	NetPrograms
	NetMaps
}

func (o *NetObjects) Close() error {
	return _NetClose(
		&o.NetPrograms,
		&o.NetMaps,
	)
}

// NetMaps contains all maps after they have been loaded into the kernel.
//
// It can be passed to LoadNetObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetMaps struct {
	AggregatedFlows *ebpf.Map `ebpf:"aggregated_flows"`
	DirectFlows     *ebpf.Map `ebpf:"direct_flows"`
}

func (m *NetMaps) Close() error {
	return _NetClose(
		m.AggregatedFlows,
		m.DirectFlows,
	)
}

// NetPrograms contains all programs after they have been loaded into the kernel.
//
// It can be passed to LoadNetObjects or ebpf.CollectionSpec.LoadAndAssign.
type NetPrograms struct {
	EgressFlowParse  *ebpf.Program `ebpf:"egress_flow_parse"`
	IngressFlowParse *ebpf.Program `ebpf:"ingress_flow_parse"`
}

func (p *NetPrograms) Close() error {
	return _NetClose(
		p.EgressFlowParse,
		p.IngressFlowParse,
	)
}

func _NetClose(closers ...io.Closer) error {
	for _, closer := range closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}
	return nil
}

// Do not access this directly.
//
//go:embed net_bpfel_arm64.o
var _NetBytes []byte
