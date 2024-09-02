# Dockerfile
FROM node:18-alpine

WORKDIR /app

# Install pnpm
RUN npm install -g pnpm

# Install dependencies
COPY pnpm-lock.yaml ./
COPY package.json ./
RUN pnpm install

# Copy the rest of the application
COPY . .

# Run Prisma migrations
CMD ["pnpm", "prisma", "migrate", "deploy"]
