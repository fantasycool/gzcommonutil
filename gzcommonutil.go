package gzcommonutil

import (
	"errors"
	"io"
	"log"
	"net/http"
	"reflect"
)

func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}
	return false, errors.New("not in array")
}

func ValidateIfGzFile(reader io.Reader) (bool, error) {
	buff := make([]byte, 512)
	log.Printf("start to fill 512 bytes to buff!")
	_, err := reader.Read(buff)

	if err != nil {
		log.Printf("GzipReader read failed! %s \n", err)
		return false, err
	}
	fileType := http.DetectContentType(buff)
	switch fileType {
	case "application/x-gzip", "application/zip":
		log.Printf("fileType is %s \n", fileType)
		return true, nil
	default:
		log.Printf("file is not compressed \n")
		return false, nil
	}
}
