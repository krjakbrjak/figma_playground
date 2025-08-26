import yargs from "yargs";
import { hideBin } from "yargs/helpers";

export type CliArgs = {
    fileKey: string;
    componentId: string;
    figmaToken: string;
};

export const args: CliArgs = yargs(hideBin(process.argv))
    .usage(
        "Usage: $0 --fileKey <fileKey> --componentId <componentId> --figmaToken <figmaToken>",
    )
    .option("fileKey", {
        type: "string",
        demandOption: true,
        describe: "Figma file key",
    })
    .option("componentId", {
        type: "string",
        demandOption: true,
        describe: "Figma component ID",
    })
    .option("figmaToken", {
        type: "string",
        demandOption: true,
        describe: "Figma API token",
    })
    .help()
    .parseSync() as CliArgs;
