package mo

import "encoding/xml"

//ManagedObject is the base object
type ManagedObject interface {
}

//ConfigConfMo is the object use by ConfigConfMo()
type ConfigConfMo struct {
	XMLName          xml.Name   `xml:"configConfMo"`
	Dn               string     `xml:"dn,attr"`
	Cookie           string     `xml:"cookie,attr,omitempty"`
	InHierarchical   string     `xml:"inHierarchical,attr,omitempty"`
	Response         string     `xml:"response,attr,omitempty"`
	InvocationResult string     `xml:"invocationResult,attr,omitempty"`
	ErrorCode        string     `xml:"errorCode,attr,omitempty"`
	ErrorDescr       string     `xml:"errorDescr,attr,omitempty"`
	InConfig         *Inconfig  `xml:"inConfig"`
	OutConfig        *Outconfig `xml:"outConfig"`
}

//FabricVlan is the VLAN object
type FabricVlan struct {
	XMLName                  xml.Name `xml:"fabricVlan"`
	Dn                       string   `xml:"dn,omitempty,attr"`
	AssocPrimaryVlanstate    string   `xml:"assocPrimaryVlanState,omitempty,attr"`
	AssocPrimaryVlanSwitchID string   `xml:"assocPrimaryVlanSwitchId,omitempty,attr"`
	Cloud                    string   `xml:"cloud,omitempty,attr"`
	CompressionType          string   `xml:"compressionType,omitempty,attr"`
	ConfigIssues             string   `xml:"configIssues,omitempty,attr"`
	DefaultNet               string   `xml:"defaultNet,omitempty,attr"`
	EpDn                     string   `xml:"epDn,omitempty,attr"`
	FltAggr                  string   `xml:"fltAggr,omitempty,attr"`
	Global                   string   `xml:"global,omitempty,attr"`
	Id                       string   `xml:"id,attr"`
	IfRole                   string   `xml:"ifRole,omitempty,attr"`
	IfType                   string   `xml:"ifType,omitempty,attr"`
	Local                    string   `xml:"local,omitempty,attr"`
	Locale                   string   `xml:"locale,omitempty,attr"`
	McastPolicyName          string   `xml:"mcastPolicyName,omitempty,attr"`
	Name                     string   `xml:"name,omitempty,attr"`
	OperMcastPolicyName      string   `xml:"operMcastPolicyName,omitempty,attr"`
	OperState                string   `xml:"operState,omitempty,attr"`
	OverlapStateForA         string   `xml:"overlapStateForA,omitempty,attr"`
	OverlapStateForB         string   `xml:"overlagStateForB,omitempty,attr"`
	PeerDn                   string   `xml:"peerDn,omitempty,attr"`
	PolicyOwner              string   `xml:"policyOwner,omitempty,attr"`
	PubNwDn                  string   `xml:"pubNwDn,omitempty,attr"`
	PubNwName                string   `xml:"pubNwName,omitempty,attr"`
	Sharing                  string   `xml:"sharing,omitempty,attr"`
	Status                   string   `xml:"status,omitempty,attr"`
	SwitchId                 string   `xml:"switchId,omitempty,attr"`
	Transport                string   `xml:"transport,omitempty,attr"`
	Type                     string   `xml:"type,omitempty,attr"`
}

//Outconfig represents the outConfig xml element
type Outconfig struct {
	Mos []ManagedObject
}

//Inconfig represents the inConfig xml element
type Inconfig struct {
	Mos []ManagedObject
}
