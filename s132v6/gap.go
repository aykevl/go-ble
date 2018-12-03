package sd

import (
	"device/arm"
)

// Advertisement encapsulates a single advertisement instance.
type Advertisement struct {
	handle uint8
}

// NewAdvertisement creates a new advertisement instance but does not configure
// it.
func NewAdvertisement() *Advertisement {
	return &Advertisement{
		handle: 0xff, // BLE_GAP_ADV_SET_HANDLE_NOT_SET
	}
}

// AdvData contains the byte arrays (as strings) to broadcast.
type AdvData struct {
	AdvData     string // broadcasted data
	ScanRspData string // data returned with a scan response
}

type AdvProperties struct {
	Type   AdvType // the advertising type
	Fields uint8 // anonymous:1; include_tx_power:1
}

type AdvType uint8

// Advertisement types.
const (
	AdvTypeConnectableScannableUndirected               = 0x01
	AdvTypeConnectableNonscannableDirectedHighDutyCycle = 0x02
	AdvTypeConnectableNonscannableDirected              = 0x03
	AdvTypeNonconnectableScannableUndirected            = 0x04
	AdvTypeNonconnectableNonscannableUndirected         = 0x05
	AdvTypeExtendedConnectableNonscannableUndirected    = 0x06
	AdvTypeExtendedConnectableNonscannableDirected      = 0x07
	AdvTypeExtendedNonconnectableScannableUndirected    = 0x08
	AdvTypeExtendedNonconnectableScannableDirected      = 0x09
	AdvTypeExtendedNonconnectableNonscannableUndirected = 0x0A
	AdvTypeExtendedNonconnectableNonscannableDirected   = 0x0B
)

type PeerAddr struct {
	// TODO
}

// AdvParams configures everything related to BLE advertisements.
type AdvParams struct {
	Properties   AdvProperties
	PeerAddr     *PeerAddr
	Interval     uint32
	Duration     uint16
	MaxAdvEvts   uint8
	ChannelMask  [5]uint8
	FilterPolicy uint8
	PrimaryPhi   uint8
	SecondaryPhi uint8
	Fields       uint8 // set_id:4; scan_req_notifications:1
}

// Configure this advertisement.
func (a *Advertisement) Configure(adv_data, scan_rsp_data string, params *AdvParams) Error {
	data := AdvData{
		AdvData:     adv_data,
		ScanRspData: scan_rsp_data,
	}
	err := arm.SVCall3(svc_SD_BLE_GAP_ADV_SET_CONFIGURE, &a.handle, &data, params)
	return Error(err)
}

// Start advertisement. May only be called after it has been configured.
func (a Advertisement) Start() Error {
	err := arm.SVCall2(svc_SD_BLE_GAP_ADV_START, a.handle, 0)
	return Error(err)
}

// Stop advertisement.
func (a Advertisement) Stop() Error {
	err := arm.SVCall1(svc_SD_BLE_GAP_ADV_STOP, a.handle)
	return Error(err)
}
