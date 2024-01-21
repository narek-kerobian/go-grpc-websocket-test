package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,
}

type GameOutcome struct {
    TeamA           string `json:"teamA"`
    TeamB           string `json:"teamB"`
    GameInProgress  bool `json:"gameInProgress"`
    Score           []int `json:"score"`
}

type Leaderboard struct {
    Sport   string `json:"sport"`
    Leaders []string `json:"leaders"`
}

type RequestData struct {
    Channels []string `json:"channels"`
}

type ResponseData struct {
    GameOutcomes *[]GameOutcome `json:"gameOutcomes"`
    Leaderboards *[]Leaderboard `json:"leaderboards"`
}

func HandleSocketConnections(c *gin.Context) {

    // Allow requests from all origins
    upgrader.CheckOrigin = func(r *http.Request) bool {
        return true
    }

    // Upgrade http connection
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println("[WS Log] Failed to upgrade http: ", err)
    }
    log.Println("[WS Log] Client successfully connected")

    var data RequestData
    dataChan := make(chan RequestData, 1)

    go func () {
        for {
            err = conn.ReadJSON(&data)
            if err != nil {
                log.Println("[WS Log] Error reading the message: ", err)
                return
            }
            log.Println(fmt.Sprintf("[WS Log] Received request for channels: %v", data.Channels))
            dataChan<-data
        }
    }()

    for {
        select  {
        case <-c.Done():
            conn.Close()
            return
        case data := <-dataChan:
            response := ResponseData{
                GameOutcomes: nil,
                Leaderboards: nil,
            }
            for _, channel := range data.Channels {
                switch channel {
                    case "leaderboard":
                        response.Leaderboards = _buildLiederboard()
                        break
                    case "outcome":
                        response.GameOutcomes = _buildOutcome()
                        break
                }
            }

            log.Println(fmt.Sprintf("[WS Log] Sending response for channels: %v", data.Channels))
            responseString, err := json.Marshal(response)
            err = conn.WriteMessage(websocket.TextMessage, responseString)

            if err != nil {
                fmt.Println("[WS Log] Write failed:", err)
                break
            }
        }
        time.Sleep(time.Second * 1)
    }

}

func _buildLiederboard() *[]Leaderboard {
    var leaderboards []Leaderboard

    for i:=0; i<=10; i++ {
        lb := Leaderboard{
            Sport: fmt.Sprintf("Sport %d", i),
            Leaders: []string{"TeamA", "TeamB", "TeamC", "TeamD", "TeamE", "TeamF"},
        }
        leaderboards = append(leaderboards, lb)
    }

    return &leaderboards
}

func _buildOutcome() *[]GameOutcome {
    var outcomes []GameOutcome

    for i:=0; i<=10; i++ {
        gmo := GameOutcome{
            TeamA: fmt.Sprintf("TeamA Name %d", i),
            TeamB: fmt.Sprintf("TeamB Name %d", i),
            GameInProgress: true,
            Score: []int{10, i},
        }
        outcomes = append(outcomes, gmo)
    }

    return &outcomes
}

