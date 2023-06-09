package spider

import (
	"fmt"
	"time"
)

type Game struct {
	HomeTeam      string
	AwayTeam      string
	StartTime     string
	Score         string
}

func Start(){
	t := time.Unix(int64(time.Now().Unix()), 0)
	date := t.Format("2006-01-02")
	url := fmt.Sprintf("https://m.china.nba.cn/stats2/season/schedule.json?gameDate=%s&locale=zh_CN&tz=%s",date, "%2B8")
	fmt.Println(url)
}