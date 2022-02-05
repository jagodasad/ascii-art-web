package main

import (
	//"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	//"regexp"
	//"strings"
	"testing"
	"web/art"
)

func ReadTestFile() string {
	data, err := ioutil.ReadFile("test_test_file.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)

	//contentSplit := strings.Split(contentString, "#")

	return contentString
}

func TestAscii_Art(t *testing.T) {

	req := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=Hello", nil)
	req.Header.Set("Content-Type", "text/html;")
	//if err != nil {
	//	t.Fatal("Could not create request: %v", err)
	//}

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	//fmt.Println(string(body))

	os.WriteFile("output.txt", []byte(body), 0o644)
	outputFile, _ := ioutil.ReadFile("output.txt")

	if ReadTestFile() != string(outputFile) {
		t.Fatal("incorrect output")

	}

}
