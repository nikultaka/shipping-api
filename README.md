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

### Concurrency
- Go routines with channels for parallel carrier API calls
- 5-second timeout to prevent request hanging
- Error isolation to prevent carrier failures from affecting others

### Carrier Simulation
- Simulated APIs for BlueDart, Delhivery, and XpressBees
- Random scenarios: API failures (10%), slow responses (20%), unserviceable locations (15%)

### Database Design
- Orders and Tracking Events tables with auto-migration
- Indexes on frequently queried fields (order_id, tracking_id)

### Error Handling
- Graceful fallback to in-memory storage during database outages
- Request validation at controller level
- Continues processing on external API failures

## Trade-offs

| Aspect | Choice | Reason |
|--------|--------|--------|
| Carrier Integration | Simulated APIs | Focus on core logic without external dependencies |
| Data Persistence | In-memory fallback | Service availability during database outages |
| Validation | Basic field validation | Core requirements focus |
| Tracking Updates | Polling API | Simpler implementation |
| Database | MySQL with GORM | ACID compliance and structured data |

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
