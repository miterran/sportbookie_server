package pickmon

import (
	"net/http"
	"encoding/xml"
)

// FetchGames ...
func FetchGames(url string) (Lines, error) {
	var lines Lines
	response, err := http.Get(url)
	if err != nil {
		return lines, err
	}
	decoder := xml.NewDecoder(response.Body)
	decoder.CharsetReader = MakeCharsetReader
	err = decoder.Decode(&lines)
	if err != nil {
		return lines, err
	}
	return lines, nil
}