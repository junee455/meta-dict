# Use the official Node.js image as a base
FROM node:18-alpine AS base

# Set the working directory inside the container
WORKDIR /app

# Install dependencies
COPY package*.json ./

# Install dependencies (including dev dependencies)
RUN yarn

# Copy the rest of the application code
COPY . .

# Set the environment to development
ENV NODE_ENV=development

# Run the Next.js development server
CMD ["yarn", "dev"]
