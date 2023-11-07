package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
	"time"
)

const userUrl = "/admin/realms/master/users"
const groupUrl = "/admin/realms/master/groups"
const membersUrl = "/admin/realms/master/groups/%s/members"

type Client struct {
	accessToken string
	httpClient  *http.Client
	baseUrl     string
}

type User struct {
	Id                         string        `json:"id"`
	CreatedTimestamp           int64         `json:"createdTimestamp"`
	Username                   string        `json:"username"`
	Enabled                    bool          `json:"enabled"`
	Totp                       bool          `json:"totp"`
	EmailVerified              bool          `json:"emailVerified"`
	DisableableCredentialTypes []interface{} `json:"disableableCredentialTypes"`
	RequiredActions            []interface{} `json:"requiredActions"`
	NotBefore                  int           `json:"notBefore"`
	Access                     struct {
		ManageGroupMembership bool `json:"manageGroupMembership"`
		View                  bool `json:"view"`
		MapRoles              bool `json:"mapRoles"`
		Impersonate           bool `json:"impersonate"`
		Manage                bool `json:"manage"`
	} `json:"access"`
}

type Group struct {
	Id        string        `json:"id"`
	Name      string        `json:"name"`
	Path      string        `json:"path"`
	SubGroups []interface{} `json:"subGroups"`
}

func New(accessToken string, baseUrl string) *Client {
	return &Client{
		accessToken: accessToken,
		baseUrl:     baseUrl,
		httpClient:  &http.Client{Timeout: 10 * time.Second}}
}

func (c *Client) ListUsers(ctx context.Context) ([]*User, error) {
	var ret []*User

	req, err := http.NewRequest(http.MethodGet, "http://"+path.Join(c.baseUrl, userUrl), nil)
	if err != nil {
		return nil, err
	}

	userBytes, err := c.do(ctx, req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(userBytes).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) ListGroups(ctx context.Context) ([]*Group, error) {
	var ret []*Group

	req, err := http.NewRequest(http.MethodGet, "http://"+path.Join(c.baseUrl, groupUrl), nil)
	if err != nil {
		return nil, err
	}

	groupBytes, err := c.do(ctx, req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(groupBytes).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) ListGroupMembers(ctx context.Context, groupId string) ([]*User, error) {
	var ret []*User

	req, err := http.NewRequest(http.MethodGet, "http://"+path.Join(c.baseUrl, fmt.Sprintf(membersUrl, groupId)), nil)
	if err != nil {
		return nil, err
	}

	userBytes, err := c.do(ctx, req)

	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(userBytes).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *Client) do(ctx context.Context, req *http.Request) (io.Reader, error) {
	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	response, err := c.httpClient.Do(req.WithContext(ctx))

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	ret, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	fmt.Printf("%s \n", ret)

	return bytes.NewBuffer(ret), nil
}
