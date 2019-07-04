package main

import (
	"net/http"
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"io"
	"time"
	"strconv"
	"bytes"
	"compress/gzip"
	"strings"
	"github.com/xxdstem/bancho/handlers"
	"github.com/xxdstem/bancho/common"
	"github.com/jmoiron/sqlx"
)
const ProtocolVersion = 19

var cnf Config


type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

type Config struct {
	DSN    		string      	`json:"dsn"` 
	Port		int				`json:"port"`
}

type ConnectionHandler struct{}

func (c ConnectionHandler) serveHTTPReal(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()
	w.Header().Add("Content-Type", "text/html; charset=utf-8")

	if r.Method != "POST" || r.UserAgent() != "osu!" {
		return
	}

	// Get data from request body
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error while attempting to understand request:", err)
		return
	}

	// We're not using .Add() because it capitalizes the string automatically. We'd rather not.
	w.Header()["cho-protocol"] = []string{strconv.Itoa(ProtocolVersion)}
	w.Header().Add("Vary", "Accept-Encoding")

	// Handle the packet
	buf := new(bytes.Buffer)
	newToken, err := handlers.Handle(data, buf, r.Header.Get("osu-token"))
	if err != nil {
		fmt.Println("Error in bancho:", err)
	}

	// Finish it up.
	if newToken != "" {
		w.Header()["cho-token"] = []string{newToken}
	}
	io.Copy(w, buf)
	fmt.Printf("> Request end - time took: %s\n", time.Since(begin).String())
}

func (c ConnectionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		c.serveHTTPReal(w, r)
		return
	}
	// Set up the chunked transfer.
	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("expected http.ResponseWriter to be an http.Flusher")
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	c.serveHTTPReal(gzr, r)
	flusher.Flush()
}

func initConfig(){
	jsonFile, err := os.Open("config.json")
    // if we os.Open returns an error then handle it
    if err != nil {
		fmt.Println(err)
		return
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(byteValue, &cnf)
}

func main(){
	initConfig()
	common.Init()
	common.DB, _ = sqlx.Open("mysql", cnf.DSN)
	fmt.Println("connected to dB!")
	defer common.DB.Close()
	handler := &ConnectionHandler{}
	http.ListenAndServe(fmt.Sprintf(":%d", cnf.Port), handler)
}
