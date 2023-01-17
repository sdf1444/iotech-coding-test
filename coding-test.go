package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
)

type Device struct {
	Name      string `json:"Name"`
	Type      string `json:"Type"`
	Info      string `json:"Info"`
	Value     string `json:"value"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Devices []Device `json:"Devices"`
}

type Output struct {
	ValueTotal int      `json:"ValueTotal"`
	UUIDS      []string `json:"UUIDS"`
}

func main() {
	data := Data{}
	output := Output{}
	output.UUIDS = make([]string, 0)
	fmt.Println("Parsing data from data.json")

	// Parse data from data.json
	file, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading data.json", err)
		return
	}
	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling data:", err)
		return
	}

	// Discard devices where timestamp is before current time
	now := time.Now().Unix()
	for _, device := range data.Devices {
		timestamp, err := strconv.ParseInt(device.Timestamp, 10, 64)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			continue
		}
		if timestamp > now {
			// Extract uuid from Info field
			uuidRegex := regexp.MustCompile("[a-f0-9-]{36}")
			uuid := uuidRegex.FindString(device.Info)
			output.UUIDS = append(output.UUIDS, uuid)

			value, err := base64.StdEncoding.DecodeString(device.Value)
			if err != nil {
				fmt.Println("Error decoding base64 value:", err)
				continue
			}
			valueInt, err := strconv.Atoi(string(value))
			if err != nil {
				fmt.Println("Error converting value to int:", err)
				continue
			}
			output.ValueTotal += valueInt
		}
	}

	// Marshal output struct to JSON
	outputJSON, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling output:", err)
		return
	}

	// Write output to file
	err = ioutil.WriteFile("output.json", outputJSON, 0644)
	if err != nil {
		fmt.Println("Error writing output to file:", err)
		return
	}

	fmt.Println("Successfully parsed data and wrote output to output.json")
	fmt.Println("Output: ")
	fmt.Println(string(outputJSON))
}
