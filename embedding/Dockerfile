FROM nvidia/cuda:12.3.1-base-ubuntu22.04

RUN groupadd -g 999 python && useradd -r -u 999 -g python python
RUN mkdir /usr/root && chown python:python /usr/root

WORKDIR /usr/root

RUN apt update
RUN apt -y install python3.12
RUN apt -y install python3-pip
RUN pip install poetry

COPY poetry.lock .
COPY pyproject.toml .

RUN poetry config virtualenvs.create false
RUN poetry install --without dev --with embedding-server

COPY embedding/src/ .src/
COPY settings.toml .

USER 999
EXPOSE 5501

ENTRYPOINT [ "python3", "src/main.py" ]
