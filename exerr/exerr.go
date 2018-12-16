package exerr

import (
	"encoding/json"
	"fmt"
	"os"
)

//LEVEL : errorlevel
type LEVEL int

const (
	//Error : Error level error
	Error LEVEL = 0
	//Warn : Error level warn
	Warn LEVEL = 1
	//Info : Error level info
	Info LEVEL = 2
)

//ExtendedError : extendederror
type ExtendedError struct {
	Level           LEVEL  `json:"level"`
	Message         string `json:"message"`
	CustomerMessage string `json:"customer_message"`
	HTTPStatus      int    `json:"http_status"`
}

var (
	errorJSON map[string]*ExtendedError
)

const (
	//ErrFailedToDecodeErrorFile : err failed decoding configuration file
	ErrFailedToDecodeErrorFile = "Failed to decode configuration file: %v"

	//ErrFailedToOpenErrorFile : info using default configuration file path
	ErrFailedToOpenErrorFile = "Failed to open configuration file: %v"
)

//Initialize :the config
func Initialize(configurationFileAbsPath string) (err error) {
	errorJSONFile, err := os.Open(configurationFileAbsPath)
	defer errorJSONFile.Close()
	if err != nil {
		err = fmt.Errorf(ErrFailedToDecodeErrorFile, err)
		return
	}
	jsonParser := json.NewDecoder(errorJSONFile)
	err = jsonParser.Decode(&errorJSON)
	if err != nil {
		err = fmt.Errorf(ErrFailedToOpenErrorFile, err)
	}
	return
}

//NewExtendedError : Create a instance of extended error
func NewExtendedError(errorCode string, args ...interface{}) (eError *ExtendedError) {
	eError = errorJSON[errorCode]
	eError.Message = fmt.Sprintf(eError.Message, args)
	eError.CustomerMessage = fmt.Sprintf(eError.CustomerMessage, args)
	return eError
}

func (e *ExtendedError) Error() (errorString string) {
	return e.Message
}
