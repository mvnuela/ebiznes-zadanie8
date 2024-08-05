# Use the official Node.js image as a base image
FROM node:14-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the source code to the working directory
COPY . .

# Build the React app
RUN npm run build

# Use a minimal Node.js image to serve the built React app
FROM nginx:alpine

# Copy the built React app to the nginx html folder
COPY --from=0 /app/build /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Command to start nginx
CMD ["nginx", "-g", "daemon off;"]
