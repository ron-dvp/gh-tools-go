package main

import (
	"flag"
	"fmt"
	"ghtools/tools"
	"log"
	"os"

	"github.com/TwiN/go-color"
)

func main() {

	// TODO: if not args show command line help tool
	// Read flags
	checkId := flag.Bool("u", false, "Check auth caller")
	listAll := flag.Bool("l", false, "List your account repos")
	listAll_ := flag.Bool("list", false, "List your account repos")
	createRepo := flag.String("n", "", "Create a new repository")
	description := flag.String("d", "", "Attach description to your repository")
	private := flag.Bool("p", false, "Set repository as private")

	flag.Parse()
	if *checkId == true {
		CheckAuth()
	}

	if *listAll || *listAll_ {
		ShowRepos()
	}

	if *createRepo != "" {
		NewRepo(*createRepo, *description, *private)
	}

}

func ShowRepos() {
	repos, err := tools.GetRepos()
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		os.Exit(1)
	}
	for _, repo := range repos {
		fmt.Printf("%s , private: %t\n", repo.Full_name, repo.Private)
	}
}

func CheckAuth() {
	res, err := tools.AuthUser()
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		os.Exit(1)
	}
	fmt.Printf(color.Ize(color.Cyan, "Autheticated user %s\n"), res.Login)
}

func NewRepo(name string, desc string, private bool) {

	res, err := tools.CreateRepo(tools.NewRepo{Name: name, Description: desc, Private: private})
	if err != nil {
		log.Fatalf("[-] Error: %s", err)
		os.Exit(1)
	}
	fmt.Printf(color.Ize(color.Yellow, "New repository created: %s\n"), res.Url)
}
