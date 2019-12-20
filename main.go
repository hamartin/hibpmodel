package hibpmodel

/*
 * This package simply contain the HIBP JSON model.
 *
 * If you are unmarshalling the JSON file given by HIBP, you can do something
 * like this.
 *
 * import (
 *     "encoding/json"
 *
 *     "github.com/hamartin/hibpmodel"
 * )
 *
 * var hibp hibpmodel.HIBP
 * err := json.NewDecoder(<FP>).Decode(&hibp)
 * if err != nil {
 *     panic(err)
 * }
 * <DO STUFF WITH THE hibp VARIABLE>
 */

type HIBP struct {
	Breach []BreachModel `json:"BreachSearchResults,omitempty"`
	Paste  []PasteModel  `json:"PasteSearchResults,omitempty"`
}

type BreachModel struct {
	Alias      string        `json:"Alias,omitempty"`
	DomainName string        `json:"DomainName,omitempty"`
	Breaches   []IndBreaches `json:"Breaches,omitempty"`
}

type IndBreaches struct {
	AddedDate    string   `json:"AddedDate,omitempty"`
	BreachDate   string   `json:"BreachDate,omitempty"`
	DataClasses  []string `json:"DataClasses,omitempty"`
	Description  string   `json:"Description,omitempty"`
	Domain       string   `json:"Domain,omitempty"`
	IsActive     bool     `json:"IsActive,omitempty"`
	IsFabricated bool     `json:"IsFabricated,omitempty"`
	IsRetired    bool     `json:"IsRetired,omitempty"`
	IsSensitive  bool     `json:"IsSensitive,omitempty"`
	IsSpamList   bool     `json:"IsSpamList,omitempty"`
	IsVerified   bool     `json:"IsVerified,omitempty"`
	LogoFileName string   `json:"LogoFileName,omitempty"`
	ModifiedDate string   `json:"ModifiedDate,omitempty"`
	Name         string   `json:"Name,omitempty"`
	PwnCount     int      `json:"PwnCount,omitempty"`
	Title        string   `json:"Title,omitempty"`
}

type PasteModel struct {
	Alias      string          `json:"Alias,omitempty"`
	DomainName string          `json:"DomainName,omitempty"`
	Pastes     []IndPasteModel `json:"Pastes,omitempty"`
}

type IndPasteModel struct {
	Date       string `json:"Date,omitempty"`
	EmailCount int    `json:"EmailCount,omitempty"`
	Id         string `json:Id,omitempty"`
	Source     string `json:Source,omitempty"`
	Title      string `json:Title,omitempty"`
}
