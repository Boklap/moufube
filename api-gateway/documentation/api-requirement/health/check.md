# Health Check API

## Overview
Performs a basic health check on the API Gateway instance to verify that the service is running and operational. This endpoint is used for monitoring, load balancer health checks, and service discovery systems.

## Endpoint
```
GET /api/v1/health
```

## Request
This endpoint does not require any request body or parameters.

### Example Request
```bash
curl -X GET https://api.example.com/api/v1/health
```

## Responses

### Success Response (200 OK)
```json
{
    "success": true,
    "message": "Instance is healthy.",
}
```

### Error Responses

#### Service Unavailable (503 Service Unavailable)
```json
{
    "success": false,
    "error": {
        "code": "SERVICE_UNAVAILABLE",
        "message": "Service is currently unavailable"
    }
}
```

#### Internal Server Error (500 Internal Server Error)
```json
{
    "message": "Internal server error occurred"
}
```

## HTTP Status Codes
- `200 OK` - Service is healthy and operational
- `500 Internal Server Error` - Unexpected server error
- `503 Service Unavailable` - Service is temporarily unavailable

## Usage Examples

### cURL Example
```bash
curl -X GET https://api.example.com/api/v1/health \
  -H "Accept: application/json"
```

### JavaScript/Fetch Example
```javascript
const checkHealth = async () => {
    try {
        const response = await fetch('/api/v1/health');
        const data = await response.json();

        if (data.success) {
            console.log('Service is healthy:', data.message);
            return { healthy: true, message: data.message };
        } else {
            console.error('Service health check failed:', data.error.message);
            return { healthy: false, error: data.error };
        }
    } catch (error) {
        console.error('Health check request failed:', error);
        return { healthy: false, error: error.message };
    }
};

// Usage
checkHealth().then(result => {
    if (result.healthy) {
        console.log('✅ API Gateway is operational');
    } else {
        console.log('❌ API Gateway health check failed');
    }
});
```

### Shell Script Example
```bash
#!/bin/bash

HEALTH_URL="https://api.example.com/api/v1/health"
STATUS=$(curl -s -o /dev/null -w "%{http_code}" $HEALTH_URL)

if [ $STATUS -eq 200 ]; then
    echo "✅ API Gateway is healthy (HTTP $STATUS)"
    exit 0
else
    echo "❌ API Gateway health check failed (HTTP $STATUS)"
    exit 1
fi
```

### Docker Health Check Example
```yaml
version: '3.8'
services:
  api-gateway:
    image: moufube/api-gateway:latest
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/api/v1/health"]
      interval: 30s
      timeout: 10s
      retries: 3
      start_period: 40s
```

## Monitoring Integration

### Prometheus Metrics
```yaml
scrape_configs:
  - job_name: 'api-gateway-health'
    static_configs:
      - targets: ['api-gateway:8080']
    metrics_path: '/api/v1/health'
    params:
      format: ['prometheus']
```

### Load Balancer Configuration
```nginx
upstream api_gateway {
    server api-gateway-1:8080;
    server api-gateway-2:8080;
    server api-gateway-3:8080;
}

server {
    listen 80;
    location / {
        proxy_pass http://api_gateway;
        health_check uri=/api/v1/health;
    }
}
```

## Notes
- This endpoint is designed to be lightweight and fast
- No authentication is required for health checks
- The endpoint should respond within milliseconds
- Consider implementing more detailed health checks for production monitoring
- Health checks are typically called frequently by monitoring systems
- The response format follows the standard API response structure