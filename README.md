# Gera-AI backend

Stack/Libs used:
- Bcrypt
- Fiber
- Gorm
- PostgreSQL
- Swagger
- Validation

Configuration via environment variables (```.env``` file for docker):
```
POSTGRES_HOST=db
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=geraai
POSTGRES_PORT=5432
JWT_SECRET=your_jwt_secret_key
OPENAI_API_KEY=your_api_key
PROXY_URL=your_proxy_url
```

To run use ```compose.yml``` from this repo:
