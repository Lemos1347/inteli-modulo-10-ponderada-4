# src/main.py

import logging

from fastapi import FastAPI, Request
from logging_config import LoggerSetup
from routers import produtos, usuarios

# Cria um logger raiz
logger_setup = LoggerSetup()

# Adiciona o logger para o módulo
LOGGER = logging.getLogger(__name__)

app = FastAPI(root_path="/fastapi")

app.include_router(usuarios.router)
app.include_router(produtos.router)


@app.get("/")
async def root(request: Request):
    LOGGER.info("Acessando a rota /")
    return {"message": "Hello World"}
