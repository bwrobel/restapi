package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
    "io/ioutil"
    "log"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusOK)

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func TodoCreate(w http.ResponseWriter, r *http.Request){
    var todo Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    s := string(body[:])
    log.Printf(s);

    if err := json.Unmarshal(body, &todo); err != nil {

        if ute, ok := err.(*json.UnmarshalTypeError); ok {
            fmt.Printf("UnmarshalTypeError %v - %v - %v\n", ute.Value, ute.Type, ute.Offset)
        } else {
            fmt.Println("Other error:", err)
        }

        w.Header().Set("Content-Type", "application/json;charset=UTF=8")
        w.WriteHeader(422) //unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
        return
    }

    log.Printf(todo.Name);
    t := RepoCreateTodo(todo)
    w.Header().Set("Content-Type", "application/json;charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }
}

