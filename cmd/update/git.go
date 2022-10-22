package update

import (
	"context"
	"log"
	"net/http"

	"github.com/google/go-github/v45/github"
)


func getGithubVersion() (string,error){
	client:=github.NewClient(http.DefaultClient)
	release,response,err:=client.Repositories.GetLatestRelease(context.Background(),"guionardo", "go-dev")
	if err!=nil{
		log.Printf("Failed to get latest release: %v", err)
		return "",err
	}
	if response.StatusCode!=http.StatusOK{
		log.Printf("Failed to get latest release: %v", response.StatusCode)
		return "",err
	}
	return release.GetTagName(),nil
}