package goucs

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vfiftyfive/cisco/goucs/methods"
	"github.com/vfiftyfive/cisco/goucs/mo"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

var debug = false

//Client abstracts connection information
type Client struct {
	http.Client
	u string
	t *http.Transport
	a AaaLogin
	i bool
}

//AaaLogin is the structure representation for the xsd object
type AaaLogin struct {
	XMLName  xml.Name `xml:"aaaLogin"`
	User     string   `xml:"inName,attr"`
	Password string   `xml:"inPassword,attr"`
	Cookie   string   `xml:"outCookie,attr,omitempty"`
}

//AaaLogout is the structure representation for the xsd object
type AaaLogout struct {
	XMLName xml.Name `xml:"aaaLogout"`
	Cookie  string   `xml:"inCookie,attr"`
}

//Login provides user login to UCS
func (c *Client) Login(ctx context.Context) error {

	resp, err := post(ctx, c, c.a)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(resp, &c.a)
	if err != nil {
		return err
	}
	return nil
}

//Logout logs user out from UCS
func (c *Client) Logout(ctx context.Context) error {

	a := AaaLogout{
		Cookie: c.a.Cookie,
	}
	_, err := post(ctx, c, a)
	if err != nil {
		return err
	}

	return nil
}

//ConfigConfMo is the method to configure a single object
func (c *Client) ConfigConfMo(ctx context.Context, dn string, m mo.ManagedObject) (mo.ConfigConfMo, error) {

	cm := mo.ConfigConfMo{
		Dn:             dn,
		Cookie:         c.a.Cookie,
		InHierarchical: "false",
		InConfig:       &m,
	}
	resp, err := post(ctx, c, cm)
	if err != nil {
		return cm, err
	}

	err = xml.Unmarshal(resp, &cm)
	if err != nil {
		return cm, err
	}

	if debug {
		fmt.Println("Debug Mode - HTTP response body:")
		spew.Dump(cm)
	}
	if "" != cm.ErrorCode {
		err := fmt.Errorf(cm.ErrorDescr)
		return cm, err
	}

	return cm, nil
}

//ConfigConfMos is the method to configure multiple objects simultaneously
func (c *Client) ConfigConfMos(ctx context.Context, pairs []mo.Pair) (mo.ConfigConfMos, error) {

	cms := mo.ConfigConfMos{
		Cookie:    c.a.Cookie,
		InConfigs: &mo.InConfigs{},
	}
	for _, p := range pairs {
		cms.InConfigs.Pairs = append(cms.InConfigs.Pairs, p)
	}
	resp, err := post(ctx, c, cms)
	if err != nil {
		return cms, err
	}

	mos, err := methods.UnmarshalXML(bytes.NewReader(resp), &cms)
	cms.OutConfigs = &mos
	if err != nil {
		return cms, err
	}

	if debug {
		fmt.Println("Debug Mode - HTTP response body:")
		spew.Dump(cms)
	}
	if "" != cms.ErrorCode {
		err := fmt.Errorf(cms.ErrorDescr)
		return cms, err
	}

	return cms, nil

}

//ConfigResolveChildren is the method to retrieve multiple objects children information
func (c *Client) ConfigResolveChildren(ctx context.Context, cid string, inDn string) (mo.ConfigResolveChildren, error) {

	crc := mo.ConfigResolveChildren{
		Cookie:         c.a.Cookie,
		ClassId:        cid,
		InDn:           inDn,
		InHierarchical: "true",
	}

	resp, err := post(ctx, c, crc)
	if err != nil {
		return crc, err
	}
	var tm mo.ManagedObject
	mos, err := methods.UnmarshalXML(bytes.NewReader(resp), &tm)
	if err != nil {
		return crc, err
	}
	crc.OutConfigs = &mos

	if debug {
		fmt.Println("Debug Mode - HTTP response spew dump:")
		spew.Dump(crc)
	}
	if "" != crc.ErrorCode {
		err := fmt.Errorf(crc.ErrorDescr)
		return crc, err
	}

	return crc, nil
}

var scheme = regexp.MustCompile(`^\w+://`)

//ParseURL parses URL information to include uri and default settings
func ParseURL(s string) (string, error) {

	if s != "" {
		// Default to https
		if !scheme.MatchString(s) {
			s = "https://" + s
		}
	}

	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}

	// Default the path to /nuova
	u.Path = "/nuova"
	us := u.String()

	return us, err

}

//NewClient genereates a new Client with context
func NewClient(ctx context.Context, u string, insecure bool, user string, pwd string) (*Client, error) {

	c := &Client{
		u: u,
		a: AaaLogin{
			User:     user,
			Password: pwd,
		},
		i: insecure,
	}
	if t, ok := http.DefaultTransport.(*http.Transport); ok {

		c.t = &http.Transport{
			Proxy:                 t.Proxy,
			DialContext:           t.DialContext,
			MaxIdleConns:          t.MaxIdleConns,
			IdleConnTimeout:       t.IdleConnTimeout,
			TLSHandshakeTimeout:   t.TLSHandshakeTimeout,
			ExpectContinueTimeout: t.ExpectContinueTimeout,
		}
	}
	if !c.i {
		err := fmt.Errorf("InsecureSkipVerify must be set to true")
		return nil, err
	}
	c.t.TLSClientConfig = &tls.Config{InsecureSkipVerify: c.i}
	c.Client.Transport = c.t
	err := c.Login(ctx)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func post(ctx context.Context, c *Client, xmlStruct interface{}) ([]byte, error) {

	xmlByte, err := xml.MarshalIndent(&xmlStruct, " ", " ")
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.u, strings.NewReader(string(xmlByte)))
	if err != nil {
		return nil, err
	}
	req.Header.Set(`Content-Type`, `application/x-www-form-urlencoded`)
	req = req.WithContext(ctx)
	fmt.Printf("Posting to URL: %s\n", c.u)
	err = methods.XMLPrint(xmlByte)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)

	//Send response to Stdout if debug is true
	if debug {
		fmt.Println("Debug Mode - raw HTTP response:")
		methods.XMLPrint(rbody)
	}

	if err != nil {
		return nil, err
	}

	return rbody, nil
}
