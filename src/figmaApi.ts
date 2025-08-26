import { GetFileNodesResponse } from "@figma/rest-api-spec";

export async function fetchComponent(
    fileKey: string,
    componentId: string,
    figmaToken: string,
): Promise<GetFileNodesResponse> {
    const url = `https://api.figma.com/v1/files/${fileKey}/nodes?ids=${componentId}`;
    const response = await fetch(url, {
        headers: {
            "X-Figma-Token": figmaToken,
        },
    });
    if (!response.ok) throw new Error(`Error: ${response.statusText}`);
    return await response.json();
}
