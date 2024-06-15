alias c := clean

python := "./.venv/bin/python"
pip := "./.venv/bin/pip"

[private]
default:
  @just --choose

[private]
venv:
  @if [ ! -d .venv ]; then \
    python3 -m venv .venv ; cd backend ; .{{pip}} install -r requirements.txt ; \
  fi 

api command='run' extra='false': venv
  @if [ '{{extra}}' = 'install' ]; then \
    {{pip}} install -r backend/requirements.txt ; \
  fi

  @{{python}} -m fastapi_cli {{if command == "dev" {"dev"} else {"run"} }} backend/src/main.py

dev:
    @just api dev

app:
  @docker compose up --build -d

clean arg='false':
  @docker compose down {{if arg == "v" {"-v"} else {""} }}
