# DeepLX Load Balancer

This project implements a high-performance, highly available load balancer for DeepLX servers using Go.

## Features

- High-performance load balancing
- Caching of translation requests
- Health checking of DeepLX servers
- Performance tracking and adaptive load balancing
- Prometheus metrics integration
- Configurable via YAML

## Getting Started

### Prerequisites

- Go 1.16 or later
- Docker (recommend, for containerized deployment)

### Installation

#### Docker

1. Clone the repository:
    
    ```bash
    git clone https://github.com/your-username/deeplx-load-balancer.git
    cd deeplx-load-balancer
    ```
2. Build the Docker image:
    
    ```bash
    docker build -t deeplx-load-balancer .
    ```
3. Run the Docker container:
    
    ```bash
    docker run -p 8080:8080 deeplx-load-balancer
    ```
4. Access the load balancer at `http://localhost:8080`.

#### Manual

1. Clone the repository:
    
    ```bash
    git clone https://github.com/your-username/deeplx-load-balancer.git
    cd deeplx-load-balancer
    ```
3. Install dependencies:

    ```bash
    go mod download
    ```

### Configuration

Edit the `config/config.yaml` file to set up your environment-specific configurations.

### Running the Application

1. Build the application:
    
    ```bash
    go build -o deeplx-load-balancer ./cmd/server
    ```

2. Run the application:
    
    ```bash
    ./deeplx-load-balancer
    ```

## Usage

The load balancer exposes a `/translate` endpoint that accepts POST requests with the following format:

    ```bash
    curl -X POST https://yourdomain.com/{API_KEY}/translate \
         -H "Content-Type: application/json" \
         -d '{
               "text": "Hello, world!",
               "source_lang": "EN",
               "target_lang": "CN"
             }'
    ```

## Monitoring

Prometheus metrics are exposed at the `/metrics` endpoint.

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct and the process for submitting pull requests.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
