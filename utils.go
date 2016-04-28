package runscope

import (
	"encoding/json"
	"fmt"
)

func unmarshal(content []byte, result interface{}) error {
	var response = Response{Data: result}
	if err := json.Unmarshal(content, &response); err != nil {
		return err
	}

	if response.Error.Message != "" {
		return fmt.Errorf("%s (%d): %+v", response.Error.Message, response.Error.Status, response.Meta)
	}
	return nil
}
