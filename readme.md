This is simple CLI app that runs HTTP server at localhost:8080 

Run:

_-> go build_
 -> httpCLI 

Test:

_-> go test_


Example usage:

httpCLI run --file <index.html> 


Available Commands: 


version                                                -      Displays app version info


help                             -       Help about any command


run [--file <index.html>]         -      Runs http server that serves alternate html file (file must be in current catalog or filepath)
