package mp

// This library is a wrapper around a media player that plays YouTube playlists.

import (
	"errors"
	"time"
)

// these are defined by the YouTube API
type State int

const (
	STATE_STOPPED   State = 0
	STATE_PLAYING         = 1
	STATE_PAUSED          = 2
	STATE_BUFFERING       = 3
)

// PlayState defines the current state of the generic MediaPlayer.
// It is shared within the MediaPlayer and used as an access token as well:
// whoever holds a pointer to this structure may access it's members.
// That also means that a pointer to this struct should be cleared when starting
// a new goroutine.
type PlayState struct {
	Playlist          []string
	Index             int
	State             State
	Volume            int
	bufferingPosition time.Duration
	newVolume         bool // true if the Volume property must be reapplied to the player
}

// Video returns the current video, or an empty string if there is no current
// video.
func (ps *PlayState) Video() string {
	if len(ps.Playlist) == 0 {
		return ""
	}
	return ps.Playlist[ps.Index]
}

type PlaylistState struct {
	Playlist []string
	Index    int
	Position time.Duration
	State    State
}

type StateChange struct {
	State    State
	Position time.Duration
}

const INITIAL_VOLUME = 80

var PROPERTY_UNAVAILABLE = errors.New("media player: property unavailable")
