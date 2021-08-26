package myapp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	FirstName string    `json:"first_name"` //어노테이션)(Annotation) 설명을 붙히는것
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) { //
	fmt.Fprint(w, "Hello goBlock화이팅") // Fprint은 w값에 "Hello World" 값을 주어서 프린팅하라는 의미
}

type fooHandler struct{} //인스턴스 생성

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //serveHTTP 인터페이스 구현
	user := new(User)                           //POST 접속방식은 a new foo를 CREAED 반환한다 new를 사용한것
	err := json.NewDecoder(r.Body).Decode(user) // Decorde 스트리밍기반 데이터전송, JSON을 go value
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) //인턴페이스, JSON형식 결과값은 바이트 어레이
	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	log.Print(string(data)) //이거 추가해줘야 json 작동됨
	fmt.Fprint(w, string(data))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	mux.HandleFunc("/bar", barHandler)

	mux.Handle("/foo", &fooHandler{})
	return mux
}
