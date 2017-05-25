package git


import (
	//"Golem/make"
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	//"bytes"
)

type (
	GitHubRepo struct {
		Name string
		Scope string
		Description string
	}
)

var (
	GitToken string 

)

func GitConnect(gitToken string){
	GitToken = gitToken

}
func CreateRepository ( g GitHubRepo ){
	//token := make.GitToken
	fmt.Println("I am here")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken : GitToken},
	)
	tc := oauth2.NewClient(ctx,ts)
	client := github.NewClient(tc)
	var scope bool
	if g.Scope == "private" || g.Scope == "Private"{
		scope = true
	} else {
		scope = false
	}

	repo := &github.Repository{
		Name : github.String(g.Name),
		Private : &scope,
		Description: github.String(g.Description),
	}
	client.Repositories.Create(ctx, "", repo)

}

