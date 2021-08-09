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
