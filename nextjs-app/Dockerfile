# Use the official Node.js image as base
FROM node:21-alpine as build

# Set working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the application code
COPY . .

# Build the Next.js application
RUN npm run build

# Expose the port Next.js is running on
EXPOSE 3000

# Command to run the application
CMD ["npm", "start"]
