package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func SendRequest(method string, url string, headers []string, body string, pretty bool) {

	if url == "" {
		fmt.Println("Error: --url is required")
	}

	req, err := http.NewRequest(method, url, bytes.NewBufferString(body))

	if err != nil {
		fmt.Println("Request creation failed:", err)
	}

	for _, h := range headers {
		var key, value string
		n, _ := fmt.Sscanf(h, "%s:%s", &key, &value)

		if n == 2 {
			req.Header.Set(key, value)
		}
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("Status: %v \n", resp.Status)
	respBody, _ := io.ReadAll(resp.Body)

	if pretty {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, respBody, "", " ")
		if err != nil {
			fmt.Println("Failed to pretty print:", err)
			fmt.Println(string(respBody))
			return
		}
		fmt.Println(prettyJSON.String())
	} else {
		fmt.Println(string(respBody))
	}

}
