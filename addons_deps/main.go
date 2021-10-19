package addons_deps

import (
	"errors"
	"strings"
)

const FROM_KEY_Value = "FROM"
const FROM_KEY_Errors = ""

const ENV_KEY = "ENV"
const ENV_VALUE_ZERO = "CGO=ENABLED=0"
const ENV_VALUE_ONE = "CGO=ENABLED=1"
const ENV_VALUE = ""

const ENV_PORT_ZERO = ":0000"
const ENV_PORT_DNSMasq = ":53"
const ENV_PORT_SYSTEM = ":44347"
const ENV_PORT_cup = ":631"

var Error_Tag = errors.New("this image is not available on dockerhub")
var CGOErrorValue = errors.New("alien code detected")
var Locked_Port = errors.New("port had already been locked")

func From(str string) (string, string, error) {

	if strings.Contains(str, "golang:1.17.2-bullseye") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-buster") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-stretch") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-alpine3.14") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-alpine3.13") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-windowsservercore-ltsc2022") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-windowsservercore-1809") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-windowsservercore-ltsc2016") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-nanoserver-ltsc2022") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.17.2-nanoserver-1809") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-bullseye") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-buster") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-stretch") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-alpine3.14") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-alpine3.13") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-windowsservercore-ltsc2022") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-windowsservercore-1809") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-windowsservercore-ltsc2016") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-nanoserver-ltsc2022") {
		return str, FROM_KEY_Value, nil
	} else if strings.Contains(str, "golang:1.16.9-nanoserver-1809") {
		return str, FROM_KEY_Value, nil
	} else {
		return "", FROM_KEY_Errors, Error_Tag
	}
}

func CGO(str string) (string, string, error) {
	if strings.Contains(str, "0") {
		return ENV_KEY, ENV_VALUE_ZERO, nil
	} else if strings.Contains(str, "1") {
		return ENV_KEY, ENV_VALUE_ONE, nil
	} else {
		return ENV_KEY, ENV_VALUE, CGOErrorValue
	}
}

func PORTMAP(str string) (string, string, error) {
	if strings.Contains(str, ENV_PORT_ZERO) {
		return ENV_KEY, ENV_PORT_ZERO, Locked_Port
	} else if strings.Contains(str, ENV_PORT_DNSMasq) {
		return ENV_KEY, ENV_PORT_DNSMasq, Locked_Port
	} else if strings.Contains(str, ENV_PORT_SYSTEM) {
		return ENV_KEY, ENV_PORT_SYSTEM, Locked_Port
	} else if strings.Contains(str, ENV_PORT_cup) {
		return ENV_KEY, ENV_PORT_cup, Locked_Port
	} else {
		return ENV_KEY, str, nil
	}
}
