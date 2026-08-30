package main

import (
	"gopkg.in/yaml.v3"
	"strings"
	"path/filepath"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func loadYaml(path string) ([]byte, error) {

	ext := strings.ToLower(filepath.Ext(path));
	if ext != ".yaml" && ext != ".yml" {
		return nil,fmt.Errorf("invalid file extension %q: must be .yaml or .yml", ext);
	}

	content, err := os.ReadFile(path);
	
	if err!=nil {
		return nil, fmt.Errorf("failed to read file: %w",err);
	}

	var node yaml.Node;
	if err:= yaml.Unmarshal(content, &node); err!=nil {
		return nil, fmt.Errorf("invalid YAML content in %s: %w",path, err);
	}

	return content, nil;
}

func main() {
	content, err := loadYaml("./test.yaml");
	if err!=nil {
		log.Fatalf("error: %v", err);
	}

	endpoint := "http://localhost:8080/yaml"

	body := bytes.NewReader(content)

	req, err := http.NewRequest(http.MethodPost, endpoint, body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	req.Header.Set("Content-Type", "application/yaml")

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("post: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		msg, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		log.Fatalf("bad status: %s: %s", resp.Status, string(msg))
	}

	if _, err := io.Copy(io.Discard, resp.Body); err != nil {
		log.Printf("drain: %v", err)
	}

	fmt.Printf("POST ok: %s\n", resp.Status)
}
