package models

import (
	"fmt"
	"strings"
)

var allowedAddressTypes = []string{"A", "AAAA", "CNAME"}
var allowedCheckTypes = []string{"HTTP", "TCP", "UDP", "PING"}

func isInList(value string, list []string) bool {
	for _, item := range list {
		if strings.ToUpper(value) == item {
			return true
		}
	}
	return false
}

type Check struct {
	ID              string   `json:"id"`
	UserParent      string   `json:"user_parent"`
	ServiceName     string   `json:"service_name"`
	AddressType     string   `json:"address_type"`
	Addresses       []string `json:"addresses"`
	CheckType       string   `json:"check_type"`
	Port            int      `json:"port"`
	IntervalSeconds int      `json:"interval_seconds"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

func (check *Check) Validate() error {

	check.ID = ""

	check.AddressType = strings.ToUpper(check.AddressType)
	check.CheckType = strings.ToUpper(check.CheckType)

	if check.UserParent == "" {
		return fmt.Errorf("user_parent is required")
	}

	if check.ServiceName == "" {
		return fmt.Errorf("service_name is required")
	}

	if !isInList(check.AddressType, allowedAddressTypes) {
		return fmt.Errorf("address_type must be one of: A, AAAA, CNAME")
	}

	if len(check.Addresses) == 0 {
		return fmt.Errorf("at least one address is required")
	}

	if !isInList(check.CheckType, allowedCheckTypes) {
		return fmt.Errorf("check_type must be one of: HTTP, TCP, UDP, PING")
	}

	if check.CheckType == "TCP" || check.CheckType == "UDP" {
		if check.Port <= 0 || check.Port > 65535 {
			return fmt.Errorf("port must be between 1 and 65535 for TCP/UDP checks")
		}
	}

	if check.IntervalSeconds <= 0 {
		return fmt.Errorf("interval_seconds must be greater than 0")
	}

	return nil

}
