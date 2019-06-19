package main

import (
	"strconv"

	"github.com/antchfx/htmlquery"
)

/*
Stocks is a struc containing an array of stock entries as well a stock market name.
*/
type Stocks struct {
	Stockmarket string       //`json:stockmarket`
	Entries     []StockEntry //`json:"stockentries"`
}

/*
StockEntry is a struct containing a single entry
*/
type StockEntry struct {
	Tickersymbol string  //`json:"tickersymbol,omitempty`
	Sell         float64 //`json:"sell,omitempty"`
	//buy          float64 `json:"buy,omitempty"`

}

/*
GetSell retrieves the current Sell value of a particular stock and assigns it to the correct struct.
Returns an error message when called
*/
func (s *StockEntry) GetSell(url, buyxpath, sellxpath string) error {

	rawhtml, err := htmlquery.LoadURL(url)
	if err != nil {
		return err
	}
	/*
		buyidx := htmlquery.Find(rawhtml, buyxpath)
		buy, err := strconv.ParseFloat(htmlquery.InnerText(buyidx[1]), 64)
		//buy := htmlquery.InnerText(buyidx[0])
		fmt.Println(buy)
		if err != nil {
			return err
		}
		s.buy = buy
	*/
	sellidx := htmlquery.Find(rawhtml, sellxpath)
	sell, err := strconv.ParseFloat(htmlquery.InnerText(sellidx[0]), 64)
	if err != nil {
		return err
	}
	s.Sell = sell

	return nil

}
