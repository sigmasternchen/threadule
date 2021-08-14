package web

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"threadule/backend/internal/app"
	"threadule/backend/internal/data/models"
)

type SessionInfo struct {
	User *models.User
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	AppCtx   *app.Context
	Session  SessionInfo
}

func (c *Context) ReadJSON(v interface{}) error {
	content, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(content, v)
}

func (c *Context) WriteJSON(v interface{}) error {
	content, err := json.Marshal(v)
	if err != nil {
		return err
	}

	_, err = c.Response.Write(content)
	return err
}
