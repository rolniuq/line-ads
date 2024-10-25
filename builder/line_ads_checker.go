package lineads

import (
	"encoding/json"
)

type lineAdsError struct {
	Errors []struct {
		Reason  LineAdsError `json:"reason"`
		Message string       `json:"message"`
	} `json:"errors"`
}

func (s *LineAdsRequest[T]) getError(body []byte) error {
	var data lineAdsError
	if err := json.Unmarshal(body, &data); err != nil {
		return err
	}
	if len(data.Errors) == 0 {
		return nil
	}

	reason := LineAdsErrorReasons.Find(data.Errors[0].Reason)
	if reason == nil {
		return nil
	}

	return nil
}
