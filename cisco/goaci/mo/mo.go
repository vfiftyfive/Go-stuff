package mo

import (
	"encoding/xml"
	"reflect"
)

//ManagedObject is the base object
type ManagedObject interface {
}

//Reg is the factory registry map for ACI objects, mapping string to corresponding type
var Reg = map[string]reflect.Type{}

//AaaLogin is the structure representation of AaaLogin response
type AaaLogin struct {
	XMLName                xml.Name `xml:"aaaLogin"`
	Token                  string   `xml:"token,attr"`
	SiteFingerprint        string   `xml:"siteFingerPrint,attr"`
	RefreshTimeoutSeconds  string   `xml:"refreshTimeoutSeconds,attr"`
	MaximumLifetimeSeconds string   `xml:"maximumLifetimeSeconds,attr"`
	GuiIdleTimeoutSeconds  string   `xml:"guiIdleTimeoutSeconds,attr"`
	RestTimeoutSeconds     string   `xml:"restTimeoutSeconds,attr"`
	CreationTime           string   `xml:"creationTime,attr"`
	FirstLoginTime         string   `xml:"firstLoginTime,attr"`
	UserName               string   `xml:"userName,attr"`
	RemoteUser             string   `xml:"remoteUser,attr"`
	UnixUserID             string   `xml:"unixUserId,attr"`
	SessionID              string   `xml:"sessionId,attr"`
	LastName               string   `xml:"lastName,attr"`
	FirstName              string   `xml:"firstName,attr"`
	ChangePassword         string   `xml:"changePassword,attr"`
	Version                string   `xml:"version,attr"`
	BuildTime              string   `xml:"buildTime,attr"`
	Node                   string   `xml:"node,attr"`
}

func init() {
	Reg["aaaLogin"] = reflect.TypeOf((*AaaLogin)(nil)).Elem()
}

//VzEntry is the structure representation of ACI vzEntry
type VzEntry struct {
	XMLName 	xml.Name 	`xml:"vzEntry"`
	Name 		string		`xml:"name,attr"`
	DFromPort	string		`xml:"dFromPort,attr"`
	DToPort		string		`xml:"dToPort,attr"`
	EtherT		string		`xml:"etherT,attr"`
	Prot 		string		`xml:"prot,attr"`
}

func init() {
	Reg["vzEntry"] = reflect.TypeOf((*VzEntry)(nil)).Elem()
}

//VzFilter is ACI Filter object
type VzFilter struct {
	XMLName		xml.Name	`xml:"vzFilter"`
	Name 		string 		`xml:"name,attr"`
	VzEntry		VzEntry 	
}

func init() {
	Reg["vzFilter"] = reflect.TypeOf((*VzFilter)(nil)).Elem()
}

//VzRsSubjFiltAtt is ACI vzRsSubjFiltAtt object
type VzRsSubjFiltAtt struct {
	XMLName			xml.Name	`xml:"vzRsSubjFiltAtt"`
	TnVzFilterName	string 		`xml:"tnVzFilterName,attr"`
}

func init() {
	Reg["vzRsSubjFiltAtt"] = reflect.TypeOf((*VzRsSubjFiltAtt)(nil)).Elem()
}

//VzSubj is ACI Subject object
type VzSubj struct {
	XMLName				xml.Name			`xml:"vzSubj"`
	Name				string				`xml:"name,attr"`
	VzRsSubjFiltAtt		*VzRsSubjFiltAtt	`xml:",omitempty"`
}

func init() {
	Reg["vzSubj"] = reflect.TypeOf((*VzSubj)(nil)).Elem()
}

//VzBrCP is ACI contract object
type VzBrCP struct {
	XMLName	xml.Name	`xml:"vzBrCP"`
	Name	string		`xml:"name,attr"`
	VzSubj	VzSubj		
}

func init() {
	Reg["vzBrCP"] = reflect.TypeOf((*VzBrCP)(nil)).Elem()
}

//FvCtx is ACI context object
type FvCtx struct {
	XMLName		xml.Name 	`xml:"fvCtx"`
	Name 		string		`xml:"name,attr"`
}

