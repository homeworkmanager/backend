package class

//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//)
//
//type UpdateClass struct {
//	Name        string `json:"summary"`
//	Description string `json:"description"`
//	StartTime   string `json:"start"`
//	EndTime     string `json:"end"`
//	Location    string `json:"location"`
//}
//
//func (s *Service) Update() error {
//
//	resp, err := http.Get("http://127.0.0.1:8000/events")
//	if err != nil {
//		return err
//	}
//	defer func() error {
//		err := resp.Body.Close()
//		if err != nil {
//			return err
//		}
//		return nil
//	}()
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return err
//	}
//
//	var updateClasses []UpdateClass
//	if err := json.Unmarshal(body, &updateClasses); err != nil {
//		fmt.Print(err.Error())
//	}
//
//	return nil
//
//}
