# Build Stage
FROM alpine:3.18 AS builder

# Install build dependencies
RUN apk add --no-cache \
    build-base \
    pcre-dev \
    zlib-dev \
    openssl-dev \
    zstd-dev \
    wget \
    git \
    ca-certificates

# Set Nginx version
ENV NGINX_VERSION 1.27.1

WORKDIR /usr/src

# Download Nginx source code
RUN wget http://nginx.org/download/nginx-$NGINX_VERSION.tar.gz && \
    tar -zxvf nginx-$NGINX_VERSION.tar.gz && \
    rm nginx-$NGINX_VERSION.tar.gz

# Clone zstd-nginx-module
RUN git clone https://github.com/tokers/zstd-nginx-module.git

WORKDIR /usr/src/nginx-$NGINX_VERSION

# Configure and build Nginx with the zstd module
RUN ./configure \
    --prefix=/etc/nginx \
    --sbin-path=/usr/sbin/nginx \
    --modules-path=/usr/lib/nginx/modules \
    --conf-path=/etc/nginx/nginx.conf \
    --error-log-path=/var/log/nginx/error.log \
    --http-log-path=/var/log/nginx/access.log \
    --pid-path=/var/run/nginx.pid \
    --lock-path=/var/run/nginx.lock \
    --with-http_ssl_module \
    --with-http_v2_module \
    --with-http_realip_module \
    --with-http_stub_status_module \
    --with-http_gzip_static_module \
    --with-http_sub_module \
    --add-module=/usr/src/zstd-nginx-module && \
    make && \
    make install

# Final Stage
FROM alpine:3.18

# Install runtime dependencies
RUN apk add --no-cache \
    pcre \
    zlib \
    openssl \
    zstd \
    ca-certificates

# Copy Nginx from the builder stage
COPY --from=builder /usr/sbin/nginx /usr/sbin/nginx
COPY --from=builder /etc/nginx /etc/nginx

WORKDIR /var/log/nginx
RUN ln -sf /dev/stdout /var/log/nginx/access.log \
    && ln -sf /dev/stderr /var/log/nginx/error.log

# Expose port 80
EXPOSE 80

# Run Nginx
CMD ["nginx", "-g", "daemon off;"]
