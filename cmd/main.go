package main

import (
	"OneFootball_Zusammenfuehren/config"
	"OneFootball_Zusammenfuehren/helpers"
	"OneFootball_Zusammenfuehren/models"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

const (
	urlStart            string = "https://api-origin.onefootball.com/score-one-proxy/api/teams/en/"
	urlEnd              string = ".json"
	maxNumOfClubs       int    = 6   // This is the total number of local soccer teams.
	maxNumOfNations     int    = 4   // This is the total number of national soccer teams.
	maxNumOfAPIRequests int    = 200 // This value sets the maximum number of server requests that may be made. This value is used as a termination condition to avoid sending an endless number of requests.
)

var tmp models.Schema
var listOfLocalFootballPlayers []*models.PlayerData
var listOfNationalFootballPlayers []*models.PlayerData
var onePlayer *models.PlayerData
var listOfFootballClubs []*string
var listOfNations []*string
var playerIdsWithErrors []*string
var conf config.Config   // conf access to configuration struct und creates a Variable in main.go
var infoLog *log.Logger  // infolog stores all relevant information, e.g. in case some fields in the JSON body have no values or if the status = not ok is found in the JSON body.
var errorLog *log.Logger // errorlog stores all errors that may occur.
var counterRequest int

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {
	counterRequest = 1
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	conf.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	conf.ErrorLog = errorLog
	helpers.NewHelpers(&conf) // This creates config file in config package.
	for {
		if len(listOfNations) != maxNumOfNations || len(listOfFootballClubs) != maxNumOfClubs {
			if counterRequest == maxNumOfAPIRequests {
				conf.ErrorLog.Printf("we have reached the maximum number of api requests: (%s out of %s) \n", strconv.Itoa(counterRequest), strconv.Itoa(maxNumOfAPIRequests))
				break
			}
			finalURL := urlGenerator(urlStart, counterRequest, urlEnd)
			apiClient := NewApiClient(finalURL, 10*time.Second)
			respBytes, err := apiClient.GetJson()
			if err != nil {
				conf.ErrorLog.Println(err)
				counterRequest++
				continue
			}
			err = apiClient.Unmarshall(respBytes)
			if err != nil {
				conf.ErrorLog.Println(err)
				counterRequest++
				continue
			}
			if tmp.Status != "ok" || tmp.Status == "" {
				conf.InfoLog.Printf("Status for URL '%s' is 'not ok' or value is empty. \n", finalURL)
				counterRequest++
				continue
			}
			if !validationClubsAndNations(tmp.Data.Team.Name) {
				counterRequest++
				continue
			}
			for _, v := range tmp.Data.Team.Players {
				onePlayer = &models.PlayerData{
					TeamId:          tmp.Data.Team.ID,
					PlayerId:        v.ID,
					FirstName:       v.FirstName,
					LastName:        v.LastName,
					Age:             v.Age,
					PlaceholderName: tmp.Data.Team.Name,
					IsNational:      tmp.Data.Team.IsNational,
				}
				if onePlayer.Age == "" || onePlayer.FirstName == "" || onePlayer.LastName == "" || onePlayer.PlaceholderName == "" {
					conf.InfoLog.Printf("URL '%s' can't be processed successfully \n Some values are missing in mandatory fields for team ID: %s / player ID: %s. \n", finalURL, strconv.Itoa(onePlayer.TeamId), onePlayer.PlayerId)
					playerIdsWithErrors = append(playerIdsWithErrors, &onePlayer.PlayerId)
					continue
				}
				appendPlayersToSlice(onePlayer.IsNational)
			}
			appendNationsAndClubsToSlice(tmp.Data.Team.IsNational)
			counterRequest++
		} else {
			conf.InfoLog.Printf("we have logged %s out of %s nations \n", strconv.Itoa(len(listOfNations)), strconv.Itoa(maxNumOfNations))
			conf.InfoLog.Printf("we have logged %s out of %s football clubs\n", strconv.Itoa(len(listOfFootballClubs)), strconv.Itoa(maxNumOfClubs))
			conf.InfoLog.Printf("we have send %s out of %s send requests\n", strconv.Itoa(counterRequest), strconv.Itoa(maxNumOfAPIRequests))
			break
		}
	} // for loop is the "main" function of this program.
	dataSort(listOfLocalFootballPlayers)                                            //dataSort sorts the slice alphabetically.
	transferNationValues(listOfLocalFootballPlayers, listOfNationalFootballPlayers) //transferNationValues extracts nationalteam value from a field of a slice 'listOfNationalPlayers' and applies it to the final struct of another slice 'listOfClubPlayers' which will be used for printing.
	stdOut(listOfLocalFootballPlayers)                                              //In the following, the output (stdout) is now created using a switch statement.
	return nil
}

// Appends national and local football players to respective slice.
func appendPlayersToSlice(l bool) {
	switch l {
	case true:
		listOfNationalFootballPlayers = append(listOfNationalFootballPlayers, onePlayer)
	default:
		listOfLocalFootballPlayers = append(listOfLocalFootballPlayers, onePlayer)
	}
}

// Appends national and local football clubs to respective slice.
func appendNationsAndClubsToSlice(t bool) {
	switch t {
	case true:
		listOfNations = append(listOfNations, &tmp.Data.Team.Name)
	default:
		listOfFootballClubs = append(listOfFootballClubs, &tmp.Data.Team.Name)
	}
}

// We have stored football players in two different slice: Slice 'listNationPlayer' and Slice 'listClubPlayer'.
// The slices differ in that the 'listNationPlayer' slice has the name of the national team stored
// and the 'listClubPlayer' slice has the name of the local soccer team stored.
// transferNationValues does three things:
// 1: It searches both slices for players with the same first name and last name. The result is the intersection of both slices.
// 2: It extracts the name of the national team of a specific player, lets say 'player A', from slice 'listNationPlayer'.
// 3: It stores this value in a field of a struct of the same player, who also exists in slice 'listClubPlayer'
func transferNationValues(listClubPlayer []*models.PlayerData, listNationPlayer []*models.PlayerData) {
	for _, nationalSpieler := range listNationPlayer {
		nationalSpieler.NameOfNationalteam = nationalSpieler.PlaceholderName
		for _, clubplayer := range listClubPlayer {
			clubplayer.NameOfFootballClub = clubplayer.PlaceholderName
			if nationalSpieler.FirstName == clubplayer.FirstName && nationalSpieler.LastName == clubplayer.LastName {
				clubplayer.NameOfNationalteam = nationalSpieler.PlaceholderName
			}
		}
	}
}

// ValidationClubs is a switch statement used to filter tmp data file - which we extracted from the JSON response - based on national and local teams named in the coding challenge task.
func validationClubsAndNations(lookup string) bool {
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
	case "Germany":
		return true
	case "England":
		return true
	case "France":
		return true
	case "Spain":
		return true
	default:
		return false
	}
}

// urlGenerator generates the API URL to query.
func urlGenerator(urlStart string, num int, urlEnd string) string {
	finalURL := fmt.Sprintf("%s%s%s", urlStart, strconv.Itoa(num), urlEnd)
	return finalURL
}

// dataSort sorts the football player slice alphabetically.
func dataSort(list []*models.PlayerData) {
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].FirstName < list[j].FirstName
	})
}

//stdOut creates output (stdout) using a switch statement.
func stdOut(listOfLocalFootballPlayers []*models.PlayerData) {
	for i, v := range listOfLocalFootballPlayers {
		switch v.NameOfNationalteam {
		case "":
			fmt.Printf("Num: %d -> %s, %s, %s, %v \n", i, v.FirstName, v.LastName, v.Age, v.NameOfFootballClub)
		default:
			fmt.Printf("Num: %d -> %s, %s, %s, %v, %v \n", i, v.FirstName, v.LastName, v.Age, v.NameOfNationalteam, v.NameOfFootballClub)
		}
	}
}
