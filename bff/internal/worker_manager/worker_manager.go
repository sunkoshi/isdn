package worker_manager

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/gorilla/websocket"
	db "github.com/orted-org-isdn-bff/db/sqlc"
	"github.com/orted-org-isdn-bff/util"
)

type NodeInfo struct {
	id   int64
	conn *websocket.Conn
}

type WorkerManager struct {
	store          *db.Queries
	nodes          map[int64]NodeInfo
	nodesLang      map[string][]int64
	requestChanMap map[int64]chan FunctionExecutionResult
}

func New(store *db.Queries) *WorkerManager {
	return &WorkerManager{
		store:          store,
		nodes:          make(map[int64]NodeInfo),
		nodesLang:      make(map[string][]int64),
		requestChanMap: make(map[int64]chan FunctionExecutionResult),
	}
}

func (wm *WorkerManager) AddNewNode(id int64, languages []string, conn *websocket.Conn) {
	log.Printf("adding new node with ID=%d, LANG=%v", id, languages)
	for _, v := range languages {
		if nodeIdArr, ok := wm.nodesLang[v]; ok {
			nodeIdArr = append(nodeIdArr, id)
			wm.nodesLang[v] = nodeIdArr
		} else {
			// array does not exists
			var nodeIdArr []int64 = make([]int64, 0)
			nodeIdArr = append(nodeIdArr, id)
			wm.nodesLang[v] = nodeIdArr
		}
	}
	wm.nodes[id] = NodeInfo{
		id:   id,
		conn: conn,
	}
	go wm.ListenOnNodeConnectionForResult(conn)
}

func (wm *WorkerManager) AddNewFunctionCallJob(fnID int64, language string, requestID int64, input map[string]interface{}, resultChan chan FunctionExecutionResult) error {
	var callRequest map[string]interface{} = map[string]interface{}{}
	callRequest["request_id"] = requestID
	callRequest["language"] = language
	callRequest["input"] = input
	callRequest["function_id"] = fnID

	payload, err := json.Marshal(callRequest)
	if err != nil {
		return err
	}

	if langSupportedNode, ok := wm.nodesLang[language]; ok {
		if len(langSupportedNode) == 0 {
			return errors.New("no node available for this load")
		}
		nodeIndex := util.RandomInt(0, int64(len(langSupportedNode)-1))
		nodeId := langSupportedNode[nodeIndex]
		node := wm.nodes[nodeId]
		wm.requestChanMap[requestID] = resultChan
		return node.conn.WriteMessage(websocket.TextMessage, payload)
	} else {
		return errors.New("no node available for this load")
	}
}

func (wm *WorkerManager) ListenOnNodeConnectionForResult(conn *websocket.Conn) {
	for {
		_, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error", err)
			continue
		}
		var res FunctionExecutionResult
		err = util.JSONParse(payload, &res)
		if err != nil {
			log.Println(err)
		}
		log.Println("received result from worker node for request id", res.RequestID)
		if channelToSendData, ok := wm.requestChanMap[int64(util.ToInt(res.RequestID))]; ok {
			channelToSendData <- res
		} else {
			log.Println("no channel found to send data")
		}
	}
}
