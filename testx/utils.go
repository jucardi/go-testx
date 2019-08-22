package testx

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

func Marshal(obj interface{}) string {
	return marshalObj(obj, yaml.Marshal)
}

func PrintObj(msg string, obj interface{}, writer ...io.Writer) {
	var w io.Writer = os.Stdout
	if len(writer) > 0 && writer[0] != nil {
		w = writer[0]
	}
	_, _ = fmt.Fprintf(w, "%s   (Type: %T)\n%s\n", msg, obj, Marshal(obj))
}

func marshalObj(obj interface{}, handler func(interface{}) ([]byte, error)) string {
	data, err := handler(obj)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
