// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package emotedownloader

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson2d7cdb3fDecodeGithubComJdavasligilEmoteDownloader(in *jlexer.Lexer, out *SevenTVEmoteCollection) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "emotes":
			if in.IsNull() {
				in.Skip()
				out.Emotes = nil
			} else {
				in.Delim('[')
				if out.Emotes == nil {
					if !in.IsDelim(']') {
						out.Emotes = make([]struct {
							Data SevenTVEmote `json:"data"`
						}, 0, 0)
					} else {
						out.Emotes = []struct {
							Data SevenTVEmote `json:"data"`
						}{}
					}
				} else {
					out.Emotes = (out.Emotes)[:0]
				}
				for !in.IsDelim(']') {
					var v1 struct {
						Data SevenTVEmote `json:"data"`
					}
					easyjson2d7cdb3fDecode(in, &v1)
					out.Emotes = append(out.Emotes, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d7cdb3fEncodeGithubComJdavasligilEmoteDownloader(out *jwriter.Writer, in SevenTVEmoteCollection) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"emotes\":"
		out.RawString(prefix)
		if in.Emotes == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Emotes {
				if v2 > 0 {
					out.RawByte(',')
				}
				easyjson2d7cdb3fEncode(out, v3)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v SevenTVEmoteCollection) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson2d7cdb3fEncodeGithubComJdavasligilEmoteDownloader(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v SevenTVEmoteCollection) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson2d7cdb3fEncodeGithubComJdavasligilEmoteDownloader(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *SevenTVEmoteCollection) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson2d7cdb3fDecodeGithubComJdavasligilEmoteDownloader(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *SevenTVEmoteCollection) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson2d7cdb3fDecodeGithubComJdavasligilEmoteDownloader(l, v)
}
func easyjson2d7cdb3fDecode(in *jlexer.Lexer, out *struct {
	Data SevenTVEmote `json:"data"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "data":
			easyjson2d7cdb3fDecodeGithubComJdavasligilEmoteDownloader1(in, &out.Data)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d7cdb3fEncode(out *jwriter.Writer, in struct {
	Data SevenTVEmote `json:"data"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"data\":"
		out.RawString(prefix[1:])
		easyjson2d7cdb3fEncodeGithubComJdavasligilEmoteDownloader1(out, in.Data)
	}
	out.RawByte('}')
}
func easyjson2d7cdb3fDecodeGithubComJdavasligilEmoteDownloader1(in *jlexer.Lexer, out *SevenTVEmote) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.Id = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "animated":
			out.Animated = bool(in.Bool())
		case "host":
			easyjson2d7cdb3fDecode1(in, &out.Host)
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d7cdb3fEncodeGithubComJdavasligilEmoteDownloader1(out *jwriter.Writer, in SevenTVEmote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.Id))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"animated\":"
		out.RawString(prefix)
		out.Bool(bool(in.Animated))
	}
	{
		const prefix string = ",\"host\":"
		out.RawString(prefix)
		easyjson2d7cdb3fEncode1(out, in.Host)
	}
	out.RawByte('}')
}
func easyjson2d7cdb3fDecode1(in *jlexer.Lexer, out *struct {
	Url   string `json:"url,intern"`
	Files []struct {
		Name       string `json:"name"`
		StaticName string `json:"static_name"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		FrameCount int    `json:"frame_count"`
		Size       uint32 `json:"size"`
		Format     string `json:"format"`
	} `json:"files"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.Url = string(in.StringIntern())
		case "files":
			if in.IsNull() {
				in.Skip()
				out.Files = nil
			} else {
				in.Delim('[')
				if out.Files == nil {
					if !in.IsDelim(']') {
						out.Files = make([]struct {
							Name       string `json:"name"`
							StaticName string `json:"static_name"`
							Width      int    `json:"width"`
							Height     int    `json:"height"`
							FrameCount int    `json:"frame_count"`
							Size       uint32 `json:"size"`
							Format     string `json:"format"`
						}, 0, 0)
					} else {
						out.Files = []struct {
							Name       string `json:"name"`
							StaticName string `json:"static_name"`
							Width      int    `json:"width"`
							Height     int    `json:"height"`
							FrameCount int    `json:"frame_count"`
							Size       uint32 `json:"size"`
							Format     string `json:"format"`
						}{}
					}
				} else {
					out.Files = (out.Files)[:0]
				}
				for !in.IsDelim(']') {
					var v4 struct {
						Name       string `json:"name"`
						StaticName string `json:"static_name"`
						Width      int    `json:"width"`
						Height     int    `json:"height"`
						FrameCount int    `json:"frame_count"`
						Size       uint32 `json:"size"`
						Format     string `json:"format"`
					}
					easyjson2d7cdb3fDecode2(in, &v4)
					out.Files = append(out.Files, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d7cdb3fEncode1(out *jwriter.Writer, in struct {
	Url   string `json:"url,intern"`
	Files []struct {
		Name       string `json:"name"`
		StaticName string `json:"static_name"`
		Width      int    `json:"width"`
		Height     int    `json:"height"`
		FrameCount int    `json:"frame_count"`
		Size       uint32 `json:"size"`
		Format     string `json:"format"`
	} `json:"files"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.Url))
	}
	{
		const prefix string = ",\"files\":"
		out.RawString(prefix)
		if in.Files == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Files {
				if v5 > 0 {
					out.RawByte(',')
				}
				easyjson2d7cdb3fEncode2(out, v6)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson2d7cdb3fDecode2(in *jlexer.Lexer, out *struct {
	Name       string `json:"name"`
	StaticName string `json:"static_name"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	FrameCount int    `json:"frame_count"`
	Size       uint32 `json:"size"`
	Format     string `json:"format"`
}) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "name":
			out.Name = string(in.String())
		case "static_name":
			out.StaticName = string(in.String())
		case "width":
			out.Width = int(in.Int())
		case "height":
			out.Height = int(in.Int())
		case "frame_count":
			out.FrameCount = int(in.Int())
		case "size":
			out.Size = uint32(in.Uint32())
		case "format":
			out.Format = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson2d7cdb3fEncode2(out *jwriter.Writer, in struct {
	Name       string `json:"name"`
	StaticName string `json:"static_name"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	FrameCount int    `json:"frame_count"`
	Size       uint32 `json:"size"`
	Format     string `json:"format"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix[1:])
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"static_name\":"
		out.RawString(prefix)
		out.String(string(in.StaticName))
	}
	{
		const prefix string = ",\"width\":"
		out.RawString(prefix)
		out.Int(int(in.Width))
	}
	{
		const prefix string = ",\"height\":"
		out.RawString(prefix)
		out.Int(int(in.Height))
	}
	{
		const prefix string = ",\"frame_count\":"
		out.RawString(prefix)
		out.Int(int(in.FrameCount))
	}
	{
		const prefix string = ",\"size\":"
		out.RawString(prefix)
		out.Uint32(uint32(in.Size))
	}
	{
		const prefix string = ",\"format\":"
		out.RawString(prefix)
		out.String(string(in.Format))
	}
	out.RawByte('}')
}
