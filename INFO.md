# OneFootball - Coding challenge

- Built in Go version 1.16.
- Uses only Golang standard library. No additional third party libraries are used.

- The program generates the output specified in the task.
- Furthermore, the program provides additional information that may be useful for debugging the JSON files:

1. According to the task, the API endpoints whose JSON body is incorrect (because some values in mandatory fields are empty) / missing are rendered in stdOut.
2. In addition, the Team Ids and the Player Ids where the error occurred are rendered in stdOut too.

- To run the code, please navigate in /cmd folder and type "go run main.go client.go" in your terminal.
