package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// https://gist.github.com/ik5/d8ecde700972d4378d87

var (
	Info = LogTeal
	Warn = LogYellow
	Fata = LogRed
)

var (
	LogBlack   = Color("\033[1;30m%s\033[0m")
	LogRed     = Color("\033[1;31m%s\033[0m")
	LogGreen   = Color("\033[1;32m%s\033[0m")
	LogYellow  = Color("\033[1;33m%s\033[0m")
	LogPurple  = Color("\033[1;34m%s\033[0m")
	LogMagenta = Color("\033[1;35m%s\033[0m")
	LogTeal    = Color("\033[1;36m%s\033[0m")
	LogWhite   = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) {
	sprint := func(args ...interface{}) {
		log.Printf(colorString, fmt.Sprint(args...))
	}
	return sprint
}

// Read from request json payload to exact entity
func ReadFromRequest[T any](req *http.Request) (*T, error) {
	var target T
	rawBodyBytes, err := io.ReadAll(req.Body)
	// log.Println("byte payload : " + string(rawBodyBytes))
	defer closeResources(req.Body)
	if err != nil {
		return &target, fmt.Errorf("failed to register : %w", err)
	}
	if err := json.Unmarshal(rawBodyBytes, &target); err != nil {
		return &target, fmt.Errorf("failed to register : %w", err)
	}
	// log.Printf("unmarshalled: %+v\n", target)
	return &target, nil
}

func Marshall[T any](rw http.ResponseWriter, entity T) ([]byte, bool) {
	respBytes, err := json.Marshal(entity)
	if err != nil {
		http.Error(rw, fmt.Errorf("failed to register : %w", err).Error(), http.StatusInternalServerError)
		return nil, false
	}
	// log.Printf("marshalled: %+v\n", string(respBytes))
	return respBytes, true
}

func SafeResponseWrite(rw http.ResponseWriter, content []byte, status int) {
	rw.WriteHeader(status)
	n, err := rw.Write(content)
	if err != nil || n < len(content) {
		http.Error(rw, "unexpected error", http.StatusInternalServerError)
		return
	}

}

func closeResources(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		log.Fatal("unable to close request body")
	}
}
