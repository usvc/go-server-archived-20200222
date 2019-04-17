# `usvc/go-server`

# Example usage

Include it in your code:

```go
package main

import (
  "net/http"
  "github.com/usvc/go-server"
)

func main() {
  config := server.NewConfigFromEnvironment()
  server := server.NewServer(config)
  handler := http.NewServeMux()
  handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello world"))
  })
  server.StartWithHandler(handler)
}

```

# Configuration

## Via Environment Variables

| Key | Description |
| --- | --- |
| `HOST` | Specifies the host interface to bind to |
| `LIVENESS_CHECK_INTERVAL` | Specifies a duration to check for responsiveness |
| `LIVENESS_CHECK_METHOD` | Specifies the method used to check for responsiveness |
| `LIVENESS_CHECK_PATH` | Specifies the path used to check for responsiveness |
| `LIVENESS_CHECK_STATUS_CODE` | Specifies the expected status code for the server to be considered responsive|
| `LIVENESS_CHECK_TIMEOUT` | Specifies the duration of time before responsiveness checks will time out and fail |
| `MAX_HEADER_BYTES` | Specifies the maximum number of bytes allowed in a header |
| `PORT` | Specifies the port to listen on to |
| `TIMEOUT_IDLE` | Specifies the duration of an idle connection before timing out |
| `TIMEOUT_READ` | Specifies the duration of a request read before timing out |
| `TLS_CERTIFICATE_PATH` | Specifies the relative path to the `.crt` file for enabling TLS |
| `TLS_KEY_PATH` | Specifies the relative path to the `.key` file for enabling TLS |

> To prefix the variables with a string, add it as a parameter in the call to `NewConfigFromEnvironment()`. For example, `server.NewConfigFromEnvironment("SERVER")` will read from `SERVER_HOST` instead of just `HOST`.

# Development

## With TLS
To test the TLS, you'll need to generate a cert/key pair. The included `Makefile` has the relevant recipes. Simply run:

```sh
make secrets
```

The above generates two sets of certificates/keys. Next, create symlinks of one of the set of keys with:

```sh
make tls
```

## Continuous Integration Pipeline

### Variables

| Key | Description |
| --- | --- |
| GITHUB_SSH_DEPLOY_KEY | Create a deploy key for your repository (generate one with `make ssh.keys`, copy contents of `./bin/_id_rsa.pub` into your repository's Deploy Keys and use contents of `./bin/_id_rsa_b64` as the value of this variable) |
| GITHUB_REPOSITORY_URL | Get this from your repository's clone address |

# License
This project is licensed under the MIT license.

# Cheers