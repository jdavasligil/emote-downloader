// TODO:
// - Create universal interface for all emotes
// - Asynchronously make requests to each provider
// - Obtain emote image data as well?
// - Provide a stream (iterator) for digesting emotes

package emotedownloader

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"slices"
)

// Example URL to access BTTV cdn for image:
// https://cdn.betterttv.net/emote/54fa8f1401e468494b85b537/1x.webp
var (
	bttvAPIVersion = "3"

	EmoteProviders = []string{
		"BTTV",
		"7TV",
		"FFZ",
	}
)

// Used to unmarshal errors from an API response
type jsonError struct {
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// BTTV format for now.
type Emote struct {
	Provider  string
	ID        string `json:"id"`
	Code      string `json:"code"`
	ImageType string `json:"imageType"`
	Animated  bool   `json:"animated"`
	UserID    string `json:"userId"`
}

type EmoteDownloader struct {
	Emotes []Emote
}

func (ed *EmoteDownloader) Download() error {
	var err error

	emotesChan := make(chan []Emote, 4)
	errorChan := make(chan error, 4)

	done := make(chan struct{}, 1)

	go func() {
		bttvEmotes, err := getBTTVGlobalEmotes()
		emotesChan <- bttvEmotes
		errorChan <- err

		close(done)
	}()

	<-done

	close(emotesChan)
	close(errorChan)

	for e := range errorChan {
		err = errors.Join(e)
	}

	for emotes := range emotesChan {
		ed.Emotes = slices.Concat(emotes)
	}

	return err
}

func getBTTVGlobalEmotes() ([]Emote, error) {
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   "api.betterttv.net",
			Path:   "/" + bttvAPIVersion + "/cached/emotes/global",
		},
		Header: http.Header{},
	}

	//req.Header.Set("Client-ID", b.ClientID)

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		errorMessage := &jsonError{}
		err = json.Unmarshal(body, errorMessage)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(errorMessage.Error.Message)
	}

	bttvEmotes := []Emote{}
	err = json.Unmarshal(body, &bttvEmotes)
	if err != nil {
		return nil, err
	}

	for i, _ := range bttvEmotes {
		bttvEmotes[i].Provider = "BTTV"
	}

	return bttvEmotes, nil
}
