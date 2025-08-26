import {
    ComponentNode,
    Node,
    RectangleNode,
    TextNode,
    RGBA,
} from "@figma/rest-api-spec";

export function isComponentNode(node: Node): node is ComponentNode {
    return node.type === "COMPONENT";
}

export function isRectangleNode(node: Node): node is RectangleNode {
    return node.type === "RECTANGLE" || node.type === "VECTOR";
}

export function isTextNode(node: Node): node is TextNode {
    return node.type === "TEXT";
}

export function parseColor(c: RGBA): string {
    if (c.a !== undefined && c.a < 1) {
        return `Qt.rgba(${c.r}, ${c.g}, ${c.b}, ${c.a})`;
    }
    return `Qt.rgba(${c.r}, ${c.g}, ${c.b}, 1)`;
}
