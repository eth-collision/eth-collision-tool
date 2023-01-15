package tool

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"math/big"
)

func FormatBigInt(n big.Int) string {
	i := n.Int64()
	p := message.NewPrinter(language.English)
	return p.Sprintf("%d", i)
}
