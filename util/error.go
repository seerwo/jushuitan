package util

import (
	"encoding/json"
	"fmt"
	"reflect"
)


// CommonError return error json
type CommonError struct {
	ErrCode string  `json:"code"`
	ErrMsg  string `json:"message"`
}

// DecodeWithCommonError  return with error
func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commError CommonError
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return
	}
	if commError.ErrCode != "0000" {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, commError.ErrCode, commError.ErrMsg)
	}
	return nil
}

// DecodeWithError return error
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, errCode.Int(), errMsg.String())
	}
	return nil
}

