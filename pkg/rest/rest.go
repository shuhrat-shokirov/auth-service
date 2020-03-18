package rest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func ReadJSONBody(request *http.Request, dto interface{}) (err error) {
	if request.Header.Get("Content-Type") != "application/json" {
		return errors.New("error")
	}

	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return errors.New("error")
	}
	defer request.Body.Close()

	err = json.Unmarshal(body, &dto)
	if err != nil {
		return errors.New("error")
	}
	return nil
}

func WriteJSONBody(response http.ResponseWriter, dto interface{}) (err error) {
	response.Header().Set("Content-Type", "application/json")

	body, err := json.Marshal(dto)
	if err != nil {
		return errors.New("error")
	}

	_, err = response.Write(body)
	if err != nil {
		return errors.New("error")
	}

	return nil
}
