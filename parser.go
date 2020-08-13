package cpquery

import (
	"encoding/xml"
	// "fmt"
	"strings"
)

type ITEM struct {
	XMLName        xml.Name `xml:"ITEM"`
	LL             string   `xml:"LL"`
	EV_FORMAT      string   `xml:"EV_FORMAT"`
	EVCODE         string   `xml:"EVCODE"`
	STATUS         string   `xml:"STATUS"`
	BRHNO          string   `xml:"BRHNO"`
	BRHNC          string   `xml:"BRHNC"`
	DATIME         string   `xml:"DATIME"`
	REVBRN_Z_TITLE string   `xml:"REVBRN-Z-TITLE"`
	REVBRN_Z       string   `xml:"REVBRN-Z"`
	REVBRC_Z_TITLE string   `xml:"REVBRC-Z-TITLE"`
	REVBRC_Z       string   `xml:"REVBRC-Z"`
	LISTNO_Z_TITLE string   `xml:"LISTNO-Z-TITLE"`
	LISTNO_Z       string   `xml:"LISTNO-Z"`
}

type BlueStar struct {
	XMLName  xml.Name `xml:"BlueStar"`
	TxCode   string   `xml:"TXCODE"`
	TxType   string   `xml:"TXTYPE"`
	MailNo   string   `xml:"MAILNO"`
	PBRNZON  string   `xml:"PBRNZON"`
	PageChk  string   `xml:"PAGECHK"`
	MailType string   `xml:"MAILTYPE"`
	DBName   string   `xml:"DBNAME"`
	PDATIME  string   `xml:"PDATIME"`
	COPBRNO  string   `xml:"COPBRNO"`
	EvtCode  string   `xml:"EVTCODE"`
	PL       string   `xml:"PL"`
	Items    []ITEM   `xml:"ITEM"`
}

func ParseMailQuerySingleResult(msg []byte) (*BlueStar, error) {

	ret := BlueStar{}

	err := xml.Unmarshal(msg, &ret)
	if err != nil {
		return nil, err
	}
	// fmt.Println(err, ret)
	ret.MailNo = strings.Trim(ret.MailNo, " ")
	ret.MailType = strings.Trim(ret.MailType, " ")

	for i, _ := range ret.Items {
		ret.Items[i].BRHNC = strings.Trim(ret.Items[i].BRHNC, " ")
	}
	return &ret, nil

}
