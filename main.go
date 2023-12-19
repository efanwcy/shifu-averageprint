package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func fetchData() (string, error) {
	// 执行 curl 命令获取数据
	cmdResult, err := execCommand("curl", "deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement")
	if err != nil {
		return "", err
	}
	return cmdResult, nil
}

func execCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error executing command: %v\n%s", err, errb.String())
	}
	return outb.String(), nil
}

func parseData(data string) ([]float64, error) {
	// 将获取的数据解析成 float64 切片
	strValues := strings.Fields(data)
	values := make([]float64, len(strValues))
	for i, strValue := range strValues {
		value, err := strconv.ParseFloat(strValue, 64)
		if err != nil {
			return nil, err
		}
		values[i] = value
	}
	return values, nil
}

func calculateAverage(values []float64) float64 {
	// 计算平均值
	sum := 0.0
	for _, value := range values {
		sum += value
	}
	average := sum / float64(len(values))
	return average
}

func main() {
	for {
		// 获取数据
		data, err := fetchData()
		if err != nil {
			log.Printf("Error fetching data: %v", err)
			continue
		}

		// 解析数据
		values, err := parseData(data)
		if err != nil {
			log.Printf("Error parsing data: %v", err)
			continue
		}

		// 计算平均值
		average := calculateAverage(values)

		// 打印平均值
		fmt.Printf("Average: %.2f\n", average)

		// 等待 5 分钟
		time.Sleep(5 * time.Minute)
	}
}
