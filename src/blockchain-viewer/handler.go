package main

import (
	"blockchain-viewer/blockchain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// NewServeMux create a new mux with request handlers
func newServerMux() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	r.HandleFunc("/api/chain/show/{id:[0-9]+}", http.HandlerFunc(handleChainShow)).Methods("GET")
	r.HandleFunc("/api/chain/create", http.HandlerFunc(handleChainCreate)).Methods("POST")
	r.HandleFunc("/api/chain/list", http.HandlerFunc(handleChainList)).Methods("GET")
	r.HandleFunc("/api/chain/add-block/{id:[0-9]+}", http.HandlerFunc(handleAddBlock)).Methods("POST")

	return r
}

// handleChainShow return a chain with a given id
func handleChainShow(rw http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
	}

	chain, ok := blockchain.DB.Get(uint64(idNum))
	if !ok {
		http.NotFound(rw, r)
		return
	}

	b, _ := json.Marshal(chain)
	jsonResponse(rw, b)
}

// createChainHandler handle requests of creating new chain
func handleChainCreate(rw http.ResponseWriter, r *http.Request) {
	chain, err := parseRequest(r, "chain")
	if err != nil {
		msg := fmt.Sprintf("error while parsing request: %s", err)
		http.Error(rw, msg, http.StatusBadRequest)
		return
	}

	c := chain.(*blockchain.Chain)
	chainId := blockchain.DB.Add(c)
	resp := struct {
		ID uint64 `json:"id"`
	}{
		chainId,
	}

	b, _ := json.Marshal(resp)
	jsonResponse(rw, b)
}

// handleChainList return the lists of created chains
func handleChainList(rw http.ResponseWriter, r *http.Request) {
	chains := blockchain.DB.GetAll()
	resp := struct {
		Chains []*blockchain.Chain `json:"chains"`
	}{
		chains,
	}

	b, err := json.Marshal(resp)
	if err != nil {
		msg := fmt.Sprintf("error while parsing chain list to JSON: %s", err)
		http.Error(rw, msg, http.StatusInternalServerError)
	}

	jsonResponse(rw, b)
}

// handleAddBlock add a new block to the chain with an given id
func handleAddBlock(rw http.ResponseWriter, r *http.Request) {
	// TODO remove code duplication from handleChainShow
	id := mux.Vars(r)["id"]
	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(rw, "Invalid ID", http.StatusBadRequest)
	}

	chain, ok := blockchain.DB.Get(uint64(idNum))
	if !ok {
		http.NotFound(rw, r)
		return
	}

	chain.AddBlock()
	b, _ := json.Marshal(chain)
	jsonResponse(rw, b)
}

// parseRequest expects the request body to be JSON, it parses the body to the concrete struct
func parseRequest(r *http.Request, dataType string) (interface{}, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading request body: %s", err)
	}
	defer r.Body.Close()

	if len(body) == 0 {
		return nil, fmt.Errorf("request body is empty")
	}

	var data interface{}
	switch dataType {
	case "chain":
		data = &blockchain.Chain{}
	}

	if err := json.Unmarshal(body, data); err != nil {
		return nil, fmt.Errorf("failed to parse request to json: %s", err)
	}

	return data, nil
}

// jsonResponse add json content header and write data to the the response
func jsonResponse(rw http.ResponseWriter, b []byte) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(b)
}
