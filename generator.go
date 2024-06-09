package main

import (
	"fmt"
	"strings"
)

func parseColor(c Color) string {
	return fmt.Sprintf("Qt.rgba(%f, %f, %f, %f)", c.Red, c.Green, c.Blue, c.Alpha)
}

func generateText(el Document, level int) string {
	indent := strings.Repeat(" ", level*4)
	color := ""
	for _, f := range el.Fills {
		color = fmt.Sprintf("%s    color: %s", indent, parseColor(f.Color))
		break
	}
	return fmt.Sprintf(`
%sText {
%s    text: "%s"
%s
%s    font.weight: %f
%s    font.pixelSize: %f
%s    baselineOffset: %f
%s}`,
		indent, indent, el.Characters,
		color,
		indent, el.Style.FontWeight,
		indent, el.Style.FontSize,
		indent, el.Style.LineHeightPx/2-el.Style.FontSize/2,
		indent)
}

func generateComponent(el Document, level int) string {
	indent := strings.Repeat(" ", level*4)

	// Transform children
	var children []string
	for _, c := range el.Children {
		children = append(children, generate(c, level+2))
	}

	layout := "ColumnLayout"
	if el.LayoutMode == HorizontalLayout {
		layout = "RowLayout"
	}

	return fmt.Sprintf(`
%sRectangle {
%s    width: %f
%s    height: %f
%s    color: %s
%s    anchors.leftMargin: %f
%s    anchors.rightMargin: %f
%s    radius: %f
%s    %s {
%s        anchors.fill: parent
%s        Layout.preferredWidth: %f
%s        Layout.preferredHeight: %f
%s        spacing: %f
%s
%s    }
%s}`,
		indent, indent, el.AbsoluteBoundingBox.Width,
		indent, el.AbsoluteBoundingBox.Height,
		indent, parseColor(el.BackgroundColor),
		indent, el.PaddingLeft,
		indent, el.PaddingRight,
		indent, el.CornerRadius,
		indent, layout,
		indent, indent, el.AbsoluteBoundingBox.Width,
		indent, el.AbsoluteBoundingBox.Height,
		indent, el.ItemsSpacing,
		strings.Join(children, "\n"+indent),
		indent,
		indent,
	)
}

func generate(el Document, level int) string {
	switch el.Type {
	case ComponentType, FrameType:
		return generateComponent(el, level)
	case TextType:
		return generateText(el, level)
	}
	return ""
}

func GenerateQml(component Component) (string, error) {
	for _, value := range component.Nodes {
		doc := value.Document
		return fmt.Sprintf(`import QtQuick
import QtQuick.Layouts
%s `,
			generate(doc, 0)), nil
	}

	return "", nil
}
