package service

import (
	"encoding/json"
	"fmt"
	"github.com/cyclops-top/auth-resource/src/context"
	"github.com/cyclops-top/auth-resource/src/model"
	"github.com/kataras/iris/v12"
	"io/ioutil"
	"net/http"
	"strings"
)

const headerAuthorization = "Authorization"
const authorizationBearerPrefix = "Bearer "

var client = http.DefaultClient

func Authorize(ctx context.Context,authorizeUrl string) (*model.User, error) {
	token := parseToken(ctx)
	if len(token) == 0 {
		return nil, nil
	}
	request, _ := http.NewRequest("GET", authorizeUrl, nil)
	request.Header.Add(headerAuthorization, authorizationBearerPrefix+token)
	resp, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("authorize request is error")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("auth request is error: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body fail", err)
		return nil, err
	}
	t := model.User{}
	err = json.Unmarshal(body, &t)
	if err != nil {
		fmt.Println("parse json error", err)
		return nil, err
	}
	return &t, nil
}

func parseToken(ctx iris.Context) string {
	authorization := ctx.GetHeader("Authorization")
	if len(authorization) == 0 || !strings.HasPrefix(authorization, authorizationBearerPrefix) {
		return ""
	}
	return authorization[len(authorizationBearerPrefix):]
}
