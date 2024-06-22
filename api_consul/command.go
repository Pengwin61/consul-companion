package api_consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"consul-companion/internal/cfg"
)

func GetMembers(config cfg.Config) []ResponseMembers {
	var response []ResponseMembers

	url := fmt.Sprintf("%s://%s/v1/agent/members", config.ConsulScheme, config.ConsulAddress)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil
	}

	req.Header.Set("X-Consul-Token", config.ConsulToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil
	}
	return response
}

func GetNodeServices(config cfg.Config, nodeName string) Data {

	var response Data

	url := fmt.Sprintf("%s://%s/v1/catalog/node-services/%s", config.ConsulScheme, config.ConsulAddress, nodeName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("X-Consul-Token", config.ConsulToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
	}

	if len(response.Services) == 0 {
		fmt.Println("No services found for node:", nodeName)
	}

	return response
}

// func RegisterService(config cfg.Config) {
// }

func DeregisterService(config cfg.Config, nodeName string, serviceID string, nodeAddress string) {

	payload := map[string]string{
		"Node":      nodeName,
		"ServiceID": serviceID,
	}
	payloadData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON data:", err)
		return
	}

	url := fmt.Sprintf("%s://%s/v1/catalog/deregister", config.ConsulScheme, config.ConsulAddress)

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(payloadData))
	if err != nil {
		fmt.Println("Error creating request:", err)
	}

	req.Header.Set("X-Consul-Token", config.ConsulToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making GET request:", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
	}

	if string(body) == "true" {
		fmt.Printf("Service: %s Node Name: %s Node IP: %s deregistered successfully\n", serviceID, nodeName, nodeAddress)
	}
}
