package methods

import (
	"encoding/xml"
	"fmt"
	"github.com/vfiftyfive/cisco/goaci/mo"
	"io"
	"os"
	"reflect"
)

//XMLPrint write bytes to Stdout
func XMLPrint(b []byte) error {

	_, err := os.Stdout.Write(b)
	if err != nil {
		return err
	}
	fmt.Printf("\n")
	return nil
}

//UnmarshalXML wraps xml.DecodeElement() to find the corresponding type
//in the registry map
func UnmarshalXML(r io.Reader, i interface{}) ([]mo.ManagedObject, error) {

	i = nil
	var mos []mo.ManagedObject
	d := xml.NewDecoder(r)
	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if s, ok := t.(xml.StartElement); ok {
			for n, m := range mo.Reg {
				if s.Name.Local == n {
					i = reflect.New(m).Interface()
					break
				}
			}
			if i != nil {
				d.DecodeElement(&i, &s)
				mos = append(mos, i)
				i = nil
			}
		}
	}
	return mos, nil
}