func init() {
	Reg["fvCtx"] = reflect.TypeOf((*FvCtx)(nil)).Elem()
}

//FvRsCtx is ACI RS context object
type FvRsCtx struct {
	XMLName			xml.Name 	`xml:"fvRsCtx"`
	TnFvCtxName		string		`xml:"tnFvCtxName,attr"`
}

func init() {
	Reg["fvRsCtx"] = reflect.TypeOf((*FvRsCtx)(nil)).Elem()
}

//FvSubnet is ACI BD Subnet object
type FvSubnet struct {
	XMLName			xml.Name 	`xml:"fvSubnet"`
	IP				string		`xml:"ip,attr"`
}

func init() {
	Reg["fvSubnet"] = reflect.TypeOf((*FvSubnet)(nil)).Elem()
}

//FvBD is ACI BD object
type FvBD struct {
	XMLName		xml.Name	`xml:"fvBD"`
	Name 		string 		`xml:"name,attr"`
	FvRsCtx		FvRsCtx
	FvSubnet	[]*FvSubnet	`xml:",omitempty"`	
}

func init() {
	Reg["fvBD"] = reflect.TypeOf((*FvBD)(nil)).Elem()
}

//FvTenant is ACI Tenant object
type FvTenant struct {
	XMLName 	xml.Name 	`xml:"fvTenant"`
	Name 		string		`xml:"name,attr"`
	VzBrCP		[]*VzBrCP	`xml:",omitempty"`
	FvCtx		[]*FvCtx	`xml:",omitempty"`
	FvBD		[]*FvBD		`xml:",omitempty"`
	VzFilter	[]*VzFilter	`xml:",omitempty"`
	FvAp		[]*FvAp		`xml:",omitempty"`
}

//FvRsDomAtt is ACI RS Domain object
type FvRsDomAtt struct {
	XMLName 		xml.Name 	`xml:"fvRsDomAtt"`
	TDn				string		`xml:"tDn,attr"`
}

func init() {
	Reg["fvRsDomAtt"] = reflect.TypeOf((*FvRsDomAtt)(nil)).Elem()
}

//FvRsBd is ACI RS BD object
type FvRsBd struct {
	XMLName 		xml.Name	`xml:"fvRsBd"`
	TnFvBDName		string		`xml:"tnFvBDName,attr"`
}

func init() {
	Reg["fvRsBd"] = reflect.TypeOf((*FvRsBd)(nil)).Elem()
}

//fvRsProv is ACI provisioned Contract
type FvRsProv struct {
	XMLName 		xml.Name 	`xml:"fvRsProv"`
	TnVzBrCPName	string 		`xml:"tnVzBrCPName,attr"`
}

func init() {
	Reg["fvRsProv"] = reflect.TypeOf((*FvRsProv)(nil)).Elem()
}

//FvRsCons is ACI consumed Contract
type FvRsCons struct {
	XMLName 		xml.Name 	`xml:"fvRsCons"`
	TnVzBrCPName	string 		`xml:"tnVzBrCPName,attr"`
}

func init() {
	Reg["fvRsCons"] = reflect.TypeOf((*FvRsCons)(nil)).Elem()
}

//FvAp is ACI ANP object
type FvAp struct {
	XMLName 		xml.Name 		`xml:"fvAp"`
	Name 			string			`xml:"name,attr"`
	FvAEPg			[]*FvAEPg			`xml:",omitempty"`
}

func init() {
	Reg["fvAp"] = reflect.TypeOf((*FvAp)(nil)).Elem()
}

//FvAEPg is ACI EPG object
type FvAEPg struct {
	XMLName 		xml.Name 	`xml:"fvAEPg"`
	Name			string		`xml:"name,attr"`
	FvRsDomAtt		FvRsDomAtt		
	FvRsBd			FvRsBd		
	FvRsProv		*FvRsProv		`xml:",omitempty"`
	FvRsCons		*FvRsCons		`xml:",omitempty"`
}

func init() {
	Reg["FvAEPg"] = reflect.TypeOf((*FvAEPg)(nil)).Elem()
}
