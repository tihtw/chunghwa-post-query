package main

import (
	"fmt"
	"github.com/tihtw/chunghwa-post-query"
)

func main() {
	msg := `<BlueStar
	MsgName="LIQD"
	RqUid="df4b4a0d-a856-4aad-8f02-f09ebe146775"
	xmlns="http://www.cedar.com.tw/bluestar/" Status="0">
	<TXCODE> </TXCODE>
	<TXTYPE>13</TXTYPE>
	<MAILNO>44261122507318      </MAILNO>
	<PBRNZON>      </PBRNZON>
	<PAGECHK>0</PAGECHK>
	<MAILTYPE>普掛大宗郵付　　　　</MAILTYPE>
	<DBNAME>   </DBNAME>
	<PDATIME></PDATIME>
	<COPBRNO>      </COPBRNO>
	<EVTCODE>  </EVTCODE>
	<PL>00558</PL>
	<ITEM>
		<LL>184</LL>
		<EV_FORMAT>Z</EV_FORMAT>
		<EVCODE>Z4</EVCODE>
		<STATUS>運輸途中　　　　　　　　　　　</STATUS>
		<BRHNO>091813</BRHNO>
		<BRHNC>台北中心掛號函件股　</BRHNC>
		<DATIME>20200714232138</DATIME>
		<REVBRN-Z-TITLE>接收局號　　　　    </REVBRN-Z-TITLE>
		<REVBRN-Z>241602</REVBRN-Z>
		<REVBRC-Z-TITLE>接收局名　　　　    </REVBRC-Z-TITLE>
		<REVBRC-Z>蘆洲郵局郵務股　　　</REVBRC-Z>
		<LISTNO-Z-TITLE>清單號碼　　　　    </LISTNO-Z-TITLE>
		<LISTNO-Z>0355       </LISTNO-Z>
	</ITEM>
	<ITEM>
		<LL>369</LL>
		<EV_FORMAT>A</EV_FORMAT><EVCODE>A1</EVCODE><STATUS>交寄郵件
		</STATUS><BRHNO>220073</BRHNO><BRHNC>中和中山路郵局　　　</BRHNC><DATIME>20200714115000</DATIME><PVAAMT-A-TITLE>報值或保價金額      </PVAAMT-A-TITLE>
		<PVAAMT-A>000000</PVAAMT-A>
		<CODAMT-A-TITLE>代收貨價金額　      </CODAMT-A-TITLE><CODAMT-A>      </CODAMT-A><POSWGT-TITLE>重量（公克）　      </POSWGT-TITLE><POSWGT>000012</POSWGT>
		<POSFEE-TITLE>郵資　　　　　      </POSFEE-TITLE><POSFEE>000028</POSFEE><SICNO-TITLE>特約戶編號　　      </SICNO-TITLE><SICNO>             </SICNO><OPBZON-TITLE>寄達局郵遞區號      </OPBZON-TITLE>
		<OPBZON>0    </OPBZON><OPBZOC-TITLE>寄達局郵遞區名      </OPBZOC-TITLE><OPBZOC>　　　　　　</OPBZOC><ORSTEL-TITLE>寄件人電話　　      </ORSTEL-TITLE><ORSTEL>                  </ORSTEL><TASTEL-TITLE>收件人電話　　      </TASTEL-TITLE><TASTEL>                  </TASTEL></ITEM></BlueStar>`

	result, err := cpquery.ParseMailQuerySingleResult([]byte(msg))
	fmt.Println(err, result, len(result.Items))

}
