import Handlebars from "handlebars";
import {
    isComponentNode,
    isRectangleNode,
    isTextNode,
    parseColor,
} from "./renderUtils";
import { Rectangle, Node } from "@figma/rest-api-spec";
import fs from "fs";
import path from "path";

function loadTemplate(path: string): Handlebars.TemplateDelegate {
    const templateSource = fs.readFileSync(path, "utf8");
    return Handlebars.compile(templateSource);
}

const renderRectangle = loadTemplate(
    path.join(__dirname, "templates/rectangle.hbs"),
);
const renderText = loadTemplate(path.join(__dirname, "templates/text.hbs"));
const renderItem = loadTemplate(path.join(__dirname, "templates/item.hbs"));

export function render(node: Node): string {
    if (isComponentNode(node) || isRectangleNode(node) || isTextNode(node)) {
        return renderNode(node, node.absoluteBoundingBox);
    }
    return "";
}

function renderNode(node: Node, scene: Rectangle | null): string {
    let x = 0;
    let y = 0;
    let width = 0;
    let height = 0;

    if (isTextNode(node) || isComponentNode(node) || isRectangleNode(node)) {
        if (node.absoluteBoundingBox) {
            x = node.absoluteBoundingBox.x;
            y = node.absoluteBoundingBox.y;
            width = node.absoluteBoundingBox.width;
            height = node.absoluteBoundingBox.height;
            if (scene) {
                x -= scene.x;
                y -= scene.y;
            }
        }
    }

    const theta = node.rotation || 0;
    const cosA = Math.abs(Math.cos(theta));
    const cos2A = Math.abs(Math.cos(2 * theta));
    const sinA = Math.abs(Math.sin(theta));
    let originalWidth = width;
    let originalHeight = height;
    if (cos2A !== 0) {
        originalWidth = (width * cosA - height * sinA) / cos2A;
        originalHeight = (height * cosA - width * sinA) / cos2A;
    }

    let borderColor = "";
    let borderWidth = 0;
    let color = "";
    let opacity = "";
    if (isTextNode(node) || isComponentNode(node) || isRectangleNode(node)) {
        for (const stroke of node.strokes || []) {
            if (stroke.type === "SOLID") {
                borderColor = parseColor(stroke.color);
                if (node.strokeWeight) {
                    borderWidth = node.strokeWeight;
                }
                break;
            }
        }

        for (const fill of node.fills) {
            if (fill.type === "SOLID") {
                color = parseColor(fill.color);
                if (fill.opacity !== undefined) {
                    opacity = fill.opacity.toString();
                }
                break;
            }
        }
    }

    let child = "";
    if (isRectangleNode(node)) {
        child = renderRectangle({ color, opacity });
    } else if (isTextNode(node)) {
        child = renderText({
            color,
            opacity,
            characters: node.characters,
            style: node.style,
            calcBaselineOffset: (lineHeight: number, fontSize: number) =>
                lineHeight - fontSize,
        });
    } else if (isComponentNode(node)) {
        child = node.children
            .map((child) => renderNode(child, scene).toString())
            .join("\n");
    }

    const result = renderItem({
        x,
        y,
        width,
        height,
        originalHeight,
        originalWidth,
        rotationDeg: (theta * 180) / Math.PI,
        borderColor,
        borderWidth,
        color,
        opacity,
        child,
        clip: isComponentNode(node),
    });
    return result;
}
