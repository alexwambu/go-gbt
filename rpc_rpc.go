package rpc

import (
    "encoding/json"
    "net/http"
    "github.com/alexwambu/go-gbt/core"
)

type RPCServer struct {
    TxPool *core.TxPool
}

func (s *RPCServer) Serve(addr string) {
    http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
        decoder := json.NewDecoder(r.Body)
        var req map[string]interface{}
        decoder.Decode(&req)

        switch req["method"] {
        case "eth_sendRawTransaction":
            // For now just accept dummy tx
            tx := &core.Transaction{From: "0xabc", To: "0xdef", Value: 1}
            s.TxPool.Add(tx)
            w.Write([]byte(`{"jsonrpc":"2.0","result":"0xtxhash123","id":1}`))
        case "eth_blockNumber":
            w.Write([]byte(`{"jsonrpc":"2.0","result":"0x1","id":1}`))
        default:
            w.Write([]byte(`{"jsonrpc":"2.0","error":"method not found","id":1}`))
        }
    })

    http.ListenAndServe(addr, nil)
}
