package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type LoadingData struct {
	MainHeader string `json:"mainHeader"`
	SubHeader  string `json:"subHeader"`
	StartTime  string `json:"startTime"`
	Timestamp  int64  `json:"timestamp"`
}

type PauseData struct {
	HeaderText   string `json:"headerText"`
	DescribeText string `json:"describeText"`
}

func WriteJSON[T LoadingData | PauseData](c *gin.Context, filePath string) {
	var instance T

	if err := c.BindJSON(&instance); err != nil {
		fmt.Println("BindJSON error:", err)
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//结构体转字节切片
	jsonBytes, marshalErr := json.MarshalIndent(instance, "", " ")
	if marshalErr != nil {
		fmt.Println(marshalErr)
		return
	}

	writeErr := os.WriteFile(filePath, jsonBytes, 0666)
	if writeErr != nil {
		fmt.Println(writeErr)
		return
	}

	c.Status(http.StatusOK)
}

func ReadJSON[T LoadingData | PauseData](c *gin.Context, filePath string) {
	var instance T

	jsonBytes, readErr := os.ReadFile(filePath)
	if readErr != nil {
		fmt.Println(555)
		fmt.Println(readErr)
		return
	}

	unmarshalErr := json.Unmarshal(jsonBytes, &instance)
	if unmarshalErr != nil {
		fmt.Println(666)
		fmt.Println(unmarshalErr)
		return
	}

	c.JSON(http.StatusOK, instance)
}
