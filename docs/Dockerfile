FROM oven/bun AS base

WORKDIR /frontend

COPY package*.json bun.lockb ./

RUN bun install

COPY . .

ENV NODE_ENV production

RUN bun run build

FROM base AS final

WORKDIR /frontend

COPY --from=base /frontend .

EXPOSE 3000

CMD [ "bun", "start" ]