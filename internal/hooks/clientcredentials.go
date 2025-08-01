// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package hooks

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/criblio/cribl-control-plane-sdk-go/internal/config"
	"github.com/criblio/cribl-control-plane-sdk-go/models/components"
	"golang.org/x/sync/singleflight"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

type session struct {
	Credentials *credentials
	Token       string
	ExpiresAt   *int64
	Scopes      []string
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   *int64 `json:"expires_in"`
}

type credentials struct {
	ClientID             string
	ClientSecret         string
	TokenURL             string
	AdditionalProperties map[string]string
}

type clientCredentialsHook struct {
	client   HTTPClient
	sessions sync.Map

	// sessionsGroup prevents concurrent token refreshes.
	sessionsGroup *singleflight.Group
}

var (
	_ sdkInitHook       = (*clientCredentialsHook)(nil)
	_ beforeRequestHook = (*clientCredentialsHook)(nil)
	_ afterErrorHook    = (*clientCredentialsHook)(nil)
)

func NewClientCredentialsHook() *clientCredentialsHook {
	return &clientCredentialsHook{
		sessionsGroup: new(singleflight.Group),
	}
}

func (c *clientCredentialsHook) SDKInit(config config.SDKConfiguration) config.SDKConfiguration {
	c.client = config.Client
	return config
}

func (c *clientCredentialsHook) BeforeRequest(ctx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if ctx.OAuth2Scopes == nil {
		// OAuth2 not in use
		return req, nil
	}

	credentials, err := c.getCredentials(ctx.HookContext, ctx.SecuritySource)
	if err != nil {
		return nil, &FailEarly{Cause: err}
	}
	if credentials == nil {
		return req, err
	}

	session, err := c.getSession(ctx, credentials)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", session.Token))

	return req, nil
}

func (c *clientCredentialsHook) AfterError(ctx AfterErrorContext, res *http.Response, err error) (*http.Response, error) {
	if ctx.OAuth2Scopes == nil {
		// OAuth2 not in use
		return res, err
	}

	// We don't want to refresh the token if the error is not related to the token
	if err != nil {
		return res, err
	}

	credentials, err := c.getCredentials(ctx.HookContext, ctx.SecuritySource)
	if err != nil {
		return nil, &FailEarly{Cause: err}
	}
	if credentials == nil {
		return res, err
	}

	if res != nil && res.StatusCode == http.StatusUnauthorized {
		sessionKey := getSessionKey(credentials.ClientID, credentials.ClientSecret)
		c.sessionsGroup.Forget(sessionKey)
		c.sessions.Delete(sessionKey)
	}

	return res, err
}

