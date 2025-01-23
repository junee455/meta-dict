# Use the official Node.js image as a base
FROM node:18-alpine AS base

RUN apk add --no-cache libc6-compat

WORKDIR /app

COPY package.json /app/package.json
COPY yarn.lock /app/yarn.lock

RUN yarn install --frozen-lockfile

COPY . .

# Set the environment to development
ENV NODE_ENV=development
ENV NEXT_TELEMETRY_DISABLED=1

# Run the Next.js development server
CMD ["yarn", "dev"]
