package src


import (
	"github.com/robfig/cron"
    "github.com/google/go-github/github"
    "log"
)

var gitcron *cron.Cron

var githubClient *github.Client
var githubUsername string
var githubRepo string

func InitCron(gc *github.Client) {
    if gitcron == nil {
		gitcron = cron.New()
	} else {
		gitcron.Stop()
	}

    githubClient = gc

	gitcron.AddFunc("@every 60s", func() {
		acceptPullRequest()
	})
	
	gitcron.AddFunc("@every 300s", func() {
		acceptAllPullRequest("tokopedia", "etl-dwh")
	})

    githubUsername = "tokopedia"
    githubRepo = "rechargeapp"

	gitcron.Start()
}

func acceptPullRequest(){
    results, _, _ := githubClient.PullRequests.List(githubUsername, githubRepo, nil)
    for _, v := range results {
        if *v.State == "open" && *v.Base.Ref != "master" {
            mergeResult, _, err := githubClient.PullRequests.Merge(githubUsername, githubRepo, *v.Number, "", nil)
            if err != nil {
                log.Println(err)
            } else {
                if mergeResult.Merged != nil && *mergeResult.Merged {
                    log.Println("Success merged : " + *v.Head.Label + "(" + *v.Base.Ref + ")")
                }
            }
        }
    }
}

func acceptAllPullRequest(username string, repo string){
    results, _, _ := githubClient.PullRequests.List(username, repo, nil)
    for _, v := range results {
        if *v.State == "open" {
            mergeResult, _, err := githubClient.PullRequests.Merge(username, repo, *v.Number, "", nil)
            if err != nil {
                log.Println(err)
            } else {
                if mergeResult.Merged != nil && *mergeResult.Merged {
                    log.Println("Success merged : " + *v.Head.Label + "(" + *v.Base.Ref + ")")
                }
            }
        }
    }
}