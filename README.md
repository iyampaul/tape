# Tape: Tiny Action Processing Engine

Tape is a lightweight, modular action processing engine designed for easy integration and automation. It allows users to define, configure, and execute custom or core actions via HTTP endpoints, supporting authentication, logging, and flexible input/output handling. Tape is ideal for orchestrating scripts, managing workflows, and securely exposing system operations in both local and containerized environments.

## Setup

1. **Clone the repository**
	```sh
	git clone <repo-url>
	cd tape
	```
2. **Install Go (if running locally)**
	- Requires Go 1.22 or newer.

3. **Install dependencies:**
	```sh
	go mod tidy
	```

4. **Edit `config.yml`** to match your environment (ports, paths, etc).

5. **Update action packages**

## Running TAPE

### Local

1. **Build package**
    ```sh
    make build
    ```
2. **Run TAPE** 
    ```sh
    make run
    ```
3. **Check initialized actions and keys**
    ```sh
    cat ./logs/event.log
    ```

### Docker

1. **Update docker-compose.yml and Dockerfile**
    - **If TLS is enabled**: Key and cert filenames must match the config.yml
    - Default port bind is 8080 => 443, modify as needed

2. **Build the Docker Image**
	```sh
	make docker-build
	```
3. **Execute TAPE container**
	```sh
	make docker-up
	```
	- This will start the app, mount configs, actions, logs, and files.
	- The app will be available on the port specified in `docker-compose.yml` (default: 8080).

## Actions

- Each action must have a unique `route`
- `generate_keys` (`keygen` for core actions) enables or disables authentication and generates a key for the action. If disabled, they key won't be required.

## Core Actions

- Core actions are built-in and configured in `config.yml`.
- Schema:
    ```yaml
        enabled:              # bool
        method:               # string
        keygen:               # bool
        input:                # bool
        data:                 # string
        output_write:         # bool
        output_file:          # string
        route:                # string
        description:          # string
    ```

### Core Action: Log

- Writes the incoming request out to event log.

## Custom Actions

### Example Action
  ```yaml
  name: "whoami"
  description: "Show current user"
  route: "whoami"
  method: POST
  generate_keys: true
  action: "/usr/bin/whoami"
  accept_input: false
  output_write: true
  output_file: "whoami.log"
  ```

## Test / Validation

### Example cURL command to reach a TAPE instance:
```sh
curl -k -X POST https://<ip_address>:<port>/<route>/<key> -H "Content-Type: application/json" -d '{"data": "test message"}
```