

[example 1](https://github.com/Shareable-Resources/eurus/blob/5b26db2a094715d7deb44e1a4f2c41cca966247c/eurus-backend/.vscode/tasks.json)

.vscode/tasks.json
```
{
    "version": "2.0.0",
    "cwd": "${workspaceFolder}",
    "tasks": [
        {
            "label": "compile release",
            "command": "go",
            "args": [
                "build",
                "./..."
            ],
            "group": "build",
            "type": "shell"
        },
        {
            "label": "compile debug",
            "command": "go",
            "args": [
                "build",
                "-gcflags",
                "all=-N -l",
                "-o",
                "${fileDirname}/__debug_bin"
                "./..."
            ],
            "options": {
                "cwd": "${workspaceFolder}",
                "env": {
                    "CGO_ENABLED": "1"
                 }
            },
            "type": "shell",
            "problemMatcher": []
        },
        {
            "label": "build TestServer debug",
            "command": "go",
            "args": [
                "build",
                "-gcflags=\"all=-N -l\"",
                "-o",
                "bin/debug/testServer",
                "./foundation/server/test/main"
            ],
            "group": "build",
            "type": "shell",
            "problemMatcher": []
        },
   ]
}
```

.vscode/settings.json
```
{
    "go.languageServerExperimentalFeatures": {
        "format": true,
        "autoComplete": true,
        "rename": true,
        "goToDefinition": true,
        "hover": true,
        "signatureHelp": true,
        "goToTypeDefinition": true,
        "goToImplementation": true,
        "documentSymbols": true,
        "workspaceSymbols": true,
        "findReferences": true,
        "diagnostics": true,
        "documentLink": true
      },
    "go.useLanguageServer": true,
    "go.delveConfig": {
      "debugAdapter": "legacy",
      "apiVersion": 2,
      "showGlobalVariables": false,
      "dlvLoadConfig": {
        "maxStringLen": 10000,
      },
    },
    "go.alternateTools": {
        "dlv": "<absolute path to your dlv binary>"
    }
    "go.lintTool": "golangci-lint",
    "go.testTags": "dev,unittest_prod",
    "go.testFlags": [
      "-v"
    ],
}    
```

.vscode/launch.json
```
{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [

        {
            "name": "Debug AuthServer",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/auth_service/AuthenMain.go",
            "env": {},
            "buildFlags": "-tags dev",
            "cwd":"${workspaceFolder}/bin/debug/",
            "args": ["--config", "${workspaceFolder}/bin/debug/AuthServerConfig.json"]
        },
        {
            "name": "Debug ConfigServer",
            "type": "go",
            "request": "launch",
            "mode": "debug",
            "program": "${workspaceFolder}/config_service/ConfigServerMain.go",
            "env": {},
            "buildFlags": "-tags dev",
            "cwd":"${workspaceFolder}/bin/debug/",
            "args": ["--config", "${workspaceFolder}/bin/debug/ConfigServerConfig.json", "--pwServer", "/tmp/passwordServerSock"]
        },
    ],
}


https://github.com/golang/vscode-go/blob/master/docs/debugging.md

```
{
    "tasks": [
        {
            "label": "go: build (debug)",
            "type": "shell",
            "command": "go",
            "args": [
                "build",
                "-gcflags=all=-N -l",
                "-o",
                "${fileDirname}/__debug_bin"
            ],
            "options": {
                "cwd": "${fileDirname}"
            },
        }
    ]
}
```


```
    ...
    "configurations": [
        {
            "name": "Launch Package as root",
            "type": "go",
            "request": "launch",
            "mode": "exec",
            "asRoot": true,
            "console": "integratedTerminal",
            "program": "${fileDirname}/__debug_bin",
            "preLaunchTask": "go: build (debug)",
        }
    ]
```

"go.alternateTools": {
    "dlv": "<absolute path to your dlv binary>"
}

