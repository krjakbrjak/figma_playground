import { fetchComponent } from "./figmaApi";
import { render } from "./renderer";
import { args, CliArgs } from "./cli";

async function main(argv: CliArgs = args) {
    try {
        const data = await fetchComponent(
            argv.fileKey,
            argv.componentId,
            argv.figmaToken,
        );
        const document = data.nodes[argv.componentId].document;

        const rendered = render(document);
        console.log(`
import QtQuick
import QtQuick.Layouts
${rendered}
`);
    } catch (error) {
        console.error("Error fetching or rendering component:", error);
        process.exit(1);
    }
}

main();
