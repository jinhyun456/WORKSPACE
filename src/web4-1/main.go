package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) { //파일업로드 랜들러 정의
	//---------------------------------상대방이 전송해준 파일을 리딩-----//
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close() //핸들러자원을 사용후 defer로 닫아주어야

	//--- 새로운 uploads Form파일을 만들고 저장할 공간----------------------------//
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777) //8진수 형태, MKdirAll(redwer Write 모두) 리눅스명령
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename)
	file, err := os.Create(filepath) //핸들은 OS자원이므로 사용후 defer로 닫아주어야
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return

	}
	//-------파일을 copy하는 공간----------------//
	io.Copy(file, uploadFile) // 위 uploads된 파일을 아래 비어있는 파일에 copy 한다
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath) //filepath는 upload경로를 알려준다
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler) //파일업로드 핸들러를 만들어준다
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
