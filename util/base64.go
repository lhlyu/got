package util

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"log"
)

func AnyEncode(v interface{}) string{
	if v == nil{
		log.Println("AnyEncode","v is nil")
		return ""
	}
	byts,err := json.Marshal(v)
	if err != nil{
		log.Println("AnyEncode","json marshal is err : ",err)
		return ""
	}
	return base64.StdEncoding.EncodeToString(byts)
}

func AnyDecode(s string,v interface{}) error{
	if v == nil{
		log.Println("AnyDecode","v is nil")
		return errors.New("v is nil")
	}
	byts,err := base64.URLEncoding.DecodeString(s)
	if err != nil{
		log.Println("AnyDecode","decodeString is err : ",err)
		return err
	}
	err = json.Unmarshal(byts,v)
	if err != nil{
		log.Println("AnyDecode","json unmarshal is err : ",err)
		return err
	}
	return nil
}
