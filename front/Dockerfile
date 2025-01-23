# Use the official Node.js image as a base
FROM node:18-alpine AS base

# Set the working directory inside the container
WORKDIR /app

# Same as npm install
RUN yarn install --frozen-lockfile

# Copy the rest of the application code
COPY . .

# Set the environment to development
ENV NODE_ENV=development

# Run the Next.js development server
CMD ["yarn", "dev"]
