# Build stage
FROM oven/bun AS builder

WORKDIR /app

COPY package*.json bun.lockb ./

RUN bun install --no-cache

COPY . .

ENV NEXT_TELEMETRY_DISABLED 1

ENV NODE_ENV production

RUN bun run build

# Run stage
FROM oven/bun AS runner

WORKDIR /app

ENV NEXT_TELEMETRY_DISABLED 1

COPY --from=builder /app/package.json .
COPY --from=builder /app/bun.lockb .
COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next/standalone ./standalone
COPY --from=builder /app/public /app/standalone/public
COPY --from=builder /app/.next/static /app/standalone/.next

EXPOSE 3000

CMD [ "bun", "start" ]