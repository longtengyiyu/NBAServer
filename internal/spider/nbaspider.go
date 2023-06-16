package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Repose struct {
	HomeTeam      string `json:"homeTeam"`
	HomeTeamEn    string `json:"homeTeamEn"`
	HomeTeamScore uint8 `json:"homeTeamScore"`
	AwayTeam      string `json:"awayTeam"`
	AwayTeamEn    string `json:"awayTeamEn"`
	AwayTeamScore uint8 `json:"awayTeamScore"`
	StartTime     string `json:"startTime"`
	Status        int `json:"status"`
	GameType      int `json:"gameType"`
}

type Schedule struct {
	Payload Payload `json:"payload"`
	Timestamp string `json:"timestamp"`
	Error Error `json:"error"`
}

type Payload struct {
	Dates [] Date `json:"dates"`
	UtcMillis string `json:"utc_millis"`
}

type Date struct {
	Games [] Game `json:"games"`
	GameCount string `json:"gameCount"`
	UtcMillis string `json:"utcMillis"`
}

type Error struct {
	Detail string `json:"detail"`
	IsError bool `json:"is_error"`
	Message string `json:"message"`
}

type Game struct {
	Boxscore Boxscore `json:"boxscore"`
	HomeTeam HomeTeam `json:"homeTeam"`
	AwayTeam AwayTeam `json:"awayTeam"`
	IfNecessary bool `json:"ifNecessary"`
	Profile GameInfo `json:"profile"`
}

type GameInfo struct {
	ArenaLocation string      `json:"arenaLocation"`
	ArenaName     string      `json:"arenaName"`
	AwayTeamId    string      `json:"awayTeamId"`
	DateTimeEt    string      `json:"dateTimeEt"`
	GameId        string      `json:"gameId"`
	HomeTeamId    string      `json:"homeTeamId"`
	Number        string      `json:"number"`
	ScheduleCode  interface{} `json:"scheduleCode"`
	SeasonType    string      `json:"seasonType"`
	Sequence      string      `json:"sequence"`
	UtcMillis     string      `json:"utcMillis"`
}

type HomeTeam struct{
	Profile Profile `json:"profile"`
	Matchup string `json:"matchup"`
}

type AwayTeam struct{
	Profile Profile `json:"profile"`
	Matchup string `json:"matchup"`
}

type Profile struct {
	Abbr              string `json:"abbr"`
	City              string `json:"city"`
	CityEn            string `json:"cityEn"`
	Code              string `json:"code"`
	Conference        string `json:"conference"`
	DisplayAbbr       string `json:"displayAbbr"`
	DisplayConference string `json:"displayConference"`
	Division          string `json:"division"`
	Id                string `json:"id"`
	IsAllStarTeam     bool   `json:"isAllStarTeam"`
	IsLeagueTeam      bool   `json:"isLeagueTeam"`
	LeagueId          string `json:"leagueId"`
	Name              string `json:"name"`
	NameEn            string `json:"nameEn"`
}

type Boxscore struct {
	Attendance            string `json:"attendance"`
	AwayScore             uint8    `json:"awayScore"`
	GameLength            string `json:"gameLength"`
	HomeScore             uint8    `json:"homeScore"`
	LeadChanges           int    `json:"leadChanges"`
	OfficialsDisplayName1 string `json:"officialsDisplayName1"`
	OfficialsDisplayName2 string `json:"officialsDisplayName2"`
	OfficialsDisplayName3 string `json:"officialsDisplayName3"`
	Period                string `json:"period"`
	PeriodClock           string `json:"periodClock"`
	Status                string `json:"status"`
	StatusDesc            string `json:"statusDesc"`
	Ties                  string `json:"ties"`
}

func Start(){
	fmt.Println(getUrl())
}

const currentTime = 1686585600 //int64(time.Now().Unix())
const gameType int = 0
var Res[] Repose


func getUrl() string{
	//t := time.Unix(int64(time.Now().Unix()), 0)
	date := getDate(int64(currentTime))
	return fmt.Sprintf("https://m.china.nba.cn/stats2/season/schedule.json?gameDate=%s&locale=zh_CN&tz=%s",date, "%2B8")
}

func Request(){
	url := getUrl()
	fmt.Println(url)
	// 发送HTTP GET请求
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}
	//清空
	Res = Res[:0]
	// 打印响应结果
	fmt.Println(string(body))
	var schedule Schedule
	err = json.Unmarshal(body, &schedule)
	if err != nil {
		fmt.Println("解析 JSON 数据出错:", err)
		return
	}
	fmt.Println(schedule)
	dates := schedule.Payload.Dates
	for _, date := range dates {
		currentDateStr := getDate(int64(currentTime))
		s,_ := strconv.ParseInt(date.UtcMillis, 10, 64)
		dateStr := getDate(s / 1000)
		println("currentDateStr:", currentDateStr)
		println("dateStr:", dateStr)
		if currentDateStr == dateStr {
			games := date.Games
			for _, game := range games {
				start,_ := strconv.ParseInt(game.Profile.UtcMillis, 10, 64)
				startTime := getTime(start / 1000)
				println("比赛开始时间：",startTime)
				homeTeam := game.HomeTeam.Profile.Name
				homeTeamEn := game.HomeTeam.Profile.Abbr
				homeScore := game.Boxscore.HomeScore
				awayTeam := game.AwayTeam.Profile.Name
				awayTeamEn := game.AwayTeam.Profile.Abbr
				awayScore := game.Boxscore.AwayScore
				status,_ := strconv.Atoi(game.Boxscore.Status)

				var re Repose
				re.HomeTeam = homeTeam
				re.HomeTeamEn = homeTeamEn
				re.AwayTeam = awayTeam
				re.AwayTeamEn = awayTeamEn
				re.StartTime = startTime
				re.HomeTeamScore = homeScore
				re.AwayTeamScore = awayScore
				re.GameType = gameType
				re.Status = status
				Res = append(Res, re)
			}
			break
		}
	}
	b,_ := json.Marshal(Res)
	println(string(b))
}

func getDate(timestamp int64) string{
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02")
}

func getTime(timestamp int64) string{
	t := time.Unix(int64(timestamp), 0)
	return t.Format("15:04:05")
}