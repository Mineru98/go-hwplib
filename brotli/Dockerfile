FROM nginx:1.27.1

# Install required packages for building ngx_brotli
RUN apt-get update && \
    apt-get install -y \
    git \
    build-essential \
    cmake \
    zlib1g-dev \
    libpcre3-dev \
    libssl-dev \
    wget \
    curl \
    gnupg2 \
    ca-certificates \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

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
    make modules && \
    cp objs/ngx_http_brotli_filter_module.so /etc/nginx/modules/ && \
    cp objs/ngx_http_brotli_static_module.so /etc/nginx/modules/ && \
    rm -rf /usr/local/src/ngx_brotli nginx-1.27.1* 

# Update nginx configuration to load the Brotli modules
RUN echo "load_module modules/ngx_http_brotli_filter_module.so;" >> /etc/nginx/nginx.conf && \
    echo "load_module modules/ngx_http_brotli_static_module.so;" >> /etc/nginx/nginx.conf

# Expose the default nginx port
EXPOSE 80

# Command to run nginx
CMD ["nginx", "-g", "daemon off;"]
