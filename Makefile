.PHONY: check-contracts check-backend check-frontend check-all

check-contracts:
	@echo "Checking Contracts..."
	cd contracts && forge build --sizes && forge fmt --check && forge test -vvv

check-backend:
	@echo "Checking Backend..."
	cd backend && go build -v ./... && go test -v ./...

check-frontend:
	@echo "Checking Frontend..."
	cd frontend && pnpm run lint && pnpm run build

check-all: check-contracts check-backend check-frontend
	@echo "All checks passed!"
