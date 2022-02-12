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

	SendOK(w, http.StatusOK, result)
	/*if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if len(link.Link) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	result, err := s.services.Link.Create(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)*/

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

	SendOK(w, http.StatusOK, result)
	/*if err := c.BindJSON(&link); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if len(link.ShortLink) == 0 {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	result, err := s.services.Link.GetByShortLink(link)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)*/
}
