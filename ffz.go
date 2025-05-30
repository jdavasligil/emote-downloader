package emodl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"unsafe"

	"github.com/mailru/easyjson"
)

// DOCUMENTATION
// https://api.frankerfacez.com/docs/

var (
	ffzAPIVersion = "v1"
	ffzHost       = "api.frankerfacez.com"
	//ffzUserPathTmpl, _ = template.New("ffzUserPath").Parse("/{{ .Version }}/cached/users/{{ .Platform }}/{{ .PlatformID }}")
)

type FFZOptions struct {
	// Platform linked to FFZ (twitch, youtube)
	Platform string

	// ID associated with Platform (not username)
	PlatformID string
}

type FFZEmote struct {
	ID     int               `json:"id"`
	Name   string            `json:"name"`
	Height int               `json:"height"`
	Width  int               `json:"width"`
	URLs   map[string]string `json:"urls"`
}

type FFZEmoteSet struct {
	ID     int        `json:"id"`
	Emotes []FFZEmote `json:"emoticons"`
}

//easyjson:json
type FFZEmoteSetResponse struct {
	DefaultSets []int                  `json:"default_sets"`
	Sets        map[string]FFZEmoteSet `json:"sets"`
}

//easyjson:json
type FFZRoomData struct {
	Room struct {
		Set int `json:"set"`
	} `json:"room"`
	Sets map[string]FFZEmoteSet `json:"sets"`
}

// Scale "1", "2", etc. for 1x, 2x, etc. scales.
func (e FFZEmote) URL(scale string) string {
	if url, ok := e.URLs[scale]; ok {
		return url
	}

	if url, ok := e.URLs["1"]; ok {
		return url
	} else if url, ok := e.URLs["2"]; ok {
		return url
	} else if url, ok := e.URLs["3"]; ok {
		return url
	} else if url, ok := e.URLs["4"]; ok {
		return url
	}

	return ""
}

// Scale "1", "2", etc. for 1x, 2x, etc. scales.
// FFZ only returns .png images.
func (e FFZEmote) Image(scale string) Image {
	return Image{
		URL:    e.URL(scale),
		Width:  e.Width,
		Height: e.Height,
		ID:     strconv.Itoa(e.ID),
	}
}

func (e FFZEmote) AsEmote() Emote {
	return Emote{
		ID:        strconv.Itoa(e.ID),
		Name:      e.Name,
		Images:    []Image{e.Image("1")},
		Locations: []string{},
	}
}

func (e FFZEmote) Size() uintptr {
	var size uintptr
	size = unsafe.Sizeof(e)
	// Not a great approximation, but it works for now.
	for k, v := range e.URLs {
		size += unsafe.Sizeof(k) + unsafe.Sizeof(v)
	}
	return size
}

// Get User -> User Emote Sets -> User Emotes
func getFFZRoomEmoteSet(platform string, platformID string) (FFZEmoteSet, error) {
	var ffzRoomData FFZRoomData
	var ffzEmoteSet FFZEmoteSet
	var path string
	sb := strings.Builder{}
	platform = strings.ToLower(platform)
	if platform == "youtube" || platform == "yt" {
		path = "room/yt"
	} else if platform == "twitch" || platform == "ttv" {
		path = "room/id"
	}
	err := apiPathOptionTmpl.Execute(&sb, apiPath{
		Version: ffzAPIVersion,
		Path:    path,
		Option:  platformID,
	})
	if err != nil {
		return ffzEmoteSet, err
	}
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   ffzHost,
			Path:   sb.String(),
		},
		Header: http.Header{},
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return ffzEmoteSet, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMessage := &jsonError{}
		err = easyjson.UnmarshalFromReader(response.Body, errorMessage)
		if err != nil {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return ffzEmoteSet, err
			}
			return ffzEmoteSet, errors.New(response.Status + "\n" + string(body))
		}
		return ffzEmoteSet, errors.New(errorMessage.Error.Message)
	}
	// DEBUG
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return ffzEmoteSet, err
	// }
	// fmt.Printf("BODY:\n%s\n", body)

	err = easyjson.UnmarshalFromReader(response.Body, &ffzRoomData)
	if err != nil {
		return ffzEmoteSet, err
	}

	strIdx := strconv.Itoa(ffzRoomData.Room.Set)
	if strIdx == "" {
		return ffzEmoteSet, errors.New(fmt.Sprintf("failed to convert set %d to string", ffzRoomData.Room.Set))
	}

	ffzEmoteSet, ok := ffzRoomData.Sets[strIdx]
	if !ok {
		return ffzEmoteSet, errors.New(fmt.Sprintf("FFZ Emote Set ID=%s does not exist", strIdx))
	}

	return ffzEmoteSet, nil
}

func getFFZEmoteSets(setID string) ([]FFZEmoteSet, error) {
	var ffzEmoteSetResponse FFZEmoteSetResponse
	var ffzEmoteSets []FFZEmoteSet
	sb := strings.Builder{}
	err := apiPathOptionTmpl.Execute(&sb, apiPath{
		Version: ffzAPIVersion,
		Path:    "set",
		Option:  setID,
	})
	if err != nil {
		return ffzEmoteSets, err
	}

	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Scheme: "https",
			Host:   ffzHost,
			Path:   sb.String(),
		},
		Header: http.Header{},
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return ffzEmoteSets, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errorMessage := &jsonError{}
		err = easyjson.UnmarshalFromReader(response.Body, errorMessage)
		if err != nil {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				return ffzEmoteSets, err
			}
			return ffzEmoteSets, errors.New(response.Status + "\n" + string(body))
		}
		return ffzEmoteSets, errors.New(errorMessage.Error.Message)
	}
	// DEBUG
	// body, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	return ffzEmoteSets, err
	// }
	// fmt.Printf("BODY:\n%s\n", body)

	err = easyjson.UnmarshalFromReader(response.Body, &ffzEmoteSetResponse)
	if err != nil {
		return ffzEmoteSets, err
	}

	for _, idx := range ffzEmoteSetResponse.DefaultSets {
		strIdx := strconv.Itoa(idx)
		if strIdx == "" {
			return ffzEmoteSets, errors.New(fmt.Sprintf("failed to convert index %d to string", idx))
		}
		if set, ok := ffzEmoteSetResponse.Sets[strIdx]; ok {
			ffzEmoteSets = append(ffzEmoteSets, set)
		} else {
			return ffzEmoteSets, errors.New(fmt.Sprintf("FFZ Emote Set ID=%s does not exist", strIdx))
		}
	}

	return ffzEmoteSets, nil
}
