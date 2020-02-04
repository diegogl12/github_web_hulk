package services

import(
	"encoding/csv"
	"os"
	"../structs"
	"log"
)

func CreateAndWriteCsv(branches []structs.Branch) {
    file, err := os.Create("result.csv")
    checkError("Cannot create file", err)
    defer file.Close()

    writer := csv.NewWriter(file)
	defer writer.Flush()
	
	for _, branch := range branches {
		branchString := toString(branch)
		err = writer.Write(branchString)
		checkError("Cannot write to file", err)
	}
}

func toString(branch structs.Branch) []string{
	var response = []string {branch.Name, branch.Commit.Sha, branch.Commit.Commit.Message, branch.Commit.Commit.Author.Name, branch.Commit.Commit.Author.Date}
	return response
}

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}
