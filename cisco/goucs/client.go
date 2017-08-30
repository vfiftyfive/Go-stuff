package goucs

import (
	"context"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

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

//AaaLogout is the structure representation of the xsd object
type AaaLogout struct {
	XMLName xml.Name `xml:"aaaLogout"`
	Cookie  string   `xml:"inCookie,attr"`
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

//Login provides user login to UCS
func (c *Client) Login(ctx context.Context) error {

	res, err := post(ctx, c, c.a)
	if err != nil {
		return err
	}
	err = xml.Unmarshal(res, &c.a)
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
	xmlPrint(c.u, xmlByte)
	resp, err := c.Do(req)
	defer resp.Body.Close()
	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return rbody, nil
}

func xmlPrint(s string, b []byte) {

	fmt.Printf("Posting to URL: %s\n", s)
	os.Stdout.Write(b)
	fmt.Printf("\n")
}
