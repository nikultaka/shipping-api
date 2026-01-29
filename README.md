# Shipping API

## How to Run

### Prerequisites
- Go 1.21 or higher
- MySQL 5.7 or higher

### Installation Steps

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd shipping-api
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Configure environment file:
   ```bash
   cp .env.example .env
   ```
   Update `.env` with your database credentials

4. Create database:
   ```sql
   CREATE DATABASE shipping_db;
   ```

5. Run the server:
   ```bash
   go run cmd/main.go
   ```

The API will be available at `http://localhost:8080`

## Design Decisions

### Layered Architecture
- Controllers: Handle HTTP requests and validation
- Services: Contain business logic
- Repository: Database operations using GORM
- Models: Data structures and DTOs
- Helpers: Utility functions

### Database Design
- Orders and Tracking Events tables with auto-migration
- Indexes on frequently queried fields (order_id, tracking_id)

### Error Handling
- Graceful fallback to in-memory storage during database outages
- Request validation at controller level
- Continues processing on external API failures

## What I'd Improve with More Time

- Redis caching for shipping rates
- JWT authentication and authorization
- Docker containerization and Kubernetes configuration
- Comprehensive unit and integration tests
- OpenAPI/Swagger documentation
- Real carrier API integrations
- Webhook support for tracking updates
- Performance monitoring with Prometheus
- Distributed tracing implementation
