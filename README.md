# Retail Pulse - Image Processing Service

This service processes thousands of store images **concurrently**, calculating their perimeters while handling multiple jobs simultaneously. It's designed to be scalable, efficient, and reliable, making it perfect for retail analytics applications.

### [Backend Task](https://kiranaclub.notion.site/Backend-Intern-Assignment-529d5850691d483db61c3561cfaa7293)

### Key Features

- Concurrent image processing
- Job status tracking
- RESTful API endpoints
- Error handling for failed downloads/processing
- Docker support
- In-memory job management

## Architecture

### API Endpoints
### Refer to Postman Tests for API structure.

- **Submit Job** (`POST /api/submit/`)
  - Creates a new image processing job
  - Returns a unique job ID
  - Validates input data

- **Get Job Status** (`GET /api/status?jobid=<id>`)
  - Retrieves current job status
  - Reports any errors encountered
  - Provides processing results

### Core Components

- Job Queue Manager
- Image Processor
- Store Master Integration
- Concurrent Download Manager
- Result Aggregator

## Assumptions

- In-memory Job Master storage (as specified)
- Images are accessible via HTTP/HTTPS URLs
- Store Master data is available and valid
- Network connectivity is stable
- Image URLs are properly formatted and accessible
- The service runs on a single instance (for the current implementation)
- Job IDs are unique integers
- Visit times are in a standard format (ISO 8601)

## Setup Instructions

### Prerequisites

- Go 1.21 or higher
- Docker (optional)
- Internet connection for image downloads

### Standard Setup

```bash
# Clone the repository
git clone https://github.com/kevintamakuwala/retail-pulse.git
cd retail-pulse

# Install dependencies
go mod download

# Run the service
go run cmd/server/main.go
```

### Docker Setup

```bash
# Build the Docker image
docker build -t retail-pulse .

# Run the container
docker run -p 8080:8080 retail-pulse
```

### Using Docker Hub

```bash
# Pull the Docker image from Docker Hub
docker pull kevintamakuwala/retail-pulse:latest

# Run the container
docker run -p 8080:8080 kevintamakuwala/retail-pulse:latest
```

## Future Improvements

- Replace in-memory storage with a database
- Add data archiving capabilities
- Add authentication/authorization
- Implement caching layer
- Optimize image download process
- Add batch processing capabilities
- Add retry mechanisms for failed downloads
- Implement circuit breakers


### Best regards,
### Kevin Tamakuwala.
