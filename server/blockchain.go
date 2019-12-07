//	Author: Shaquille Hall
//	Date: December 6, 2019
// 	Title: Blockchain Logic in Golang
//
// 	This package executes all blockchain logical operations:
// 	* 	Create
// 	* 	Mine
// 	*	Update
//
// 	This package also coordinates node and hashing functionality interplay

package OpenFaces

import (
	"encoding/json"		
	"net/http"

	"hashing"
	"node"
)

// These must be global to keep state between invocations
var nodes []OpenFacesNode.Node
var peerNetworkNodes []OpenFacesNode.Node

var Command struct {
   	Info OpenFacesNode.Node `json:"info"`
}

func enableCors(w *http.ResponseWriter) {
      (*w).Header().Set("Access-Control-Allow-Origin", "https://open-faces-8877b.firebaseapp.com")
  		(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func newRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", create)
	router.HandleFunc("/create", create)
	router.HandleFunc("/mine", mineNode)
	router.HandleFunc("/peerNodes", getPeerNodes)
	router.HandleFunc("/update", updateNode)
	return router
}

func BuildBlockchain(response http.ResponseWriter, request *http.Request) {
	enableCors(&response)
	json.NewDecoder(request.Body).Decode(&Command)

	router := newRouter()
	router.ServeHTTP(response, request)
}

func create(response http.ResponseWriter, request *http.Request) {
// Clear state from previous application invocations, and initiate the node
// creation workflow.


	enableCors(&response)
	
	// Google Cloud Platforms may (or may not) hold on to state in its execution 
	// environment. Let's explicitly clear it on startup. 
	nodes = []OpenFacesNode.Node{}
	peerNetworkNodes = []OpenFacesNode.Node{}

	
	jsNodes := encodeResponse(response, nodes, "create")
	response.Write([]byte(jsNodes))
}

func mineNode(response http.ResponseWriter, request *http.Request) {
// Prep a new input node, and if applicable add it to the blockchain.
// We only mine the node at index len(nodes) on the blockchain
// If a node at index < len(nodes) has been altered and not mined, this has no
// effect on the subsequent node(s).


	enableCors(&response)
	json.NewDecoder(request.Body).Decode(&Command)
	
	if Command.Info.Idx == 0 || (Command.Info.Idx > 0 && nodes[Command.Info.Idx - 1].IsMined) {
		if Command.Info.Idx >= len(nodes) {
			hashOfPrev := getHashOfPrevious(len(nodes))
			nonce, hashKey := OpenFacesHashing.GenerateHashAndNonce(len(nodes), Command.Info.Data, hashOfPrev)

			node := OpenFacesNode.ConstructNode(len(nodes), Command.Info.Data, hashKey, nonce, hashOfPrev, true, true) 
			nodes = append(nodes,node)
			
			peerNetworkNodes = append(peerNetworkNodes, node)
		} else {	
			hashOfPrev := getHashOfPrevious(Command.Info.Idx)
			nonce, hashKey := OpenFacesHashing.GenerateHashAndNonce(len(nodes), Command.Info.Data, hashOfPrev)
	 		node := OpenFacesNode.ConstructNode(
				Command.Info.Idx,
				Command.Info.Data, hashKey,
      	nonce, hashOfPrev, true, true)

	 		nodes = append(nodes[:node.Idx], append([]OpenFacesNode.Node{node}, nodes[node.Idx+1:]...)...)
  	}
	}

	jsNodes := encodeResponse(response, nodes, "update")
	response.Write([]byte(jsNodes))
}

func getPeerNodes(response http.ResponseWriter, request *http.Request) {
  enableCors(&response)

	jsNodes := encodeResponse(response, peerNetworkNodes, "get")
	response.Write([]byte(jsNodes))
}

func updateNode(response http.ResponseWriter, request *http.Request) {
// Determining three possibilites of the node in the request:
// * Nothing in the reqeust differs from what's already present => don't update
// * The node is the last in blackchain => update with no invalidation
// * The node is earlier than the last node => update with invalidation


	enableCors(&response)
	node := OpenFacesNode.Node{}
	isValidBlockchain := true

	json.NewDecoder(request.Body).Decode(&Command)
	hashOfPrev := getHashOfPrevious(Command.Info.Idx)
  newHashKey := OpenFacesHashing.GenerateHash(Command.Info.Idx, Command.Info.Data, hashOfPrev)

	// If nothing changes we don't need to do anything, so reasign previous values
	if Command.Info.Hash == nodes[Command.Info.Idx].Hash {
		node = OpenFacesNode.ConstructNode(
			Command.Info.Idx,
  		Command.Info.Data, Command.Info.Hash, 
			Command.Info.Nonce, Command.Info.HashOfPrev, 
			Command.Info.IsValid, Command.Info.IsMined) 		
	}

	if Command.Info.Idx == len(nodes)-1 {		
		if Command.Info.Hash == nodes[Command.Info.Idx].Hash {
  		node = OpenFacesNode.ConstructNode(
					Command.Info.Idx,
  				Command.Info.Data, newHashKey, 
					Command.Info.Nonce, hashOfPrev, true, Command.Info.IsMined) 		
		}
	} 

	if Command.Info.Idx < len(nodes)-1 {
  		node = OpenFacesNode.ConstructNode(
				Command.Info.Idx,
				Command.Info.Data, newHashKey, 
				0, Command.Info.HashOfPrev, false, Command.Info.IsMined) 
				isValidBlockchain = false
	}

	nodes = append(nodes[:node.Idx], append([]OpenFacesNode.Node{node}, nodes[node.Idx+1:]...)...)
	
	// Update the isValid and isMined annotation
	if !isValidBlockchain {
		for i := Command.Info.Idx; i < len(nodes); i++ {
			nodes[i].IsValid = false
			nodes[i].IsMined = false
		}
	}

	jsNodes := encodeResponse(response, nodes, "update")
	response.Write([]byte(jsNodes))
}



// -------------------------- Auxiliary functions -----------------------------


func getHashOfPrevious(posOfCurrNode int) string {
	if posOfCurrNode == 0 {
  	 return OpenFacesHashing.GetGenesisHash() 
	} else {
		return nodes[posOfCurrNode-1].Hash
	}
}

func encodeResponse(w http.ResponseWriter, nodes []OpenFacesNode.Node, caller string) ([]byte) {
// Encodes json for http response. Throw error on failure	

	jsNodes, err := json.Marshal(nodes)
	if err != nil {
    http.Error(w, "Can't process " + caller, http.StatusInternalServerError)
		return []byte{}
	}

	return jsNodes
}

