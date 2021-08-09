package models

// Schema is the data structure we need to parse the JSON server responses correctly. Furthermore we need some fields & their values to control Switch / If Else statements.
type Schema struct {
	Status string `json:"status"`
	Data   struct {
		Team struct {
			ID         int    `json:"id"`
			Name       string `json:"name"`
			IsNational bool   `json:"isNational"`
			Players    []struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Age       string `json:"age"`
			} `json:"players"`
		} `json:"team"`
	} `json:"data"`
}

// The Playerdata struct defines the data structure of a football player. The data structure is derived from the task.
type PlayerData struct {
	TeamId             int
	PlayerId           string
	FirstName          string
	LastName           string
	Age                string
	NameOfFootballClub string
	NameOfNationalteam string
	PlaceholderName    string
	IsNational         bool
}

//JsonSchema represents the Go type definition of the original json response, which is send back by server API.
//In other words it's a conversion from JSON response body schema to GO struct.
type jsonSchema struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   struct {
		Team struct {
			ID          int    `json:"id"`
			OptaID      int    `json:"optaId"`
			Country     string `json:"country"`
			CountryName string `json:"countryName"`
			Name        string `json:"name"`
			LogoUrls    []struct {
				Size string `json:"size"`
				URL  string `json:"url"`
			} `json:"logoUrls"`
			IsNational      bool `json:"isNational"`
			HasOfficialPage bool `json:"hasOfficialPage"`
			Competitions    []struct {
				CompetitionID   int    `json:"competitionId"`
				CompetitionName string `json:"competitionName"`
			} `json:"competitions"`
			Players []struct {
				ID           string `json:"id"`
				Country      string `json:"country"`
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				Name         string `json:"name"`
				Position     string `json:"position"`
				Number       int    `json:"number"`
				BirthDate    string `json:"birthDate"`
				Age          string `json:"age"`
				Height       int    `json:"height"`
				Weight       int    `json:"weight"`
				ThumbnailSrc string `json:"thumbnailSrc"`
				Affiliation  struct {
					Name         string `json:"name"`
					ThumbnailSrc string `json:"thumbnailSrc"`
				} `json:"affiliation"`
			} `json:"players"`
			Officials []struct {
				CountryName  string `json:"countryName"`
				ID           string `json:"id"`
				FirstName    string `json:"firstName"`
				LastName     string `json:"lastName"`
				Country      string `json:"country"`
				Position     string `json:"position"`
				ThumbnailSrc string `json:"thumbnailSrc"`
				Affiliation  struct {
					Name         string `json:"name"`
					ThumbnailSrc string `json:"thumbnailSrc"`
				} `json:"affiliation"`
			} `json:"officials"`
			Colors struct {
				ShirtColorHome string `json:"shirtColorHome"`
				ShirtColorAway string `json:"shirtColorAway"`
				CrestMainColor string `json:"crestMainColor"`
				MainColor      string `json:"mainColor"`
			} `json:"colors"`
		} `json:"team"`
	} `json:"data"`
	Message string `json:"message"`
}
