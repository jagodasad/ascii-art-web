package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"regexp"
	"strings"
	"testing"
	"web/art"
)

func ReadTestString() []string {
	data, err := ioutil.ReadFile("test-output.txt")
	if err != nil {
		panic(err)
	}

	contentString := string(data)
	contentSplit := strings.Split(contentString, "#")

	return contentSplit
}

func TestAscii_Art_End(t *testing.T) {
	/*	tt := []struct {
			req string
			res string
		}{
			{string{"%21%22%23"}, ReadTestString()[0]}, //test for space!"#
			{string{"OPEN"}, ReadTestString()[1]},
			{string{"xyz{|}~"}, ReadTestString()[2]}, //test for xyz{|}~
		}

		for _, tc := range tt {
			s := tc.req
			if s != tc.res {
				t.Error("Output does not equal expected result")
			}
		}
	*/
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=xyz{|}~", nil)
	req1.Header.Set("Content-Type", "text/html;")

	// req2 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=echo", nil)
	// req2.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)
	// art.Asciiart(rec, req2)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	// titles := r.FindAllString(actual, -1)
	titlesMaybe := r.FindStringSubmatch(actual)

	fmt.Println(titlesMaybe[1])
	// fmt.Println(ReadTestString()[1])

	if titlesMaybe[1] != ReadTestString()[2] {
		t.Fatal("still needs work")
	}
}

func TestAscii_Art_Middle(t *testing.T) {
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=OPEN", nil)
	req1.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	titlesMaybe := r.FindStringSubmatch(actual)

	fmt.Println(titlesMaybe[1])

	if titlesMaybe[1] != ReadTestString()[1] {
		t.Fatal("still needs work")
	}
}

func TestAscii_Art_Beginning(t *testing.T) {
	req1 := httptest.NewRequest("POST", "localhost:8070/ascii-art?banner=standard&text=+%21%22%23", nil)
	req1.Header.Set("Content-Type", "text/html;")

	rec := httptest.NewRecorder()
	art.Asciiart(rec, req1)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK: got %v", res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Could not read response: %v", err)
	}

	os.WriteFile("output.txt", []byte(body), 0o644)

	actual := string(body)

	r, _ := regexp.Compile("<pre name=\"outtext\" >(.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+\n.+)\\/pre>")

	titlesMaybe := r.FindStringSubmatch(actual)

	fmt.Println(titlesMaybe[1])

	if titlesMaybe[1] != ReadTestString()[0] {
		t.Fatal("still needs work")
	}
}
