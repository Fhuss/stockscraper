package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/antchfx/htmlquery"
)

var (
	defaultURL      string = "https://finance.yahoo.com/quote/"
	mostactivestock string = "https://finance.yahoo.com/most-active"
	tickerxpath     string = "//*[@id='scr-res-table']/div[1]/table/tbody/tr"
	buyxpath        string = "//*[@id='quote-header-info']/div[3]/div[1]/div/span[1]"
	sellxpath       string = "//*[@id='quote-header-info']/div[3]/div[1]/p/span[1]"
)

func main() {

	tickerSymbols, err := GetTickerSymbols(mostactivestock, tickerxpath)
	if err != nil {
		fmt.Println(tickerSymbols[0])
		panic(err)
	}
	stocks := &Stocks{Stockmarket: "Most Active"}
	stocks.Entries = make([]StockEntry, len(tickerSymbols))
	for i := 0; i < len(tickerSymbols); i++ {
		url := strings.Join([]string{defaultURL, tickerSymbols[i], "?p=", tickerSymbols[i]}, "")
		stocks.Entries[i].Tickersymbol = tickerSymbols[i]
		err := stocks.Entries[i].GetSell(url, buyxpath, sellxpath)
		if err != nil {
			panic(err)
		}
	}
	res, err := MarshalAndSave(stocks)
	if err != nil {
		fmt.Println(res)
		panic(err)
	}
	fmt.Println(res)
}

func MarshalAndSave(data interface{}) (string, error) {
	date, time := GetDateAndTime()
	str, err := json.MarshalIndent(&data, "", " ")
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(date, 0777)
	if err != nil {
		return "Cannot create directory", err
	}

	f, err := os.Create(strings.Join([]string{date, "/", time, ".json"}, ""))
	if err != nil {
		return "Cannot create file", err
	}
	defer f.Close()

	f.Write(str)

	return "Successfully saved json file!", nil
}

func GetTickerSymbols(url, xpath string) ([]string, error) {
	var tickerSymbols []string

	rawhtml, err := htmlquery.LoadURL(url)
	if err != nil {
		tickerSymbols = append(tickerSymbols, "Could not find any ticker symbols")
		return tickerSymbols, err
	}

	list := htmlquery.Find(rawhtml, xpath)
	nrofTickerSymbols := len(list)
	for i := 1; i <= nrofTickerSymbols; i++ {
		ticker := strings.Join([]string{xpath, "[", strconv.Itoa(i), "]/td[1]"}, "")
		tickerVal := htmlquery.Find(rawhtml, ticker)
		tickerSymbols = append(tickerSymbols, htmlquery.InnerText(tickerVal[0]))

	}
	return tickerSymbols, nil

}

func GetDateAndTime() (string, string) {
	t := time.Now()
	return t.Format("2006-01-02"), t.Format("15:04")
}
