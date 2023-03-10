package main

import (
	"encoding/hex"
	"fmt"
	"strings"
)

const (
	proactiveCommand = "D0" // 37 deve calcular o tamanho
	commandDetails   = "8103012400"
	deviceIdentities = "82028182"
	title            = "05"
	item             = "8F"
)

func main() {

	titleHex := convertText2ApduFormat("Para ativar sua linha movel, selecione:", false, 0)
	itemHex1 := convertText2ApduFormat("DDD 11", true, 1)
	itemHex2 := convertText2ApduFormat("Outro DDD", true, 2)
	itemHex3 := convertText2ApduFormat("Sair", true, 3)

	//fmt.Println("====== SIZE APDU ======")
	apduSizeDec := len(commandDetails+deviceIdentities+title+titleHex+itemHex1+itemHex2+itemHex3) / 2
	apduSizeHex := strings.ToUpper(fmt.Sprintf("%02x", apduSizeDec))
	//fmt.Println(apduSizeHex)

	fmt.Println("Proactive Command: " + proactiveCommand + apduSizeHex)
	fmt.Println("Command Details: " + commandDetails)
	fmt.Println("Device Identities: " + deviceIdentities)
	fmt.Println("Title: " + title + " " + titleHex)
	fmt.Println("Item: " + itemHex1)
	fmt.Println("Item: " + itemHex2)
	fmt.Println("Item: " + itemHex3)

	fmt.Println("====== FULL APDU ======")
	fmt.Println(proactiveCommand + apduSizeHex + commandDetails + deviceIdentities + title + titleHex + itemHex1 + itemHex2 + itemHex3)
	//convertText2ApduFormat("Banca Movil")
	//convertText2ApduFormat("Registro Banco")
	//convertText2ApduFormat("Ultimos Movs.")

}

func convertText2ApduFormat(text string, isItem bool, pos int) string {

	hexStr := hex.EncodeToString([]byte(text))
	hexLen := len(hexStr) / 2

	if isItem {
		hexPos := strings.ToUpper(fmt.Sprintf("%02x", pos))
		hexStr = hexPos + hexStr
		hexLen = len(hexStr) / 2
		return item + strings.ToUpper(fmt.Sprintf("%02x", hexLen)) + strings.ToUpper(hexStr)
	}
	//fmt.Println(hexStr)
	//fmt.Println("====== DYNAMIC TEXT ======")
	//fmt.Println(strings.ToUpper(hexSize))
	//fmt.Println(strings.ToUpper(hexStr))
	return strings.ToUpper(fmt.Sprintf("%02x", hexLen)) + strings.ToUpper(hexStr)
}
