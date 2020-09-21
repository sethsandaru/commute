# Frontier

Frontier is the Presentation-Layer-Service of Commute. It's using Laravel 6 (LTS).

## Start-up

### Requirements
- Docker & Compose
- You might configure your own environment
- Create your own `.env` file based on [.env.example](./.env.example)

#### Docker Environment Configuration

Visit [docker](./docker) folder

### Build Docker Container
```bash
cd /<path>/frontier
docker-compose build
```

### Start the Project
```bash
docker-compose up -d
```

### Build the project
```bash
docker-compose run composer install
docker-compose run artisan key:generate
docker-compose run artisan migrate
docker-compose run npm install
docker-compose run npm run production
```

Then kindly visit and test it: [https://localhost:8080](https://localhost:8080)

### Test
```bash
docker-compose run phpunit
```
