package main

import (
	"encoding/json"
	"errors"
)

type LayoutMode string
type ItemType string

const (
	HorizontalLayout LayoutMode = "HORIZONTAL"
	VerticalLayout   LayoutMode = "VERTICAL"

	ComponentType ItemType = "COMPONENT"
	FrameType     ItemType = "FRAME"
	TextType      ItemType = "TEXT"
)

type Color struct {
	Red   float64 `json:"r"`
	Green float64 `json:"g"`
	Blue  float64 `json:"b"`
	Alpha float64 `json:"a"`
}

type Fill struct {
	Color Color `json:"color"`
}

type Style struct {
	FontWeight   float64 `json:"fontWeight"`
	FontSize     float64 `json:"fontSize"`
	LineHeightPx float64 `json:"lineHeightPx"`
}

type AbsoluteBoundingBox struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type Document struct {
	Name                string              `json:"name"`
	Children            []Document          `json:"children"`
	Type                ItemType            `json:"type"`
	Characters          string              `json:"characters"`
	LayoutMode          LayoutMode          `json:"layoutMode"`
	AbsoluteBoundingBox AbsoluteBoundingBox `json:"absoluteBoundingBox"`
	Style               Style               `json:"style"`
	PaddingLeft         float64             `json:"paddingLeft"`
	PaddingRight        float64             `json:"paddingRight"`
	PaddingTop          float64             `json:"paddingTop"`
	PaddingBottom       float64             `json:"paddingBottom"`
	CornerRadius        float64             `json:"cornerRadius"`
	ItemsSpacing        float64             `json:"itemSpacing"`
	BackgroundColor     Color               `json:"backgroundColor"`
	Fills               []Fill              `json:"fills"`
}

type Node struct {
	Document Document `json:"document"`
}

type Component struct {
	Name  string          `json:"name"`
	Nodes map[string]Node `json:"nodes"`
}

// UnmarshalJSON validates the layoutMode value during JSON unmarshaling
func (s *LayoutMode) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	candidate := LayoutMode(temp)
	if !candidate.IsValid() {
		return errors.New("Invalid layout mode")
	}

	*s = candidate
	return nil
}

func (s LayoutMode) IsValid() bool {
	switch s {
	case HorizontalLayout, VerticalLayout:
		return true
	}
	return false
}

// UnmarshalJSON validates the item type value during JSON unmarshaling
func (s *ItemType) UnmarshalJSON(data []byte) error {
	var temp string
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	candidate := ItemType(temp)
	if !candidate.IsValid() {
		return errors.New("Invalid item type")
	}

	*s = candidate
	return nil
}

func (s ItemType) IsValid() bool {
	switch s {
	case ComponentType, FrameType, TextType:
		return true
	}
	return false
}
