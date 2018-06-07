package main

///@NOTE Shyft handler functions when endpoints are hit
import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	"github.com/ethereum/go-ethereum/shyftdb"
	"github.com/gorilla/mux"
	"bytes"
	"bufio"

	"github.com/ethereum/go-ethereum/rpc"

	"reflect"
)

type RWC struct {
	*bufio.ReadWriter
}

func (rwc *RWC) Close() error {
	return nil
}

// Error wraps RPC errors, which contain an error code in addition to the message.
type Error interface {
	Error() string  // returns the message
	ErrorCode() int // returns the code
}

// callback is a method callback which was registered in the server
type callback struct {
	rcvr        reflect.Value  // receiver of method
	method      reflect.Method // callback
	argTypes    []reflect.Type // input argument types
	hasCtx      bool           // method's first argument is a context (not included in argTypes)
	errPos      int            // err return idx, of -1 when method cannot return error
	isSubscribe bool           // indication if the callback is a subscription
}

// serverRequest is an incoming request
type serverRequest struct {
	id            interface{}
	svcname       string
	callb         *callback
	args          []reflect.Value
	isUnsubscribe bool
	err           Error
}

type rpcRequest struct {
	service  string
	method   string
	id       interface{}
	isPubSub bool
	params   interface{}
	err      Error // invalid batch element
}

// GetTransaction gets txs
func GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txHash := vars["txHash"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	getTxResponse := shyftdb.GetTransaction(blockExplorerDb, txHash)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getTxResponse)
}

// GetAllTransactions gets txs
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	txs := shyftdb.GetAllTransactions(blockExplorerDb)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}


	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, txs)
}

// GetAllTransactions gets txs
func GetAllTransactionsFromBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	txsFromBlock := shyftdb.GetAllTransactionsFromBlock(blockExplorerDb, blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, txsFromBlock)
}

func GetAllBlocksMinedByAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coinbase := vars["coinbase"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	blocksMined := shyftdb.GetAllBlocksMinedByAddress(blockExplorerDb, coinbase)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, blocksMined)
}

// GetAccount gets balance
func GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	getAccountBalance := shyftdb.GetAccount(blockExplorerDb, address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getAccountBalance)
}

// GetAccount gets balance
func GetAccountTxs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	getAccountTxs := shyftdb.GetAccountTxs(blockExplorerDb, address)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getAccountTxs)
}

// GetAllAccounts gets balances
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}
	allAccounts := shyftdb.GetAllAccounts(blockExplorerDb)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, allAccounts)
}

//GetBlock returns block json
func GetBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNumber := vars["blockNumber"]
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}

	getBlockResponse := shyftdb.GetBlock(blockExplorerDb, blockNumber)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, getBlockResponse)
}

// GetAllBlocks response
func GetAllBlocks(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}
	block3 := shyftdb.GetAllBlocks(blockExplorerDb)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, block3)
}

func GetRecentBlock(w http.ResponseWriter, r *http.Request) {
	connStr := "user=postgres dbname=shyftdb sslmode=disable"
	blockExplorerDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return
	}
	mostRecentBlock := shyftdb.GetRecentBlock(blockExplorerDb)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, mostRecentBlock)

}

//GetInternalTransactions gets internal txs
func GetInternalTransactions(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Get InternalTransactions", address)
}

//GetInternalTransactionsHash gets internal txs hash
func GetInternalTransactionsHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionHash := vars["transaction_hash"]

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Get Internal Transaction Hash", transactionHash)
}

func BroadcastTx(w http.ResponseWriter, r *http.Request) {
	//{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x82B7ec1bfC2a8D5799B723AA6248c1BfE3eAc0cc", "latest"],"id":
	vars := mux.Vars(r)
	transactionHash := vars["transaction_hash"]
	fmt.Println("THE tx hash is")
	fmt.Println(transactionHash)

	req := bytes.NewBufferString(`{"jsonrpc":"2.0","method":"eth_getBalance","params":["0x82B7ec1bfC2a8D5799B723AA6248c1BfE3eAc0cc", "latest"],"id":1}`)

	var str string
	reply := bytes.NewBufferString(str)
	rw := &RWC{bufio.NewReadWriter(bufio.NewReader(req), bufio.NewWriter(reply))}

	codec := rpc.NewJSONCodec(rw)
	fmt.Println("the codec is")
	fmt.Println(codec)

	requests, _, _ := codec.ReadRequestHeaders()
	fmt.Println("THe requests are")
	fmt.Println(requests)
	fmt.Println(requests[0])

	//var requ *rpcRequest
	requ := requests[0]
	fmt.Println(requ)

	// since serverRequest is not exported, we don't have access to the
	// fields in serverRequest, even though that is the return type
	// of ReadRequestHeaders
	// fmt.Println(requ.id)

	//new_request := &serverRequest{id: requ.id, err: requ.err}

	//req := requests[0]
	//arguments := []reflect.Value{requests[0].callb.rcvr}
	//if requests[0].callb.hasCtx {
	//	arguments = append(arguments, reflect.ValueOf(ctx))
	//}
	//if len(req.args) > 0 {
	//	arguments = append(arguments, req.args...)
	//}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Get Transaction Hash", transactionHash)
}