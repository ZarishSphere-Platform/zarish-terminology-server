# Zarish Terminology Server

**Part of the ZarishSphere Platform - A No-Code FHIR Healthcare Data Management System**

The Zarish Terminology Server provides FHIR-compliant terminology services for managing code systems, value sets, and concept maps. It enables standardized healthcare terminology management including SNOMED CT, LOINC, ICD-10, and custom terminologies.

## 🚀 Technology Stack

- **Language**: Go 1.23+
- **Web Framework**: Gin
- **Data Storage**: File-based JSON storage
- **Containerization**: Docker

## 📋 Prerequisites for Local Development

- **Go**: Version 1.23 or higher ([Download Go](https://go.dev/dl/))
- **Docker** (optional): For containerized deployment ([Download Docker](https://www.docker.com/))
- **Git**: For version control ([Download Git](https://git-scm.com/))

### Checking Your Installation

```bash
go version      # Should show go1.23 or higher
docker --version # Should show Docker version 20.x or higher
git --version   # Should show git version 2.x.x
```

## 🛠️ Step-by-Step Development Setup

### Step 1: Clone the Repository

```bash
cd ~/Desktop
git clone https://github.com/ZarishSphere-Platform/zarish-terminology-server.git
cd zarish-terminology-server
```

### Step 2: Install Dependencies

```bash
go mod download
go mod tidy
```

### Step 3: Configure Environment

Create a `.env` file:

```env
SERVER_PORT=8082
SERVER_HOST=0.0.0.0
TERMINOLOGY_DATA_PATH=./terminology
```

### Step 4: Start the Server

```bash
# Run the server
go run cmd/server/main.go

# Or build and run
go build -o zarish-terminology-server cmd/server/main.go
./zarish-terminology-server
```

Server starts at `http://localhost:8082`

### Step 5: Test the API

```bash
# Check server health
curl http://localhost:8082/health

# Get FHIR metadata
curl http://localhost:8082/fhir/metadata

# Search code systems
curl http://localhost:8082/fhir/CodeSystem

# Search value sets
curl http://localhost:8082/fhir/ValueSet
```

## 🔧 Available Commands

| Command | Description |
|---------|-------------|
| `go run cmd/server/main.go` | Start development server |
| `go build -o zarish-terminology-server cmd/server/main.go` | Build binary |
| `go test ./...` | Run all tests |
| `go mod tidy` | Clean up dependencies |

## 📁 Project Structure

```
zarish-terminology-server/
├── cmd/
│   └── server/
│       └── main.go         # Entry point
├── internal/
│   ├── api/               # HTTP handlers
│   ├── models/            # Data models
│   └── storage/           # Data storage
├── terminology/           # Terminology data files
│   ├── codesystems/      # Code system definitions
│   ├── valuesets/        # Value set definitions
│   └── conceptmaps/      # Concept map definitions
├── Dockerfile
├── go.mod
└── README.md
```

## 📚 Supported Terminologies

- **SNOMED CT**: Clinical terminology
- **LOINC**: Laboratory and clinical observations
- **ICD-10**: Disease classification
- **RxNorm**: Medication terminology
- **Custom**: Organization-specific terminologies

## 🔍 API Endpoints

### Code Systems

```bash
# List all code systems
GET /fhir/CodeSystem

# Get specific code system
GET /fhir/CodeSystem/{id}

# Lookup a code
GET /fhir/CodeSystem/$lookup?system={system}&code={code}
```

### Value Sets

```bash
# List all value sets
GET /fhir/ValueSet

# Get specific value set
GET /fhir/ValueSet/{id}

# Expand a value set
GET /fhir/ValueSet/$expand?url={url}

# Validate a code
GET /fhir/ValueSet/$validate-code?url={url}&code={code}
```

## 🐳 Docker Deployment

```bash
# Build image
docker build -t zarish-terminology-server .

# Run container
docker run -p 8082:8082 zarish-terminology-server
```

## 🐛 Troubleshooting

### Port Already in Use

```bash
lsof -i :8082
kill -9 <PID>
# Or change SERVER_PORT in .env
```

### Module Download Fails

```bash
go clean -modcache
go mod download
```

## 📚 Learning Resources

- [FHIR Terminology Services](https://www.hl7.org/fhir/terminology-service.html)
- [Go Documentation](https://go.dev/doc/)
- [SNOMED CT](https://www.snomed.org/)
- [LOINC](https://loinc.org/)

## 🤝 Contributing

1. Create a feature branch
2. Make changes
3. Write/update tests
4. Submit pull request

## 🔗 Related Repositories

- [zarish-fhir-server](https://github.com/ZarishSphere-Platform/zarish-fhir-server)
- [zarish-frontend-shell](https://github.com/ZarishSphere-Platform/zarish-frontend-shell)
- [zarish-fhir-data](https://github.com/ZarishSphere-Platform/zarish-fhir-data)
