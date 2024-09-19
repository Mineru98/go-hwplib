# Stage 1: Build the necessary modules on Alpine
FROM alpine:3.18 AS builder

# Install required packages for building ngx_brotli
RUN apk add --no-cache \
    git \
    build-base \
    cmake \
    pcre-dev \
    zlib-dev \
    openssl-dev \
    wget

# Download and compile ngx_brotli module
RUN git clone --recurse-submodules https://github.com/google/ngx_brotli.git /usr/local/src/ngx_brotli && \
    cd /usr/local/src/ngx_brotli && \
    cd deps/brotli && mkdir out && cd out && \
    cmake -DCMAKE_BUILD_TYPE=Release -DBUILD_SHARED_LIBS=OFF .. && \
    cmake --build . --config Release --target brotlienc && \
    cd ../../.. && \
    wget http://nginx.org/download/nginx-1.27.1.tar.gz && \
    tar -zxvf nginx-1.27.1.tar.gz && \
    cd nginx-1.27.1 && \
    ./configure --with-compat --add-dynamic-module=/usr/local/src/ngx_brotli && \
    make modules

# Stage 2: Create a minimal runtime image
FROM nginx:1.27.1-alpine

# Copy Brotli modules from the builder stage
COPY --from=builder /usr/local/src/ngx_brotli/nginx-1.27.1/objs/ngx_http_brotli_filter_module.so /etc/nginx/modules/
COPY --from=builder /usr/local/src/ngx_brotli/nginx-1.27.1/objs/ngx_http_brotli_static_module.so /etc/nginx/modules/

# Update nginx configuration to load the Brotli modules
RUN echo "load_module modules/ngx_http_brotli_filter_module.so;" >> /etc/nginx/nginx.conf && \
    echo "load_module modules/ngx_http_brotli_static_module.so;" >> /etc/nginx/nginx.conf

# Expose the default nginx port
EXPOSE 80

# Command to run nginx
CMD ["nginx", "-g", "daemon off;"]
