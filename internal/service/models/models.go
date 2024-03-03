package models

// @name SearchRequest
// @description SearchRequest represents the request body for address search
type SearchRequest struct {
	Query string `json:"query"`
}

// @name GeocodeRequest
// @description GeocodeRequest represents the request body for address geocoding.
type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// @name AddressGeo
// @description AddressGeo represents the geocode result for an address.
type AddressGeo struct {
	Suggestions []struct {
		Value             string `json:"value"`
		UnrestrictedValue string `json:"unrestricted_value"`
		Data              struct {
			PostalCode string `json:"postal_code"`
			Country    string `json:"country"`
		} `json:"data"`
	} `json:"suggestions"`
}

// @name AddressSearch
// @description AddressSearch represents the search result for an address.
type AddressSearch struct {
	Source string `json:"source"`
	Result string `json:"result"`
	Metro  []struct {
		Distance float64 `json:"distance"`
		Line     string  `json:"line"`
		Name     string  `json:"name"`
	} `json:"metro"`
}
