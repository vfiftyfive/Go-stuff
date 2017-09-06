package mo

import (
	"encoding/xml"
	"reflect"
)

//ManagedObject is the base object
type ManagedObject interface{}

//Reg is the factory registry map for UCS objects, mapping string to corresponding type
var Reg = map[string]reflect.Type{}

//ConfigConfMo is the input struct used by ConfigConfMo()
type ConfigConfMo struct {
	XMLName          xml.Name         `xml:"configConfMo"`
	Dn               string           `xml:"dn,attr"`
	Cookie           string           `xml:"cookie,attr,omitempty"`
	InHierarchical   string           `xml:"inHierarchical,attr,omitempty"`
	Response         string           `xml:"response,attr,omitempty"`
	InvocationResult string           `xml:"invocationResult,attr,omitempty"`
	ErrorCode        string           `xml:"errorCode,attr,omitempty"`
	ErrorDescr       string           `xml:"errorDescr,attr,omitempty"`
	InConfig         *ManagedObject   `xml:"inConfig"`
	OutConfig        *[]ManagedObject `xml:"outConfig"`
}

//ConfigConfMos is the input struct used by ConfigConfMos()
type ConfigConfMos struct {
	XMLName          xml.Name         `xml:"configConfMos"`
	Cookie           string           `xml:"cookie,attr,omitempty"`
	Response         string           `xml:"response,attr,omitempty"`
	InvocationResult string           `xml:"invocationResult,attr,omitempty"`
	ErrorCode        string           `xml:"errorCode,attr,omitempty"`
	ErrorDescr       string           `xml:"errorDescr,attr,omitempty"`
	InConfigs        *InConfigs       `xml:"inConfigs"`
	OutConfigs       *[]ManagedObject `xml:"outConfigs"`
}

//ConfigResolveChildren is the input struct used by ConfigResolveChildren()
type ConfigResolveChildren struct {
	XMLName          xml.Name         `xml:"configResolveChildren"`
	ClassId          string           `xml:"classId,attr,omitempty"`
	Cookie           string           `xml:"cookie,attr,omitempty"`
	InDn             string           `xml:"inDn,attr,omitempty"`
	InHierarchical   string           `xml:"inHierarchical,attr,omitempty"`
	Response         string           `xml:"response,attr,omitempty"`
	InvocationResult string           `xml:"invocationResult,attr,omitempty"`
	ErrorCode        string           `xml:"errorCode,attr,omitempty"`
	ErrorDescr       string           `xml:"errorDescr,attr,omitempty"`
	OutConfigs       *[]ManagedObject `xml:"outConfigs"`
}

//FabricVlan represents the fabricVlan xml element
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
	PubNwId                  string   `xml:"pubNwId,omitempty,attr"`
	PubNwName                string   `xml:"pubNwName,omitempty,attr"`
	Sharing                  string   `xml:"sharing,omitempty,attr"`
	Status                   string   `xml:"status,omitempty,attr"`
	SwitchId                 string   `xml:"switchId,omitempty,attr"`
	Transport                string   `xml:"transport,omitempty,attr"`
	Type                     string   `xml:"type,omitempty,attr"`
}

func init() {
	Reg["fabricVlan"] = reflect.TypeOf((*FabricVlan)(nil)).Elem()
}

