se image
FROM nginx:alpine

# Set working directory to Nginx web root
WORKDIR /usr/share/nginx/html

# Copy your HTML files from build context
COPY . .

# Expose HTTP port
EXPOSE 80

# Nginx image already has the correct CMD
# so no need to override, but just for clarity:
CMD ["nginx", "-g", "daemon off;"]

