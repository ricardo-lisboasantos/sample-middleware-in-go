# sample-middleware-in-go

A simple middleware written in go. This is a learning exercise project.
The idea is to write a simple middleware in which i'm applying the concepts as i learn them.
At the moment of this writing, the middleware is able to:
- Serve static files
- Serve dynamic content
- Handle GET and POST requests
- Log all requests
- Validate the request method

New features will be added as i learn them.

## How to build

On linux:
```bash
./buildxc.sh
```

On windows:
```powershell
NOT IMPLEMENTED YET
```

The executable will be in the build/<platform> folder.

## How to install

The server can be started simply by running the executable.

On linux:
1. Copy the executable to a folder in your PATH.
2. Run the executable.


## How to run

On linux (assuming the executable is in your PATH):
```bash
mw-server
```
else
```bash
./mw-server
```

On windows:
```powershell
./mw-server.exe
```

## How to use

The server can be accessed on port 8080.

The server has 4 endpoints:
| Endpoint | Description |
| --- | --- |
| /hello | Returns a simple hello world message |
| /contactme | Returns a simple contact me form |
| /static/ | Returns a static file |
| /api/contactme | Returns a json response |


