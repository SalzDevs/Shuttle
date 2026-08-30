package main

import (
	"fmt"
	"log"
	"gopkg.in/yaml.v3"
	"os"
)

func main() {
	content, err := os.ReadFile("./test.yaml");
	if err!=nil { 
		log.Fatalf("error:%v",err);
	}

	parsedData := make(map[string] any);
	

	err = yaml.Unmarshal([]byte(content), &parsedData);

	fmt.Printf("ParsedData: %v\n",parsedData);

	if err!=nil {
		log.Fatalf("error: %v",err);
	}

	for k,v := range parsedData {
		fmt.Printf("(k:%s) (v:%v)\n",k,v);	
	}


}
