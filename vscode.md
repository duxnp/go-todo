# VS Code Settings

## Extensions

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)
- [Tailwind CSS IntelliSense](https://marketplace.visualstudio.com/items?itemName=bradlc.vscode-tailwindcss)
- [templ-vscode](https://marketplace.visualstudio.com/items?itemName=a-h.templ)

## Tasks

This configuration will tell VS Code to launch two terminal commands simultaneously in a split panel. One terminal panel will run Air. Air will automatically run `make build` whenever any project file changes. The other terminal panel will run `make templ`. This will start a templ watcher that will automatically regenerate the templates when a \*.templ file changes, then refresh the page in the browser after the app is finished reloading.

Put the following in ./.vscode/tasks.json:

```json
{
  // See https://go.microsoft.com/fwlink/?LinkId=733558
  // for the documentation about the tasks.json format
  "version": "2.0.0",
  "tasks": [
    {
      "label": "air",
      "type": "shell",
      "command": "air",
      "presentation": {
        "panel": "shared",
        "group": "watchers"
      },
      "problemMatcher": []
    },
    {
      "label": "templ watch",
      "type": "shell",
      "command": "make templ",
      "presentation": {
        "panel": "shared",
        "group": "watchers"
      },
      "problemMatcher": []
    },
    {
      "label": "dev hot reload",
      "detail": "Go, Templ, and Tailwind hot reloading",
      "icon": {
        "color": "terminal.ansiGreen",
        "id": "eye-watch"
      },
      "type": "shell",
      "dependsOrder": "parallel",
      "dependsOn": ["air", "templ watch"],
      "problemMatcher": []
    }
  ]
}
```

## Debugging

Put the following in ./.vscode/launch.json

```json
{
  // Use IntelliSense to learn about possible attributes.
  // Hover to view descriptions of existing attributes.
  // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Launch Package",
      "type": "go",
      "request": "launch",
      "mode": "auto",
      "envFile": "${workspaceFolder}/.env",
      "cwd": "${workspaceFolder}",
      "program": "${workspaceFolder}/cmd/api/main.go"
    }
  ]
}
```

## Syntax Hightlighting

If you are using gopls, you can enable Semantic Highlighting for more accurate syntax highlighting based on semantic tokenization using this setting:

```json
"gopls": {
  "ui.semanticTokens": true,

  // you can optionally turn on these features for more colors
  // see https://go.dev/issue/45753 and https://go.dev/issue/45792
  "ui.noSemanticString": true,  // delegates string syntax highlighting to vscode
  "ui.noSemanticNumber": true,  // delegates number syntax highlighting to vscode
}
```

## Templ

```json
  "emmet.includeLanguages": {
    "templ": "html"
  },
  "tailwindCSS.includeLanguages": {
    "templ": "html"
  },
```
