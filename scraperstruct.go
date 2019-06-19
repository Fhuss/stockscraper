package main

import (
	"strconv"

	"github.com/antchfx/htmlquery"
)

type Stocks struct {
	Stockmarket string       //`json:stockmarket`
	Entries     []StockEntry //`json:"stockentries"`
}

type StockEntry struct {
	Tickersymbol string  //`json:"tickersymbol,omitempty`
	Sell         float64 //`json:"sell,omitempty"`
	//buy          float64 `json:"buy,omitempty"`

}

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
