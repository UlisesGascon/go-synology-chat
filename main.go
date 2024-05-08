package main

import (
	"fmt"
)

type SynologyChat struct {
	baseUrl         string
	token           string
	ignoreSSLErrors bool
	channelsUrl     string
	chatbotUrl      string
	usersUrl        string
}

func New(baseUrl string, token string, ignoreSSLErrors bool) (*SynologyChat, error) {
	sc := &SynologyChat{
		baseUrl:         baseUrl,
		token:           token,
		ignoreSSLErrors: ignoreSSLErrors,
	}
	var err error
	sc.channelsUrl, err = GetUrlByMethod(sc.baseUrl, sc.token)("channel_list")
	if err != nil {
		return nil, err
	}
	sc.chatbotUrl, err = GetUrlByMethod(sc.baseUrl, sc.token)("chatbot")
	if err != nil {
		return nil, err
	}
	sc.usersUrl, err = GetUrlByMethod(sc.baseUrl, sc.token)("user_list")
	if err != nil {
		return nil, err
	}
	return sc, nil
}

func (sc *SynologyChat) GetUsers() (map[string]interface{}, error) {
	return MakeGetRequest(sc.usersUrl, sc.ignoreSSLErrors)()
}

func (sc *SynologyChat) GetChannels() (map[string]interface{}, error) {
	return MakeGetRequest(sc.channelsUrl, sc.ignoreSSLErrors)()
}

func (sc *SynologyChat) SendDirectMessage(userIDs []int, text string, file_url ...string) (map[string]interface{}, error) {
	payloadStr, err := GeneratePayloadContent(userIDs, text, file_url...)
	if err != nil {
		return nil, fmt.Errorf("failed to generate payload content: %w", err)
	}
	return MakePostRequest(sc.chatbotUrl, payloadStr, sc.ignoreSSLErrors)
}
