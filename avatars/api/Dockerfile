FROM python:3.10-slim-bullseye AS base

ENV PIP_DISABLE_PIP_VERSION_CHECK="1"

WORKDIR /app
RUN --mount=type=bind,source=api/requirements.txt,target=requirements.txt \
  --mount=type=cache,target=/root/.cache/pip \
  pip install -r requirements.txt


FROM base AS dev

WORKDIR /app
ENV FLASK_APP=/app/app.py
ENV FLASK_ENV=development

CMD [ "flask", "run", "--host=0.0.0.0", "--port=80" ]


FROM base AS dev_watch

COPY api/ ./api/

ENV FLASK_APP=./api/app.py
ENV FLASK_ENV=development

CMD [ "flask", "run", "--host=0.0.0.0", "--port=80" ]

