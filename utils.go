package runscope

import (
	"encoding/json"
	"errors"
	"fmt"
)

func unmarshal(content []byte, result interface{}) error {
	var response = Response{Data: result}
	if err := json.Unmarshal(content, &response); err != nil {
		return err
	}
	if response.Error.Message != "" {
		message := fmt.Sprintf("%s: %+v", response.Error, response.Meta)
		return errors.New(message)
	}
	return nil
}
