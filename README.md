# Backend Take Home Assingment

# The Task

Given API endpoint:
`https://api-origin.onefootball.com/score-one-proxy/api/teams/en/%team_id%.json`
(`%team_id%` must be replaced with an unsigned integer)

Using the API endpoint, find the following teams by name:

* Germany
* England
* France
* Spain
* Manchester United
* Arsenal
* Chelsea
* Barcelona
* Real Madrid
* Bayern Munich

Extract all the players from the given teams and render to stdout the information about players alphabetically ordered by name.

Each player entry should contain the following information: full name; age; list of teams.

**Output Example:**

```
1. Alexander Mustermann; 25; France, Manchester Utd
2. Brad Exampleman; 30; Arsenal
3. ...
```

**Requirements:**

go >= 1.8

**Delivery of the task:**

Push your code to this repository. This repo is yours, feel free to
change any thing in this repo (including the README.md)

**FAQ:**

* Discovering a valid **range** of IDs is part of the task
* You do not need to scan all the IDs on the API, only enough to get
information about all the teams listed
* Some IDs might return errors, your code should handle that
  gracefully, however all teams on the requested list can be found
  through the API

# Coding Solution

- Built in Go version 1.16.
- Uses only Golang standard library. No additional third party libraries are used.

- The program generates the output specified in the task.
- Furthermore, the program provides additional information that may be useful for debugging the JSON files:

1. According to the task, the API endpoints whose JSON body is incorrect (because some values in mandatory fields are empty) / missing are rendered in stdOut.
2. In addition, the Team Ids and the Player Ids where the error occurred are rendered in stdOut too.

- To run the code, please navigate in /cmd folder and type "go run main.go client.go" in your terminal.

