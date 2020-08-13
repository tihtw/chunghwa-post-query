package cpquery

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	ConnectionHost string
}

func NewClient() *Client {
	return &Client{}

}

type Body []byte

func (c *Client) QueryMailDetail(mailNo string) (*BlueStar, error) {

	requestBody := `<?xml version="1.0" encoding="UTF-8"?>
  <soap:Envelope
    xmlns:soap="http://www.w3.org/2003/05/soap-envelope"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xmlns:xsd="http://www.w3.org/2001/XMLSchema"
    xmlns:tns="http://tempuri.org/"
     xmlns:tn="http://tempuri.org/">
      <soap:Body>
        <tn:MailQuerySingle>
          <tns:MAIL_NO>` + mailNo + `</tns:MAIL_NO>
        </tn:MailQuerySingle>
      </soap:Body>
    </soap:Envelope>`

	host := "sprws.post.gov.tw"

	// fmt.Println("request body:", requestBody)
	data, err := c.doRquest(host, requestBody)
	if err != nil {
		return nil, err
	}

	type Envelope struct {
		Envelope Body `xml:"Body>MailQuerySingleResponse>MailQuerySingleResult"`
	}

	env := Envelope{}
	if err := xml.Unmarshal(data, &env); err != nil {
		fmt.Println(err)
		return nil, nil
	}

	d, err := ParseMailQuerySingleResult(env.Envelope)

	return d, err

}

func (c *Client) doRquest(host string, requestBody string) ([]byte, error) {

	body := strings.NewReader(requestBody)

	hostname := host

	if c.ConnectionHost != "" {
		hostname = c.ConnectionHost
	}

	req, err := http.NewRequest("POST", "http://"+hostname+"/PSTTP_MailQuery.asmx", body)
	if err != nil {
		// handle err
		return nil, fmt.Errorf("new request: %q", err)
	}

	req.Header.Set("Content-Type", "application/soap+xml")
	req.Host = "sprws.post.gov.tw"

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client connection: %q", err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, err

}
