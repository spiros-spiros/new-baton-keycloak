package keycloak

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Clarilab/gocloaksession"
	"github.com/Nerzal/gocloak/v13"
)

type Client struct {
	client       *gocloak.GoCloak
	realm        string
	clientID     string
	clientSecret string
	session      gocloaksession.GoCloakSession
}

func NewClient(serverURL, realm, clientID, clientSecret string) (*Client, error) {
	session, err := gocloaksession.NewSession(clientID, clientSecret, realm, serverURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		client:       gocloak.NewClient(serverURL),
		realm:        realm,
		clientID:     clientID,
		clientSecret: clientSecret,
		session:      session,
	}, nil
}

func (c *Client) AddUserToGroup(ctx context.Context, userID, groupID string) error {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	return c.client.AddUserToGroup(ctx, token.AccessToken, c.realm, userID, groupID)
}

func (c *Client) RemoveUserFromGroup(ctx context.Context, userID, groupID string) error {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}

	return c.client.DeleteUserFromGroup(ctx, token.AccessToken, c.realm, userID, groupID)
}

func (c *Client) GetUsers(ctx context.Context, first int) ([]*gocloak.User, string, error) {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return nil, strconv.Itoa(first), fmt.Errorf("failed to get token: %w", err)
	}

	max := 300

	users, err := c.client.GetUsers(ctx, token.AccessToken, c.realm, gocloak.GetUsersParams{
		First: pointer(first),
		Max:   pointer(max),
	})
	if err != nil {
		return nil, strconv.Itoa(first), fmt.Errorf("failed to get users: %w", err)
	}

	if len(users) == 0 {
		return nil, "", nil
	}

	return users, strconv.Itoa(first + max), nil
}

func (c *Client) GetGroupMembers(ctx context.Context, groupID string) ([]*gocloak.User, error) {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	return c.client.GetGroupMembers(ctx, token.AccessToken, c.realm, groupID, gocloak.GetGroupsParams{})
}

func (c *Client) GetGroups(ctx context.Context, first int) ([]*gocloak.Group, string, error) {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return nil, strconv.Itoa(first), fmt.Errorf("failed to get token: %w", err)
	}

	max := 300

	groups, err := c.client.GetGroups(ctx, token.AccessToken, c.realm, gocloak.GetGroupsParams{
		First: pointer(first),
		Max:   pointer(max),
	})
	if err != nil {
		return nil, strconv.Itoa(first), fmt.Errorf("failed to get groups: %w", err)
	}

	if len(groups) == 0 {
		return nil, "", nil
	}

	return groups, strconv.Itoa(first + max), nil
}

func (c *Client) GetUserGroups(ctx context.Context, userID string) ([]*gocloak.Group, error) {
	token, err := c.session.GetKeycloakAuthToken()
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	return c.client.GetUserGroups(ctx, token.AccessToken, c.realm, userID, gocloak.GetGroupsParams{})
}

func (c *Client) Close() error {
	return nil
}

func pointer[T any](v T) *T {
	return &v
}
