FROM node:18-bullseye-slim AS base

ENV YARN_CACHE_FOLDER=/cache/yarn

WORKDIR /app

# install dependencies
RUN --mount=type=bind,source=web/package.json,target=package.json \
  --mount=type=bind,source=web/yarn.lock,target=yarn.lock \
  --mount=type=cache,target=/cache/yarn \
  yarn install


FROM base AS dev

ENV VITE_HOST=0.0.0.0
ENV VITE_PORT=5173
CMD ["yarn", "dev"]


FROM base AS dev_watch

# copy app source
COPY web/ ./

ENV VITE_HOST=0.0.0.0
ENV VITE_PORT=5173
CMD ["yarn", "dev"]
