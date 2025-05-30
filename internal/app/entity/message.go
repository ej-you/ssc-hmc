// Package entity contains all app entities.
package entity

import (
	"time"

	"sschmc/internal/pkg/text"
)

const (
	MaxDisplayedLines = 4 // max lines amount that can be displayed simultaneously

	_maxLineLen   = 16 // max len of line
	_separatorLen = 8  // len of separator between message header and content
)

// Log message model.
type Message struct {
	ID        string    `gorm:"primaryKey;autoIncrement;type:INT"`
	Level     string    `gorm:"not null;size:1"`
	Header    string    `gorm:"not null;size:30"`
	Content   string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"not null;type:TIMESTAMP"`

	FirstLine int      `gorm:"-"` // idx of first displayed line on the device (default: 0)
	Lines     []string `gorm:"-"` // a slice of lines formatted for display
}

func (Message) TableName() string {
	return "storage"
}

// Format create lines slice of message Text to display it on device as text lines.
func (m *Message) Format() {
	// join header and content to a single string
	fullText := m.Header + "\n\n" + m.Content
	m.Lines = text.Normalize(fullText, _maxLineLen)
}

// ScrollUp updates message FirstLine for scrolling up imitation.
func (m *Message) ScrollUp() {
	// extreme up position
	if m.FirstLine == 0 {
		return
	}
	m.FirstLine--
}

// ScrollDown updates message FirstLine for scrolling down imitation.
func (m *Message) ScrollDown() {
	// extreme down position
	if m.FirstLine >= len(m.Lines)-MaxDisplayedLines {
		return
	}
	m.FirstLine++
}

// MessageLevelCount is a subset of Message model fields for "levels count" query.
type MessageLevelCount struct {
	Level int
	Count int
}

// MessageWithLevel is a subset of Message model fields for "messages with level" query.
type MessageWithLevel struct {
	ID        string
	CreatedAt time.Time
}
