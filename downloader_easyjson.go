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

func easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader(in *jlexer.Lexer, out *jsonError) {
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
		case "error":
			easyjson3fc20da9Decode(in, &out.Error)
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
func easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader(out *jwriter.Writer, in jsonError) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix[1:])
		easyjson3fc20da9Encode(out, in.Error)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v jsonError) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v jsonError) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *jsonError) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *jsonError) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader(l, v)
}
func easyjson3fc20da9Decode(in *jlexer.Lexer, out *struct {
	Message string `json:"message"`
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
		case "message":
			out.Message = string(in.String())
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
func easyjson3fc20da9Encode(out *jwriter.Writer, in struct {
	Message string `json:"message"`
}) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message\":"
		out.RawString(prefix[1:])
		out.String(string(in.Message))
	}
	out.RawByte('}')
}
func easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader1(in *jlexer.Lexer, out *BTTVEmotes) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(BTTVEmotes, 0, 0)
			} else {
				*out = BTTVEmotes{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 BTTVEmote
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader1(out *jwriter.Writer, in BTTVEmotes) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v BTTVEmotes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BTTVEmotes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BTTVEmotes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BTTVEmotes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader1(l, v)
}
func easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader2(in *jlexer.Lexer, out *BTTVEmote) {
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
		case "Provider":
			out.Provider = string(in.String())
		case "id":
			out.ID = string(in.String())
		case "code":
			out.Code = string(in.String())
		case "imageType":
			out.ImageType = string(in.StringIntern())
		case "animated":
			out.Animated = bool(in.Bool())
		case "userId":
			out.UserID = string(in.StringIntern())
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
func easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader2(out *jwriter.Writer, in BTTVEmote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"Provider\":"
		out.RawString(prefix[1:])
		out.String(string(in.Provider))
	}
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix)
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"code\":"
		out.RawString(prefix)
		out.String(string(in.Code))
	}
	{
		const prefix string = ",\"imageType\":"
		out.RawString(prefix)
		out.String(string(in.ImageType))
	}
	{
		const prefix string = ",\"animated\":"
		out.RawString(prefix)
		out.Bool(bool(in.Animated))
	}
	{
		const prefix string = ",\"userId\":"
		out.RawString(prefix)
		out.String(string(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BTTVEmote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BTTVEmote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3fc20da9EncodeGithubComJdavasligilEmoteDownloader2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BTTVEmote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BTTVEmote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3fc20da9DecodeGithubComJdavasligilEmoteDownloader2(l, v)
}
