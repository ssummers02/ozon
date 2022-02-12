package handler

import (
	"encoding/json"
	"net/http"
	"ozon/pkg/restmodel"
)

func (s *Server) postLink(w http.ResponseWriter, r *http.Request) {
	var data restmodel.Link

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "json")
		return
	}

	if len(data.Link) == 0 {
		SendErr(w, http.StatusBadRequest, "json")
		return
	}

	result, err := s.services.Link.Create(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	response := LinkResponse{
		Link:      result.Link,
		ShortLink: result.ShortLink,
	}

	SendOK(w, http.StatusOK, response)
}
func (s *Server) getLink(w http.ResponseWriter, r *http.Request) {
	var data restmodel.Link

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, "json")
		return
	}

	if len(data.ShortLink) == 0 {
		SendErr(w, http.StatusBadRequest, "json")
		return
	}

	result, err := s.services.Link.GetByShortLink(data)
	if err != nil {
		SendErr(w, http.StatusBadRequest, err.Error())
		return
	}

	response := LinkResponse{
		Link:      result.Link,
		ShortLink: result.ShortLink,
	}

	SendOK(w, http.StatusOK, response)
}
