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

//AaaLogin is the structure representation of the AaaLogin response
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
	UnixUserId             string   `xml:"unixUserId,attr"`
	SessionId              string   `xml:"sessionId,attr"`
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
