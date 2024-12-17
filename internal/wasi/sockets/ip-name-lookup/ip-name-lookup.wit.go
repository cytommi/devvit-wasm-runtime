// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package ipnamelookup represents the imported interface "wasi:sockets/ip-name-lookup@0.2.0".
package ipnamelookup

import (
	"devvit-wasm/internal/wasi/io/poll"
	"devvit-wasm/internal/wasi/sockets/network"
	"go.bytecodealliance.org/cm"
)

// Pollable represents the imported type alias "wasi:sockets/ip-name-lookup@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// Network represents the imported type alias "wasi:sockets/ip-name-lookup@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// ErrorCode represents the type alias "wasi:sockets/ip-name-lookup@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddress represents the type alias "wasi:sockets/ip-name-lookup@0.2.0#ip-address".
//
// See [network.IPAddress] for more information.
type IPAddress = network.IPAddress

// ResolveAddressStream represents the imported resource "wasi:sockets/ip-name-lookup@0.2.0#resolve-address-stream".
//
//	resource resolve-address-stream
type ResolveAddressStream cm.Resource

// ResourceDrop represents the imported resource-drop for resource "resolve-address-stream".
//
// Drops a resource handle.
//
//go:nosplit
func (self ResolveAddressStream) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResourceDrop((uint32)(self0))
	return
}

// ResolveNextAddress represents the imported method "resolve-next-address".
//
//	resolve-next-address: func() -> result<option<ip-address>, error-code>
//
//go:nosplit
func (self ResolveAddressStream) ResolveNextAddress() (result cm.Result[OptionIPAddressShape, cm.Option[IPAddress], ErrorCode]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_ResolveAddressStreamResolveNextAddress((uint32)(self0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self ResolveAddressStream) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_ResolveAddressStreamSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// ResolveAddresses represents the imported function "resolve-addresses".
//
//	resolve-addresses: func(network: borrow<network>, name: string) -> result<resolve-address-stream,
//	error-code>
//
//go:nosplit
func ResolveAddresses(network_ Network, name string) (result cm.Result[ResolveAddressStream, ResolveAddressStream, ErrorCode]) {
	network0 := cm.Reinterpret[uint32](network_)
	name0, name1 := cm.LowerString(name)
	wasmimport_ResolveAddresses((uint32)(network0), (*uint8)(name0), (uint32)(name1), &result)
	return
}
