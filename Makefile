.PHONY: dev dev-backend dev-frontend build clean

dev:
	@echo "Starting development servers..."
	@make -j2 dev-backend dev-frontend

dev-backend:
	@echo "Starting Go backend on :8080..."
	@cd backend && go run main.go

dev-frontend:
	@echo "Starting Vue frontend on :5173..."
	@cd frontend && npm run dev

build:
	@echo "Building applications..."
	@cd backend && go build -o bin/app main.go
	@cd frontend && npm run build

clean:
	@rm -rf backend/bin
	@rm -rf frontend/dist

install:
	@echo "Installing dependencies..."
	@cd backend && go mod tidy
	@cd frontend && npm install
