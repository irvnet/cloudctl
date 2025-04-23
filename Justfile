# Justfile for cloudctl project
# Run `just` to see the default command below.
# Uncomment as you implement each task.

# Set environment vars from .env (if present)
set dotenv-load

# Default: list available commands
default:
    @just --summary

# --- Provisioning ---

# create-devbox:
#     cloudctl create devbox

# create-ctrl:
#     cloudctl create ctrl

# create-workers:
#     cloudctl create workers --count 2

# delete-devbox:
#     cloudctl delete devbox

# delete-ctrl:
#     cloudctl delete ctrl

# delete-workers:
#     cloudctl delete workers

# --- Dev/Test ---

# test:
#     go test ./...

# lint:
#     staticcheck ./...

# run *args:
#     go run main.go {{args}}

# --- Helpers ---

# status:
#     cloudctl status

# setup-env:
#     echo "export HCLOUD_TOKEN=your-token-here" > .env
