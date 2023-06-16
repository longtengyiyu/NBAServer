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
        "dates": [
            {
                "games": [
                    {
                        "profile": {
                            "arenaLocation": "Denver, CO",
                            "arenaName": "鲍尔体育馆",
                            "awayTeamId": "1610612748",
                            "dateTimeEt": "2023-06-12T20:30",
                            "gameId": "0042200405",
                            "homeTeamId": "1610612743",
                            "number": "5",
                            "scheduleCode": null,
                            "seasonType": "4",
                            "sequence": "1",
                            "utcMillis": "1686616200000"
                        },
                        "boxscore": {
                            "attendance": "19,537",
                            "awayScore": 66,
                            "gameLength": "100",
                            "homeScore": 64,
                            "leadChanges": 2,
                            "officialsDisplayName1": "马克 戴维斯",
                            "officialsDisplayName2": "埃德 Malloy",
                            "officialsDisplayName3": "大卫 Guthrie",
                            "period": "3",
                            "periodClock": "01:49",
                            "status": "2",
                            "statusDesc": "第3节",
                            "ties": "4"
                        },
                        "urls": [
                            {
                                "displayText": "联盟通",
                                "type": "leaguepass",
                                "value": "https://kbs.sports.qq.com/kbsweb/game.htm?mid=100000:56356206"
                            }
                        ],
                        "broadcasters": [],
                        "homeTeam": {
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
                        "awayTeam": {
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
                        "ifNecessary": false
                    }
                ],
                "gameCount": "1",
                "utcMillis": "1686585600000"
            }
        ],
        "utcMillis": "1686585600000"
    },
    "timestamp": "1686622798798"
}
```

新版本api接口数据

```
{
	"success": true,
	"gameCount": 0,
	"utcMillis": 1686585600000,
	"list": [{
			"homeTeamName": "",
			"homeTeamNameEn": "",
			"awayTeamName": "",
			"awayTeamNameEn": "",
			"homeTeamScore": 0,
			"awayTeamScore": 0,
			"status": 1,
			"startTime": 1686702600000,
			"gameType": 0 //0常规赛 1季后赛 2半决赛 3决赛
		},
		{
			"homeTeamName": "",
			"homeTeamNameEn": "",
			"awayTeamName": "",
			"awayTeamNameEn": "",
			"homeTeamScore": 0,
			"awayTeamScore": 0,
			"status": 1,
			"startTime": 1686702600000,
			"gameType": 0
		}
	]
}
```