//VnicEther represents the vnicEther xml element
type VnicEther struct {
	XMLName                  xml.Name     `xml:"vnicEther"`
	AdaptorProfileName       string       `xml:"adaptorProfileName,attr,omitempty"`
	Addr                     string       `xml:"addr,attr,omitempty"`
	AdminCdnName             string       `xml:"adminCdnName,attr, omitempty"`
	AdminHostPort            string       `xml:"adminHostPort,attr,omitempty"`
	AdminVcon                string       `xml:"adminVcon,attr,omitempty"`
	BootDev                  string       `xml:"bootDev,attr,omitempty"`
	Children                 *VnicEtherIf `xml:",omitempty"`
	CdnPropInSync            string       `xml:"cdnPropInSync,attr,omitempty"`
	CdnSource                string       `xml:"cdnSource,attr,omitempty"`
	ChildAction              string       `xml:"childAction,attr,omitempty"`
	ConfigQualifier          string       `xml:"configQualifier,attr,omitempty"`
	ConfigState              string       `xml:"configState,attr,omitempty"`
	Dn                       string       `xml:"dn,attr,omitempty"`
	DynamicId                string       `xml:"dynamicId,attr,omitempty"`
	EquipmentDn              string       `xml:"equipmentDn,attr,omitempty"`
	FltAggr                  string       `xml:"fltAggr,attr,omitempty"`
	IdentPoolName            string       `xml:"identPoolName,attr,omitempty"`
	InstType                 string       `xml:"instType,attr,omitempty"`
	Mtu                      string       `xml:"mtu,attr,omitempty"`
	Name                     string       `xml:"name,attr,omitempty"`
	NwCtrlPolicyName         string       `xml:"nwCtrlPolicyName,attr,omitempty"`
	NwTemplName              string       `xml:"nwTemplName,attr,omitempty"`
	OperAdaptorProfileName   string       `xml:"operAdaptorProfileName,attr,omitempty"`
	OperCdnName              string       `xml:"operCdnName,attr,omitempty"`
	OperHostPort             string       `xml:"operHostPort,attr,omitempty"`
	OperIdentPoolName        string       `xml:"operIdentPoolName,attr,omitempty"`
	OperNwCtrlPolicyName     string       `xml:"operNwCtrlPolicyName,attr,omitempty"`
	OperNwTemplName          string       `xml:"operNwTemplName,attr,omitempty"`
	OperOrder                string       `xml:"operOrder,attr,omitempty"`
	OperPinToGroupName       string       `xml:"operPinToGroupName,attr,omitempty"`
	OperQosPolicyName        string       `xml:"operQosPolicyName,attr,omitempty"`
	OperSpeed                string       `xml:"operSpeed,attr,omitempty"`
	OperStatsPolicyName      string       `xml:"operStatsPolicyName,attr,omitempty"`
	OperVcon                 string       `xml:"operVcon,attr,omitempty"`
	Order                    string       `xml:"order,attr,omitempty"`
	Owner                    string       `xml:"owner,attr,omitempty"`
	PfDn                     string       `xml:"pfDn,attr,omitempty"`
	PinToGroupName           string       `xml:"pinToGroupName,attr,omitempty"`
	PropAcl                  string       `xml:"propAcl,attr,omitempty"`
	Purpose                  string       `xml:"purpose,attr,omitempty"`
	QosPolicyName            string       `xml:"qosPolicyName,attr,omitempty"`
	StatsPolicyName          string       `xml:"statsPolicyName,attr,omitempty"`
	SwitchId                 string       `xml:"switchId,attr,omitempty"`
	Type                     string       `xml:"type,attr,omitempty"`
	VirtualizationPreference string       `xml:"virtualizationPreference,attr,omitempty"`
}

func init() {
	Reg["vnicEther"] = reflect.TypeOf((*VnicEther)(nil)).Elem()
}

//VnicEtherIf represents the vnicEtherIf xml element
type VnicEtherIf struct {
	XMLName             xml.Name `xml:"vnicEtherIf"`
	Addr                string   `xml:"addr,attr,omitempty"`
	ChildAction         string   `xml:"childAction,attr,omitempty"`
	ConfigQualifier     string   `xml:"configQualifier,attr,omitempty"`
	DefaultNet          string   `xml:"defaultNet,attr,omitempty"`
	FltAggr             string   `xml:"fltAggr,attr,omitempty"`
	Name                string   `xml:"name,attr,omitempty"`
	OperPrimaryVnetD    string   `xml:"operPrimaryVnetDn,attr,omitempty"`
	OperPrimaryVnetName string   `xml:"operPrimaryVnetName,attr,omitempty"`
	OperState           string   `xml:"operState,attr,omitempty"`
	OperVnetDn          string   `xml:"operVnetDn,attr,omitempty"`
	OperVnetName        string   `xml:"operVnetName,attr,omitempty"`
	Owner               string   `xml:"owner,attr,omitempty"`
	PropAcl             string   `xml:"propAcl,attr,omitempty"`
	PubNwId             string   `xml:"pubNwId,attr,omitempty"`
	Rn                  string   `xml:"rn,attr,omitempty"`
	Sharing             string   `xml:"sharing,attr,omitempty"`
	SwitchId            string   `xml:"switchId,attr,omitempty"`
	Type                string   `xml:"type,attr,omitempty"`
	Vnet                string   `xml:"vnet,attr,omitempty"`
}

func init() {
	Reg["vnicEtherIf"] = reflect.TypeOf((*VnicEtherIf)(nil)).Elem()
}

//InConfigs represents the InConfigs xml element
type InConfigs struct {
	Pairs []Pair
}

//Pair represents the pair xml element
type Pair struct {
	XMLName xml.Name `xml:"pair"`
	Key     string   `xml:"key,attr"`
	Mo      ManagedObject
}
