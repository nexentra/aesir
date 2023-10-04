FROM oven/bun

WORKDIR /frontend

COPY package*.json bun.lockb ./

RUN bun install

COPY . .

ENV NODE_ENV production

RUN bun run build

COPY . .

EXPOSE 3000

CMD [ "bun", "start" ]