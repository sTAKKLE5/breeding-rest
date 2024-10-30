# Breeding Probability Calculator REST API

A Go-based REST API for calculating breeding probabilities in plant genetics. 
This service helps determine the likelihood of specific trait combinations in F2 generation plants.

## 📋 Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [API Documentation](#api-documentation)
- [Usage Examples](#usage-examples)
- [Development](#development)

## 🚀 Features
- Calculate F2 generation trait probabilities
- Support for multiple genetic traits
- Handle dominant and recessive gene expressions
- RESTful API with Swagger documentation
- JSON request/response format
- Input validation and error handling

## 🔧 Prerequisites
- Go 1.19 or higher
- Git

## 💻 Installation

1. Clone the repository:
```bash
git clone https://github.com/stakkle5/breeding-rest.git
cd breeding-rest
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o breeding-rest
```

## ⚙️ Configuration

The application uses environment variables for configuration. Create a `.env` file in the project root:

```env
PORT=8080
GIN_MODE=debug  # Use 'release' for production
```

## 📚 API Documentation

### Swagger Documentation

The API is documented using Swagger/OpenAPI. Access the Swagger UI at:
```
http://localhost:8080/swagger/index.html
```

To regenerate Swagger documentation:
```bash
swag init
```

### API Endpoints

#### Calculate Trait Probability
```
POST /api/v1/trait-probability
```

Request body example:
```json
{
  "motherPlant": {
    "name": "Purple Flash",
    "traits": [
      {
        "name": "Regular Leave Shape",
        "dominant": true,
        "geneLabel": "L"
      },
      {
        "name": "Purple Foliage",
        "dominant": false,
        "geneLabel": "C"
      }
    ]
  },
  "fatherPlant": {
    "name": "Candlelight Mutant",
    "traits": [
      {
        "name": "Mutant Leave Shape",
        "dominant": false,
        "geneLabel": "L"
      },
      {
        "name": "Green Foliage",
        "dominant": true,
        "geneLabel": "C"
      }
    ]
  },
  "plantCount": 64
}
```

Response example:
```json
{
  "plantCount": 64,
  "motherPlantName": "Purple Flash",
  "fatherPlantName": "Candlelight Mutant",
  "combinations": [
    {
      "description": "Regular Leave Shape, Green Foliage",
      "genotype": "L_ C_",
      "numberOfPlants": 36,
      "percentage": 56.25
    },
    {
      "description": "Regular Leave Shape, Purple Foliage",
      "genotype": "L_ cc",
      "numberOfPlants": 12,
      "percentage": 18.75
    }
    // ... more combinations
  ]
}
```

## 📝 Usage Examples

### Using cURL:
```bash
curl -X POST "http://localhost:8080/api/v1/trait-probability" \
     -H "Content-Type: application/json" \
     -d @example-request.json
```

### Using Go Client:
```go
package main

import (
    "bytes"
    "encoding/json"
    "net/http"
)

func main() {
    request := models.TraitRequest{
        // ... populate request
    }
    
    jsonData, _ := json.Marshal(request)
    resp, err := http.Post(
        "http://localhost:8080/api/v1/trait-probability",
        "application/json",
        bytes.NewBuffer(jsonData),
    )
    // Handle response
}
```

## 🛠️ Development

### Project Structure
```
breeding-rest/
├── main.go
├── internal/
│   ├── handlers/
│   │   └── traitProbability.go
│   ├── models/
│   │   └── trait.go
│   └── routes/
│       └── routes.go
├── docs/
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
