package testx

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func Marshal(obj interface{}) string {
	return marshalObj(obj, json.Marshal)
}

func MarshalIndent(obj interface{}) string {
	return marshalObj(obj, func(o interface{}) ([]byte, error) { return json.MarshalIndent(o, "", "  ") })
}

func PrintObj(msg string, obj interface{}, writer ...io.Writer) {
	var w io.Writer = os.Stdout
	if len(writer) > 0 && writer[0] != nil {
		w = writer[0]
	}
	_, _ = fmt.Fprintf(w, "%s   (Type: %T)\n%s\n", msg, obj, MarshalIndent(obj))
}

func marshalObj(obj interface{}, handler func(interface{}) ([]byte, error)) string {
	data, err := handler(obj)
	if err != nil {
		return err.Error()
	}
	return string(data)
}
