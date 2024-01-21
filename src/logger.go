package src

import (
	"encoding/json"
	"fmt"
	"os"
)

func LogObject(obj interface{}) {
	fmt.Println("----------------------------------------------------------------")
	b, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(b))
	fmt.Println("----------------------------------------------------------------")
}
