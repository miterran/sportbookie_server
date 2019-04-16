package pickmon

import (
	"encoding/xml"
	"io"
	"strings"
	"time"

	"golang.org/x/text/encoding/charmap"
)

// Score ...
type Score struct {
	Home   float64 `xml:"team1"`
	Away   float64 `xml:"team2"`
	Winner string  `xml:"winner"`
}

// Team ...
type Team struct {
	Rot  string         `xml:"rotnum"`
	Name CustomTeamName `xml:"name"`
}

// Money ...
type Money struct {
	Home float64 `xml:"team1"`
	Away float64 `xml:"team2"`
	Draw float64 `xml:"draw"`
}

// Spread ...
type Spread struct {
	Points float64 `xml:"points"`
	Home   float64 `xml:"team1"`
	Away   float64 `xml:"team2"`
}

// Total ...
type Total struct {
	Points float64 `xml:"points"`
	Over   float64 `xml:"over"`
	Under  float64 `xml:"under"`
}

// Line ...
type Line struct {
	Periodnum  int        `xml:"periodnum"`
	Period     string     `xml:"perioddesc"`
	CutOffTime CustomTime `xml:"wagercutoff"`
	Money      Money      `xml:"money"`
	Spread     Spread     `xml:"spread"`
	Total      Total      `xml:"total"`
	Score      Score      `xml:"score"`
}

// Game ...
type Game struct {
	ID        string       `xml:"id"`
	Header    string       `xml:"header"`
	Sport     CustomSport  `xml:"sporttype"`
	League    CustomLeague `xml:"sportsubtype"`
	Void      int          `xml:"void"`
	MatchTime CustomTime   `xml:"gamedate"`
	HomeTeam  Team         `xml:"team1"`
	AwayTeam  Team         `xml:"team2"`
	Line      Line         `xml:"line"`
}

// Lines ...
type Lines struct {
	Games []Game `xml:"game"`
}

// CustomTime ...
type CustomTime struct {
	Parse time.Time
}

// UnmarshalXML ...
func (ct *CustomTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	t = t.Add(time.Hour * 4)
	*ct = CustomTime{t}
	return nil
}

// SportType ...
type SportType map[string]string

// SportTypeConvert ...
var SportTypeConvert = SportType{
	"Basketball": "Basketball",
	"Football":   "Football",
	"Baseball":   "Baseball",
	"Soccer":     "Soccer",
	"Hockey":     "Hockey",
	"Fighting":   "Fighting",
}

// CustomSport ...
type CustomSport struct {
	Parse string
}

// UnmarshalXML ...
func (cst *CustomSport) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	v, ok := SportTypeConvert[s]
	if ok {
		*cst = CustomSport{v}
	}
	return nil
}

// LeagueType ...
type LeagueType map[string]string

// LeagueTypeConvert ...
var LeagueTypeConvert = LeagueType{
	"NBA":          "NBA",
	"NFL":          "NFL",
	"MLB":          "MLB",
	"NHL":          "NHL",
	"European Cup": "European Cup",
	// "Boxing":        "Boxing",
	"UFC": "UFC",
	// "International": "International",
}

// CustomLeague ...
type CustomLeague struct {
	Parse string
}

// UnmarshalXML ...
func (cst *CustomLeague) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	err := d.DecodeElement(&s, &start)
	if err != nil {
		return err
	}
	v, ok := LeagueTypeConvert[s]
	if ok {
		*cst = CustomLeague{v}
	}
	return nil
}

// MakeCharsetReader ...
func MakeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	return charmap.Windows1252.NewDecoder().Reader(input), nil
}

// CustomTeamName ...
type CustomTeamName struct {
	Parse string
}

// UnmarshalXML ...
func (cst *CustomTeamName) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var name string
	err := d.DecodeElement(&name, &start)
	if err != nil {
		return err
	}
	i := strings.Index(name, "(")
	if i >= 0 {
		b := []byte(name)
		name = string(b[:i])
		name = strings.TrimSpace(name)
	}
	*cst = CustomTeamName{name}
	return nil
}
