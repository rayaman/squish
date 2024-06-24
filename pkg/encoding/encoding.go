package encoding

import (
	"fmt"
	"strings"

	"github.com/fatih/structs"
)

type options struct {
	max int64
}

func GetOptions(obj any) options {
	fields := structs.Fields(obj)
	for _, field := range fields {
		tag := field.Tag("binary")
		tags := strings.Split(tag, ",")
		fmt.Println(tags)
		for _, val := range tags {
			fmt.Println(strings.Split(val, "="))
		}
	}
	return options{}
}

func Marshal(obj any, version int16) error {
	return nil
}

func Unmarshal(data []byte) (any, int16, error) {
	return nil, 0, nil
}
