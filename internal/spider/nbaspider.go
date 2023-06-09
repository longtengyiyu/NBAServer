package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strings"
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
	url := fmt.Sprintf("https://china.nba.cn/schedule/#!/%s",date)
	fmt.Println(url)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var games []Game

	doc.Find(".game-box").Each(func(i int, s *goquery.Selection) {
		homeTeam := strings.TrimSpace(s.Find(".game-box-l .team-name").Text())
		awayTeam := strings.TrimSpace(s.Find(".game-box-r .team-name").Text())
		startTime := strings.TrimSpace(s.Find(".game-box-l .game-time").Text())
		score := strings.TrimSpace(s.Find(".game-box-l .game-score").Text())

		game := Game{
			HomeTeam:  homeTeam,
			AwayTeam:  awayTeam,
			StartTime: startTime,
			Score:     score,
		}
		games = append(games, game)
	})

	for _, game := range games {
		fmt.Printf("主队: %s\n", game.HomeTeam)
		fmt.Printf("客队: %s\n", game.AwayTeam)
		fmt.Printf("开始时间: %s\n", game.StartTime)
		fmt.Printf("比分: %s\n\n", game.Score)
	}
}