package main

import (
	"errors"
	"fmt"
	"strconv"
)

func getTLVType(str string) (string, error) {
	t, err := getTLVValue(str, 3)
	if err != nil {
		return "", err
	}
	if t[0] != 'A' && t[0] != 'N' {
		return "", errors.New("type invalid")
	}
	n := t[1:]
	_, err = strconv.Atoi(n)
	if err != nil {
		return "", errors.New("type invalid number")
	}
	return t, nil
}

func getTLVLength(str string) (string, error) {
	l, err := getTLVValue(str, 2)
	if err != nil {
		return "", err
	}
	_, err = strconv.Atoi(l)
	if err != nil {
		return "", errors.New("length must be number")
	}
	return l, nil
}

func getTLVValue(str string, size int) (string, error) {
	if len(str) < size {
		return "", errors.New(fmt.Sprintf("string require min %d chars", size))
	}
	subStr := str[:size]
	return subStr, nil
}

func processTLVStr(str string) ([]map[string]string, error) {
	if str == "" {
		return nil, errors.New("str is empty")
	}

	var tlvMap []map[string]string

	var displaceStr = func(size int) {
		str = str[size:]
	}

	for len(str) > 0 {
		t, err := getTLVType(str)
		if err != nil {
			return nil, err
		}
		displaceStr(len(t))

		l, err := getTLVLength(str)
		if err != nil {
			return nil, err
		}
		displaceStr(len(l))

		size, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}

		v, err := getTLVValue(str, size)
		if err != nil {
			return nil, err
		}
		displaceStr(len(v))

		tlvMap = append(tlvMap, map[string]string{
			"type":   t,
			"length": l,
			"value":  v,
		})
	}

	return tlvMap, nil
}

// https://github.com/falabella-fif-inte/test-1
func main() {
	/*value, err := processTLVStr("A0511AB398765UJ1N230200")
	fmt.Println(fmt.Sprintf("map: %+v error: %+v", value, err))*/
}