func (c *clientCredentialsHook) doTokenRequest(ctx HookContext, credentials *credentials, scopes []string) (*session, error) {
	values := url.Values{}
	values.Set("grant_type", "client_credentials")
	values.Set("client_id", credentials.ClientID)
	values.Set("client_secret", credentials.ClientSecret)

	if len(scopes) > 0 {
		values.Set("scope", strings.Join(scopes, " "))
	}

	for key, value := range credentials.AdditionalProperties {
		values.Set(key, value)
	}

	tokenURL := credentials.TokenURL
	u, err := url.Parse(tokenURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse token URL: %w", err)
	}
	if !u.IsAbs() {
		tokenURL, err = url.JoinPath(ctx.BaseURL, tokenURL)
		if err != nil {
			return nil, fmt.Errorf("failed to parse token URL: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx.Context, http.MethodPost, tokenURL, bytes.NewBufferString(values.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send token request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("unexpected status code: %d: %s", res.StatusCode, body)
	}

	var tokenRes tokenResponse
	if err := json.NewDecoder(res.Body).Decode(&tokenRes); err != nil {
		return nil, fmt.Errorf("failed to decode token response: %w", err)
	}

	if tokenRes.TokenType != "Bearer" {
		return nil, fmt.Errorf("unexpected token type: %s", tokenRes.TokenType)
	}

	var expiresAt *int64
	if tokenRes.ExpiresIn != nil {
		expiresAt = new(int64)
		*expiresAt = time.Now().Unix() + *tokenRes.ExpiresIn
	}

	return &session{
		Credentials: credentials,
		Token:       tokenRes.AccessToken,
		ExpiresAt:   expiresAt,
		Scopes:      scopes,
	}, nil
}

func (c *clientCredentialsHook) getCredentials(ctx HookContext, source func(ctx context.Context) (any, error)) (*credentials, error) {
	if source == nil {
		return nil, nil
	}

	sec, err := source(ctx.Context)
	if err != nil {
		return nil, err
	}

	return c.getCredentialsGlobal(sec)
}

func (c *clientCredentialsHook) getCredentialsGlobal(sec any) (*credentials, error) {
	security, ok := sec.(components.Security)

	if !ok {
		return nil, fmt.Errorf("unexpected security type: %T", sec)
	}

	if security.ClientOauth == nil {
		return nil, nil
	}
	secType := reflect.TypeOf(security.ClientOauth)
	if secType.Kind() == reflect.Ptr {
		secType = secType.Elem()
	}
	secValue := reflect.ValueOf(security.ClientOauth)
	if secValue.Kind() == reflect.Ptr {
		secValue = secValue.Elem()
	}
	if security.ClientOauth.TokenURL == "" {

		tokenURLField, ok := secType.FieldByName("TokenURL")
		if !ok {
			return nil, fmt.Errorf("TokenURL is required for security type %s", secType.Name())
		}
		tokenURLDefault := tokenURLField.Tag.Get("default")
		security.ClientOauth.TokenURL = tokenURLDefault
	}

	additionalProperties := make(map[string]string)
	for i := 0; i < secType.NumField(); i++ {
		field := secType.Field(i)
		if field.Name != "TokenURL" && field.Name != "ClientID" && field.Name != "ClientSecret" {
			// Get the field value using reflection
			fieldValue := secValue.Field(i)
			if fieldValue.IsValid() {
				tag := field.Tag.Get("security")
				parts := strings.Split(tag, ",")
				for _, part := range parts {
					if strings.HasPrefix(part, "name=") {
						additionalProperties[strings.TrimPrefix(part, "name=")] = fieldValue.String()
						break
					}
				}
			}
		}
	}

	return &credentials{
		ClientID:             security.ClientOauth.ClientID,
		ClientSecret:         security.ClientOauth.ClientSecret,
		TokenURL:             security.ClientOauth.TokenURL,
		AdditionalProperties: additionalProperties,
	}, nil
}

func (c *clientCredentialsHook) getSession(ctx BeforeRequestContext, credentials *credentials) (*session, error) {
	sessionKey := getSessionKey(credentials.ClientID, credentials.ClientSecret)

	var cachedSession *session

	rawCachedSession, ok := c.sessions.Load(sessionKey)

	if ok {
		cachedSession = rawCachedSession.(*session)

		if !hasRequiredScopes(cachedSession.Scopes, ctx.OAuth2Scopes) || hasTokenExpired(cachedSession.ExpiresAt) {
			c.sessionsGroup.Forget(sessionKey)
			c.sessions.Delete(sessionKey)
		}
	}

	rawSession, err, _ := c.sessionsGroup.Do(sessionKey, func() (any, error) {
		refreshedSession, err := c.doTokenRequest(ctx.HookContext, credentials, getScopes(ctx.OAuth2Scopes, cachedSession))

		if err != nil {
			return nil, fmt.Errorf("failed to get token: %w", err)
		}

		c.sessions.Store(sessionKey, refreshedSession)

		return refreshedSession, err
	})

	if err != nil {
		return nil, err
	}

	session := rawSession.(*session)

	return session, nil
}

func getSessionKey(clientID, clientSecret string) string {
	key := fmt.Sprintf("%s:%s", clientID, clientSecret)
	hash := md5.Sum([]byte(key))
	return hex.EncodeToString(hash[:])
}

func hasRequiredScopes(scopes []string, requiredScopes []string) bool {
	for _, requiredScope := range requiredScopes {
		found := false
		for _, scope := range scopes {
			if scope == requiredScope {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func getScopes(requiredScopes []string, sess *session) []string {
	scopes := requiredScopes
	if sess != nil {
		for _, scope := range sess.Scopes {
			found := false
			for _, requiredScope := range requiredScopes {
				if scope == requiredScope {
					found = true
					break
				}
			}
			if !found {
				scopes = append(scopes, scope)
			}
		}
	}

	return scopes
}

func hasTokenExpired(expiresAt *int64) bool {
	return expiresAt == nil || time.Now().Unix()+60 >= *expiresAt
}
