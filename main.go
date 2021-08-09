package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

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

type schema struct {
	Status string `json:"status"`
	Data   struct {
		Team struct {
			Name        string `json:"name"`
			CountryName string `json:"countryName"`
			IsNational  bool   `json:"isNational"`
			Players     []struct {
				Name      string `json:"name"`
				FirstName string `json:"firstName"`
				LastName  string `json:"lastName"`
				Age       string `json:"age"`
			} `json:"players"`
		} `json:"team"`
	} `json:"data"`
}

type playerData struct {
	FirstName    string
	LastName     string
	Age          string
	Team         string
	Country      string
	NationalTeam bool
}

const (
	urlStart = "https://api-origin.onefootball.com/score-one-proxy/api/teams/en/"
	urlEnd   = ".json"
	maxTeams = 6
)

func main() {

	var tmp schema
	var listOfAllPlayer []playerData
	var onePlayer playerData
	var counterUrlID int
	var teamList []string
	// var teamName string
	var finalURL string

	// // Data Part
	// jsonByteData := []byte(`
	//  {"status":"ok","code":0,"data":{"team":{"id":1,"optaId":0,"country":"CY","countryName":"Germany","name":"Apoel FC","logoUrls":[{"size":"56x56","url":"https://images.onefootball.com/icons/teams/56/1.png"},{"size":"164x164","url":"https://images.onefootball.com/icons/teams/164/1.png"}],"isNational":false,"hasOfficialPage":false,"competitions":[{"competitionId":140,"competitionName":"Club Friendly Games"}],"players":[{"id":"62290","country":"Georgia","firstName":"Giorgi","lastName":"Kvilitaia","name":"Giorgi Kvilitaia","position":"Forward","number":0,"birthDate":"1993-10-01","age":"27","height":193,"weight":74,"thumbnailSrc":"https://images.onefootball.com/players/180/62290.jpg","affiliation":{"name":"Georgia","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ge.png"}},{"id":"35165","country":"Norway","firstName":"Ghayas","lastName":"Zahid","name":"Ghayas Zahid","position":"Midfielder","number":17,"birthDate":"1994-09-08","age":"26","height":174,"weight":66,"thumbnailSrc":"https://images.onefootball.com/players/180/35165.jpg","affiliation":{"name":"Norway","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/no.png"}},{"id":"133462","country":"Ireland","firstName":"Jack","lastName":"Byrne","name":"Jack Byrne","position":"Midfielder","number":29,"birthDate":"1996-04-24","age":"25","height":176,"weight":73,"thumbnailSrc":"https://images.onefootball.com/players/180/133462.jpg","affiliation":{"name":"Ireland","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ie.png"}},{"id":"431250","country":"Cyprus","firstName":"Stavros","lastName":"Georgiou","name":"Stavros Georgiou","position":"Forward","number":74,"birthDate":"2004-10-19","age":"16","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/431250.jpg","affiliation":{"name":"Germany","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"373387","country":"Cyprus","firstName":"Manolis","lastName":"Charalambous","name":"Manolis Charalambous","position":"Midfielder","number":0,"birthDate":"2003-04-01","age":"18","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373387.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"6762","country":"Serbia","firstName":"Vujadin","lastName":"Savic","name":"Vujadin Savic","position":"Defender","number":0,"birthDate":"1990-07-01","age":"31","height":194,"weight":83,"thumbnailSrc":"https://images.onefootball.com/players/180/6762.jpg","affiliation":{"name":"Serbia","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/rs.png"}},{"id":"382019","country":"Brazil","firstName":"Carlos Eduardo","lastName":"Oliveira Dias","name":"Carlos Eduardo","position":"Defender","number":5,"birthDate":"2000-01-23","age":"21","height":179,"weight":73,"thumbnailSrc":"https://images.onefootball.com/players/180/382019.jpg","affiliation":{"name":"Brazil","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/br.png"}},{"id":"268659","country":"Cyprus","firstName":"Andreas","lastName":"Katsantonis","name":"Andreas Katsantonis","position":"Forward","number":16,"birthDate":"2000-02-16","age":"21","height":181,"weight":81,"thumbnailSrc":"https://images.onefootball.com/players/180/268659.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"434222","country":"Cyprus","firstName":"Ioannis","lastName":"Tsoutsouki","name":"Ioannis Tsoutsouki","position":"Midfielder","number":13,"birthDate":"2004-04-14","age":"17","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/434222.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"337577","country":"Cyprus","firstName":"Constantinos","lastName":"Karayiannis","name":"Constantinos Karayiannis","position":"Defender","number":0,"birthDate":"2000-04-02","age":"21","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/337577.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"310591","country":"Cyprus","firstName":"Giorgos","lastName":"Theodoulidis","name":"Giorgos Theodoulidis","position":"Goalkeeper","number":0,"birthDate":"2001-01-12","age":"20","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/310591.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"6683","country":"Belgium","firstName":"Richard Maciel","lastName":"Sousa Campos","name":"Danilo","position":"Midfielder","number":0,"birthDate":"1990-01-13","age":"31","height":174,"weight":71,"thumbnailSrc":"https://images.onefootball.com/players/180/6683.jpg","affiliation":{"name":"Belgium","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/be.png"}},{"id":"74667","country":"Cyprus","firstName":"Christos","lastName":"Wheeler","name":"Christos Wheeler","position":"Defender","number":42,"birthDate":"1997-06-29","age":"24","height":171,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/74667.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"303639","country":"Greece","firstName":"Apostolos","lastName":"Tsilingiris","name":"Apostolos Tsilingiris","position":"Goalkeeper","number":75,"birthDate":"2000-09-06","age":"20","height":190,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/303639.jpg","affiliation":{"name":"Greece","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/gr.png"}},{"id":"265651","country":"Brazil","firstName":"Rafael","lastName":"Santos de Sousa","name":"Rafael Santos","position":"Defender","number":4,"birthDate":"1998-02-02","age":"23","height":184,"weight":72,"thumbnailSrc":"https://images.onefootball.com/players/180/265651.jpg","affiliation":{"name":"Brazil","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/br.png"}},{"id":"223713","country":"Nigeria","firstName":"Francis","lastName":"Uzoho","name":"Francis Uzoho","position":"Goalkeeper","number":23,"birthDate":"1998-10-28","age":"22","height":196,"weight":91,"thumbnailSrc":"https://images.onefootball.com/players/180/223713.jpg","affiliation":{"name":"Nigeria","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ng.png"}},{"id":"26783","country":"Congo (Kinshasa)","firstName":"Dieumerci","lastName":"Ndongala","name":"Dieumerci Ndongala","position":"Forward","number":77,"birthDate":"1991-06-14","age":"30","height":170,"weight":61,"thumbnailSrc":"https://images.onefootball.com/players/180/26783.jpg","affiliation":{"name":"Congo (Kinshasa)","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cd.png"}},{"id":"373395","country":"Cyprus","firstName":"Marios","lastName":"Kokkinoftas","name":"Marios Kokkinoftas","position":"Midfielder","number":40,"birthDate":"2003-03-15","age":"18","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373395.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"373289","country":"Cyprus","firstName":"Constantinos","lastName":"Vrontis","name":"Constantinos Vrontis","position":"Defender","number":0,"birthDate":"2002-05-17","age":"19","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373289.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"373420","country":"Cyprus","firstName":"Nicolas","lastName":"Demetriou","name":"Nicolas Demetriou","position":"Midfielder","number":59,"birthDate":"2002-04-22","age":"19","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373420.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"268713","country":"Greece","firstName":"Dimitrios","lastName":"Priniotaki","name":"Dimitrios Priniotaki","position":"Goalkeeper","number":12,"birthDate":"1999-03-11","age":"22","height":186,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/268713.jpg","affiliation":{"name":"Greece","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/gr.png"}},{"id":"373126","country":"Cyprus","firstName":"Stavros","lastName":"Gavriel","name":"Stavros Gavriel","position":"Defender","number":25,"birthDate":"2002-01-29","age":"19","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373126.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"373328","country":"Cyprus","firstName":"Giannis","lastName":"Satsias","name":"Giannis Satsias","position":"Midfielder","number":18,"birthDate":"2002-12-28","age":"18","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373328.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"398434","country":"Cyprus","firstName":"Cháris","lastName":"Chatzigavriíl","name":"Cháris Chatzigavriíl","position":"Goalkeeper","number":80,"birthDate":"2003-11-22","age":"17","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/398434.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"27303","country":"Hungary","firstName":"Paulo","lastName":"Vinícius Souza dos Santos","name":"Paulo Vinícius","position":"Defender","number":3,"birthDate":"1990-02-21","age":"31","height":184,"weight":77,"thumbnailSrc":"https://images.onefootball.com/players/180/27303.jpg","affiliation":{"name":"Hungary","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/hu.png"}},{"id":"430302","country":"Cyprus","firstName":"Stylianos","lastName":"Vrontis","name":"Stylianos Vrontis","position":"Midfielder","number":66,"birthDate":"2004-11-05","age":"16","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/430302.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"182914","country":"Cyprus","firstName":"Antreas","lastName":"Paraskevas","name":"Antreas Paraskevas","position":"Goalkeeper","number":0,"birthDate":"1998-09-15","age":"22","height":187,"weight":79,"thumbnailSrc":"https://images.onefootball.com/players/180/182914.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"316166","country":"Cyprus","firstName":"Paris","lastName":"Polikarpou","name":"Paris Polikarpou","position":"Midfielder","number":0,"birthDate":"2000-09-23","age":"20","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/316166.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"216339","country":"Greece","firstName":"Konstantinos","lastName":"Apostolakis","name":"Konstantinos Apostolakis","position":"Midfielder","number":2,"birthDate":"1999-05-28","age":"22","height":178,"weight":74,"thumbnailSrc":"https://images.onefootball.com/players/180/216339.jpg","affiliation":{"name":"Greece","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/gr.png"}},{"id":"373253","country":"Cyprus","firstName":"Iasonas","lastName":"Toumazos","name":"Iasonas Toumazos","position":"Midfielder","number":99,"birthDate":"2003-04-24","age":"18","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373253.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"3889","country":"Equatorial Guinea","firstName":"Emilio","lastName":"Nsue","name":"Emilio Nsue","position":"Defender","number":22,"birthDate":"1989-09-30","age":"31","height":182,"weight":77,"thumbnailSrc":"https://images.onefootball.com/players/180/3889.jpg","affiliation":{"name":"Equatorial Guinea","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/gq.png"}},{"id":"33674","country":"Argentina","firstName":"Tomás Sebastián","lastName":"De Vincenti","name":"Tomás De Vincenti","position":"Midfielder","number":10,"birthDate":"1989-02-09","age":"32","height":178,"weight":72,"thumbnailSrc":"https://images.onefootball.com/players/180/33674.jpg","affiliation":{"name":"Argentina","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ar.png"}},{"id":"62690","country":"Norway","firstName":"Marius","lastName":"Lundemo","name":"Marius Lundemo","position":"Midfielder","number":8,"birthDate":"1994-04-11","age":"27","height":189,"weight":82,"thumbnailSrc":"https://images.onefootball.com/players/180/62690.jpg","affiliation":{"name":"Norway","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/no.png"}},{"id":"336317","country":"Argentina","firstName":"Facundo Gabriel","lastName":"Zabala","name":"Facundo Zabala","position":"Defender","number":36,"birthDate":"1999-01-02","age":"22","height":172,"weight":65,"thumbnailSrc":"https://images.onefootball.com/players/180/336317.jpg","affiliation":{"name":"Argentina","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ar.png"}},{"id":"167760","country":"Cyprus","firstName":"Antreas","lastName":"Karo","name":"Antreas Karo","position":"Defender","number":0,"birthDate":"1996-09-09","age":"24","height":190,"weight":80,"thumbnailSrc":"https://images.onefootball.com/players/180/167760.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"17010","country":"Georgia","firstName":"Tornike","lastName":"Okriashvili","name":"Tornike Okriashvili","position":"Midfielder","number":0,"birthDate":"1992-02-12","age":"29","height":181,"weight":71,"thumbnailSrc":"https://images.onefootball.com/players/180/17010.jpg","affiliation":{"name":"Georgia","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ge.png"}},{"id":"8029","country":"Cyprus","firstName":"Giorgios","lastName":"Merkis","name":"Giorgios Merkis","position":"Defender","number":30,"birthDate":"1984-07-30","age":"37","height":187,"weight":83,"thumbnailSrc":"https://images.onefootball.com/players/180/8029.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"373353","country":"Cyprus","firstName":"Nicolas","lastName":"Koutsakos","name":"Nicolas Koutsakos","position":"Forward","number":89,"birthDate":"2003-11-14","age":"17","height":0,"weight":0,"thumbnailSrc":"https://images.onefootball.com/players/180/373353.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"11944","country":"England","firstName":"Joe","lastName":"Garner","name":"Joe Garner","position":"Forward","number":41,"birthDate":"1988-04-12","age":"33","height":178,"weight":75,"thumbnailSrc":"https://images.onefootball.com/players/180/11944.jpg","affiliation":{"name":"England","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/gb-eng.png"}},{"id":"337568","country":"Jordan","firstName":"Omar Hani Ismail","lastName":"Al Zebdieh","name":"Omar Hani","position":"Defender","number":0,"birthDate":"1999-06-27","age":"22","height":168,"weight":76,"thumbnailSrc":"https://images.onefootball.com/players/180/337568.jpg","affiliation":{"name":"Jordan","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/jo.png"}},{"id":"143201","country":"Cyprus","firstName":"Neophytos","lastName":"Michael","name":"Neophytos Michael","position":"Goalkeeper","number":0,"birthDate":"1993-12-16","age":"27","height":190,"weight":85,"thumbnailSrc":"https://images.onefootball.com/players/180/143201.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}},{"id":"11216","country":"Georgia","firstName":"Murtaz","lastName":"Daushvili","name":"Murtaz Daushvili","position":"Midfielder","number":0,"birthDate":"1989-05-01","age":"32","height":176,"weight":74,"thumbnailSrc":"https://images.onefootball.com/players/180/11216.jpg","affiliation":{"name":"Georgia","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/ge.png"}},{"id":"7598","country":"Cyprus","firstName":"Giorgos","lastName":"Efrem","name":"Giorgos Efrem","position":"Midfielder","number":7,"birthDate":"1989-07-05","age":"32","height":171,"weight":65,"thumbnailSrc":"https://images.onefootball.com/players/180/7598.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}}],"officials":[{"countryName":"Cyprus","id":"60056","firstName":"Savvas","lastName":"Poursaitidis","country":"CY","position":"Coach","thumbnailSrc":"https://images.onefootball.com/coaches/180/60056.jpg","affiliation":{"name":"Cyprus","thumbnailSrc":"https://images.onefootball.com/icons/countries/56/cy.png"}}],"colors":{"shirtColorHome":"FFF347","shirtColorAway":"09004D","crestMainColor":"4F2C7D","mainColor":"FFF347"}}},"message":"Team feed successfully generated"}
	//  `)

	counterUrlID = 1

	for {
		if len(teamList) < maxTeams {
			finalURL = urlGenerator(urlStart, counterUrlID, urlEnd)
			resp, err := http.Get(finalURL)
			if err != nil {
				log.Println("ERROR Get: ", err)
				counterUrlID++
			}
			respByte, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println("ERROR ReadAll: ", err)
				counterUrlID++
			}
			err = json.Unmarshal(respByte, &tmp)
			if err != nil {
				log.Println("ERROR Unmarshal: ", err)
				counterUrlID++
			}

			if tmp.Status == "ok" && validationClub(tmp.Data.Team.Name) {
				for _, v := range tmp.Data.Team.Players {
					onePlayer = playerData{
						FirstName:    v.FirstName,
						LastName:     v.LastName,
						Age:          v.Age,
						Team:         tmp.Data.Team.Name,
						Country:      tmp.Data.Team.CountryName,
						NationalTeam: tmp.Data.Team.IsNational,
					}
					if onePlayer.Age != "" && onePlayer.Country != "" && onePlayer.FirstName != "" && onePlayer.LastName != "" && onePlayer.Team != "" {
						listOfAllPlayer = append(listOfAllPlayer, onePlayer)
					}
				}
				teamList = append(teamList, onePlayer.Team)
			}
			if len(teamList) == maxTeams {
				break
			}
			counterUrlID++
			log.Println("Len of Teams", len(teamList))
			time.Sleep(250 * time.Millisecond)
		}
	}

	// fmt.Println("----")
	// fmt.Println("Ungeordnet: ", listOfAllPlayer)
	// fmt.Println("+++++", listOfAllPlayer)

	//Sort Part
	dataSort(listOfAllPlayer)

	for i, v := range listOfAllPlayer {
		fmt.Printf("Num: %d -> %s, %s, %s, %v, %s, %s \n", i, v.FirstName, v.LastName, v.Age, v.NationalTeam, v.Team, v.Country)
	}

	// fmt.Println("xxxxxx")
	// fmt.Println("by name:", listOfAllPlayer)
	// fmt.Println("yyyyyy")
}

func validationClub(lookup string) bool {
	switch lookup {
	case "Manchester United":
		return true
	case "Arsenal":
		return true
	case "Chelsea":
		return true
	case "Barcelona":
		return true
	case "Bayern Munich":
		return true
	case "Real Madrid":
		return true
	default:
		return false
	}
}

func urlGenerator(urlStart string, num int, urlEnd string) string {
	finalURL := fmt.Sprintf("%s%s%s", urlStart, strconv.Itoa(num), urlEnd)
	return finalURL
}

// func addTeamToList(teamList []string, team string) []string {
// 	teamList = append(teamList, team)
// 	return teamList
// }

// func unmarshalBytes(jsonByteData []byte, tmp *schema) error {
// 	err := json.Unmarshal(jsonByteData, &tmp)
// 	if err != nil {
// 		log.Println("ERROR Unmarshal: ", err)
// 		return err
// 	}
// 	return nil
// }

func dataSort(listOfAllPlayer []playerData) {
	sort.SliceStable(listOfAllPlayer, func(i, j int) bool {
		return listOfAllPlayer[i].FirstName < listOfAllPlayer[j].FirstName
	})
}
