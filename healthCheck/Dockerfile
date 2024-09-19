# Stage 1: Build NGINX with the upstream check module
FROM alpine:3.18 AS builder

# Install necessary packages for building NGINX
RUN apk add --no-cache \
    build-base \
    curl \
    pcre-dev \
    zlib-dev \
    openssl-dev \
    linux-headers \
    git

# Set the NGINX and module versions
ENV NGINX_VERSION=1.27.1
ENV MODULE_REPO=https://github.com/yaoweibin/nginx_upstream_check_module.git

# Download NGINX source
RUN curl -fSL https://nginx.org/download/nginx-${NGINX_VERSION}.tar.gz -o nginx.tar.gz \
    && tar -zxC /tmp -f nginx.tar.gz

# Clone the upstream check module
RUN git clone ${MODULE_REPO} /tmp/nginx_upstream_check_module

# Configure and build NGINX with the upstream check module
WORKDIR /tmp/nginx-${NGINX_VERSION}
RUN ./configure \
    --with-http_ssl_module \
    --add-module=/tmp/nginx_upstream_check_module \
    --prefix=/etc/nginx \
    --sbin-path=/usr/sbin/nginx \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --pid-path=/var/run/nginx.pid \
    --lock-path=/var/run/nginx.lock \
    --with-http_stub_status_module \
    --with-stream \
    --with-stream_ssl_module \
    && make && make install

# Stage 2: Create the final lightweight NGINX image
FROM alpine:3.18

# Install runtime dependencies
RUN apk add --no-cache \
    pcre \
    zlib \
    openssl \
    bash

# Copy NGINX from the builder stage
COPY --from=builder /usr/sbin/nginx /usr/sbin/nginx
COPY --from=builder /etc/nginx /etc/nginx

# Create necessary directories
RUN mkdir -p /var/log/nginx /var/run/nginx

# Expose HTTP and HTTPS ports
EXPOSE 80 443

# Set the default command to run NGINX
CMD ["nginx", "-g", "daemon off;"]
