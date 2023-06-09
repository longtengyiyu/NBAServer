CREATE TABLE SCORE(
  id INT(6) AUTO_INCREMENT PRIMARY KEY,
  arenaName VARCHAR(10) NOT NULL COMMENT '比赛地址',
  homeTeam VARCHAR(10) NOT NULL COMMENT '主队',
  homeTeamEn VARCHAR(10) NOT NULL COMMENT '英文名' ,
  awayTeam VARCHAR(10) NOT NULL COMMENT  '客队',
  awayTeamEn VARCHAR(10) NOT NULL COMMENT '英文名',
  homeTeamScores int(3)  DEFAULT 0 COMMENT '主场得分',
  awayTeamScores int(3) DEFAULT 0  COMMENT '客场得分',
  status int(1) DEFAULT 0 COMMENT '0 未开始 1进行中 2已结束',
  Started DATE NOT NULL  '比赛开始时间',
  Created DATE NOT NULL,
  Updated DATE NOT NULL
 );

go run nbaserver.go -f etc/nbaserver-api.yaml

要使用Go语言爬取https://china.nba.cn/schedule/#!/2023-06-08比赛动态，可以使用Go语言中的net/http包和goquery包。具体步骤如下：

1. 使用net/http包中的Get函数获取网页内容。
2. 使用goquery包解析网页内容，获取需要的比赛动态信息。
3. 将获取到的比赛动态信息存储到本地或者数据库中。

以下是一个简单的示例代码，可以获取到比赛的时间、参赛队伍、比分等信息：

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/PuerkitoBio/goquery"
)

func main() {
    url := "https://china.nba.cn/schedule/#!/2023-06-08"
    resp, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    doc.Find(".game-item").Each(func(i int, s *goquery.Selection) {
        time := s.Find(".game-time").Text()
        team1 := s.Find(".team1 .team-name").Text()
        team2 := s.Find(".team2 .team-name").Text()
        score1 := s.Find(".team1 .team-score").Text()
        score2 := s.Find(".team2 .team-score").Text()

        fmt.Printf("比赛时间：%s\n", time)
        fmt.Printf("参赛队伍：%s vs %s\n", team1, team2)
        fmt.Printf("比分：%s:%s\n", score1, score2)
    })
}
```


https://m.china.nba.cn/stats2/season/schedule.json?gameDate=2023-06-08&locale=zh_CN&tz=%2B8

```
{
	"context": {
		"user": {
			"countryCode": "CN",
			"countryName": "China",
			"locale": "zh_CN",
			"timeZone": "+08:00",
			"timeZoneCity": "Australia/Perth"
		},
		"device": {
			"clazz": null
		}
	},
	"error": {
		"detail": null,
		"isError": "false",
		"message": null
	},
	"payload": {
		"league": {
			"id": "00",
			"name": "NBA"
		},
		"season": {
			"isCurrent": "true",
			"rosterSeasonType": 4,
			"rosterSeasonYear": "2022",
			"rosterSeasonYearDisplay": "2022-2023",
			"scheduleSeasonType": 4,
			"scheduleSeasonYear": "2022",
			"scheduleYearDisplay": "2022-2023",
			"statsSeasonType": 4,
			"statsSeasonYear": "2022",
			"statsSeasonYearDisplay": "2022-2023",
			"year": "2022",
			"yearDisplay": "2022-2023"
		},
		"dates": [{
			"games": [{
				"profile": {
					"arenaLocation": "Miami, FL",
					"arenaName": "Kaseya Center",
					"awayTeamId": "1610612743",
					"dateTimeEt": "2023-06-07T20:30",
					"gameId": "0042200403",
					"homeTeamId": "1610612748",
					"number": "3",
					"scheduleCode": null,
					"seasonType": "4",
					"sequence": "1",
					"utcMillis": "1686184200000"
				},
				"boxscore": {
					"attendance": "20,019",
					"awayScore": 109,
					"gameLength": "145",
					"homeScore": 94,
					"leadChanges": 7,
					"officialsDisplayName1": "托尼 Brothers",
					"officialsDisplayName2": "约什 Tiven",
					"officialsDisplayName3": "凯文 斯科特",
					"period": "4",
					"periodClock": "",
					"status": "3",
					"statusDesc": "结束",
					"ties": "7"
				},
				"urls": [{
					"displayText": "联盟通",
					"type": "leaguepass",
					"value": "https://kbs.sports.qq.com/kbsweb/game.htm?mid=100000:56356204"
				}],
				"broadcasters": [],
				"homeTeam": {
					"profile": {
						"abbr": "MIA",
						"city": "迈阿密",
						"cityEn": "Miami",
						"code": "heat",
						"conference": "Eastern",
						"displayAbbr": "热火",
						"displayConference": "东部",
						"division": "东南分区",
						"id": "1610612748",
						"isAllStarTeam": false,
						"isLeagueTeam": true,
						"leagueId": "00",
						"name": "热火",
						"nameEn": "Heat"
					},
					"matchup": null
				},
				"awayTeam": {
					"profile": {
						"abbr": "DEN",
						"city": "丹佛",
						"cityEn": "Denver",
						"code": "nuggets",
						"conference": "Western",
						"displayAbbr": "掘金",
						"displayConference": "西部",
						"division": "西北分区",
						"id": "1610612743",
						"isAllStarTeam": false,
						"isLeagueTeam": true,
						"leagueId": "00",
						"name": "掘金",
						"nameEn": "Nuggets"
					},
					"matchup": null
				},
				"ifNecessary": false
			}],
			"gameCount": "1",
			"utcMillis": "1686153600000"
		}],
		"utcMillis": "1686153600000"
	},
	"timestamp": "1686301547346"
}
```

