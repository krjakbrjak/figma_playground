# Figma QML Code Generation Playground

This project demonstrates how to generate QML components from Figma components using TypeScript.  
It leverages Figma's official API types and schema published as an npm package, allowing direct integration without custom API definitions.

<p align="center">
  <img src="./figma2qml.svg" width="550">
</p>

## Features

- Fetches Figma component data using the Figma REST API.
- Uses TypeScript for type safety and maintainability.
- Generates QML code from Figma nodes via Handlebars templates.
- Modular structure for easy extension and understanding.

## Build and Run

```bash
yarn install
yarn build
yarn start --fileKey FILE_KEY --componentId COMPONENT_ID --figmaToken FIGMA_TOKEN
```

Or for development:

```bash
yarn dev --fileKey FILE_KEY --componentId COMPONENT_ID --figmaToken FIGMA_TOKEN
```

## Requirements
* Node.js (v18+ recommended)
* Figma API token

## How it works

The CLI fetches a Figma component using the provided file key and component ID.
The data is parsed and rendered into QML using Handlebars templates.
The generated QML is printed to the console.
