// Package consts for storing util consts
package consts

import (
	"net"
)

// Type 1 devices: Heaters (v2, touch, v4, Heater), Plug
// Type 2 devices: Breeze, Runners

// Port types and values
const (
	UDPPortType1    = 20_002
	UDPPortType1New = 10_002
	UDPPortType2    = 20_003
	UDPPortType2New = 10_003
	TCPPortType1    = 9_957
	TCPPortType2    = 10_000
)

// Device catgories
const (
	DeviceCategoryWaterHeater = iota
	DeviceCategoryPowerPlug
	DeviceCategoryThermostat
	DeviceCategoryShutter
)

// Device types
const (
	DeviceTypeMini = iota
	DeviceTypePowerPlug
	DeviceTypeTouch
	DeviceTypeV2ESP
	DeviceTypeV2QCA
	DeviceTypeV4
	DeviceTypeBreeze
	DeviceTypeRunner
	DeviceTypeRunnerMini
)

// Message lengths for different devices
const (
	MessageLengthDefault = 165
	MessageLengthBreeze  = 168
	MessageLengthRunner  = 159
)

// Login packets
const (
	LoginPacketType1Template = "fef052000232a10000000000340001000000000000000000%s00000000000000000000f0fe%s00000000000000000000000000000000000000000000000000000000000000000000000000"
)

// DefaultIP the fallback IP
var DefaultIP = net.IP{10, 100, 102, 82}

// Devices join maps
var (
	DeviceCategoryToUDPPort = map[int]int{
		DeviceCategoryWaterHeater: UDPPortType1New,
		DeviceCategoryPowerPlug:   UDPPortType1,
		DeviceCategoryThermostat:  UDPPortType2,
		DeviceCategoryShutter:     UDPPortType2,
	}
	DeviceCategoryToTCPPort = map[int]int{
		DeviceCategoryWaterHeater: TCPPortType1,
		DeviceCategoryPowerPlug:   TCPPortType1,
		DeviceCategoryThermostat:  TCPPortType2,
		DeviceCategoryShutter:     TCPPortType2,
	}
	DeviceTypeToCategory = map[int]int{
		DeviceTypeMini:       DeviceCategoryWaterHeater,
		DeviceTypePowerPlug:  DeviceCategoryPowerPlug,
		DeviceTypeTouch:      DeviceCategoryWaterHeater,
		DeviceTypeV2ESP:      DeviceCategoryWaterHeater,
		DeviceTypeV2QCA:      DeviceCategoryWaterHeater,
		DeviceTypeV4:         DeviceCategoryWaterHeater,
		DeviceTypeBreeze:     DeviceCategoryThermostat,
		DeviceTypeRunner:     DeviceCategoryShutter,
		DeviceTypeRunnerMini: DeviceCategoryShutter,
	}
)
