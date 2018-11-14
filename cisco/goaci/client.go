package goaci

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/vfiftyfive/cisco/goaci/methods"
	"github.com/vfiftyfive/cisco/goaci/mo"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const aaaLogin = "/api/mo/aaaLogin.xml"
const aaaLogout = "/api/mo/aaaLogout.xml"

//Client abstracts connection information
type Client struct {
	http.Client
	u  string
	t  *http.Transport
	a  AaaUser
	tk string
	i  bool
	d  bool
}

//AaaUser is the structure representation of aaaUser
type AaaUser struct {
	XMLName xml.Name `xml:"aaaUser"`
	Name    string   `xml:"name,attr"`
	Pwd     string   `xml:"pwd,attr,omitempty"`
}

//Login provides user login to ACI
func (c *Client) Login(ctx context.Context) (mo.AaaLogin, error) {
	c.u = c.u + aaaLogin
	var l mo.AaaLogin
	var i interface{}
	resp, err := post(ctx, c, c.a)
	if err != nil {
		return l, err
	}
	ls, err := methods.UnmarshalXML(bytes.NewReader(resp), i)
	if err != nil {
		return l, err
	}
	if al, ok := ls[0].(*mo.AaaLogin); ok {
		c.tk = al.Token
		return *al, nil
	}
	err = fmt.Errorf("Type insertion for aaaLogin failed")
	return l, err
}

//Logout logs user out from ACI
func (c *Client) Logout(ctx context.Context) error {
	u, err := url.Parse(c.u)
	if err != nil {
		return err
	}
	u.Path = aaaLogout
	c.u = u.String()
	c.a.Pwd = ""
	_, err = post(ctx, c, c.a)
	if err != nil {
		return err
	}
	return nil
}

//SetDebug sets debug to true or false
func (c *Client) SetDebug(b bool) {
	c.d = true
}

//NewClient genereates a new Client with context
func NewClient(ctx context.Context, u string, insecure bool, user string, pwd string, debug bool) (*Client, error) {
	c := &Client {
		u: u,
		a: AaaUser {
			Name: user,
			Pwd:  pwd,
		},
		i: insecure,
		d: debug,
	}
	if t, ok := http.DefaultTransport.(*http.Transport); ok {

		c.t = &http.Transport {
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
	l, err := c.Login(ctx)
	//print aaaLogin mo is debug is true
	if c.d {
		spew.Dump(l)
	}
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
	//Add context to the request
	req = req.WithContext(ctx)
	//Add Authentication token as cookie
	ck := http.Cookie{
		Name:  "APIC-cookie",
		Value: c.tk,
	}
	req.AddCookie(&ck)
	fmt.Printf("Posting to URL: %s\n", c.u)
	err = methods.XMLPrint(xmlByte)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	//Send response to Stdout if debug is true
	if c.d {
		fmt.Println("Debug Mode - raw HTTP response:")
		methods.XMLPrint(rbody)
	}
	if err != nil {
		return nil, err
	}
	return rbody, nil
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
	return u.String(), err
}
