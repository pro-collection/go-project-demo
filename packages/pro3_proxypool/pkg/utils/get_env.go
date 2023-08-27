package utils

import (
	"os"
	"strings"
)

type paramsType = map[string]string

func GetEnv(keyList ...string) (string, paramsType) {
	args := os.Args
	params := make(paramsType)

	if len(args) > 1 {
		for _, arg := range args[1:] {
			keyValue := strings.Split(arg, "=")
			if len(keyValue) == 2 {
				params[keyValue[0]] = keyValue[1]
			}
		}
	}

	if len(keyList) == 0 {
		return "", params
	}

	if len(keyList) == 1 {
		return params[keyList[0]], params
	}

	var confirmRunParams = make(paramsType)
	for _, key := range keyList {
		if params[key] != "" {
			confirmRunParams[key] = params[key]
		}
	}
	return "", params
}
