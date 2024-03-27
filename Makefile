service-mocks:
	mockery --dir=internal/app/pismo/services --name=IAccountService --filename=account_service.go --output=internal/app/pismo/services/mocks --outpkg=service_mocks
	mockery --dir=internal/app/pismo/services --name=IOperationTypeService --filename=operation_type_service.go --output=internal/app/pismo/services/mocks --outpkg=service_mocks
	mockery --dir=internal/app/pismo/services --name=ITransactionService --filename=transaction_service.go --output=internal/app/pismo/services/mocks --outpkg=service_mocks

repository-mocks:
	mockery --dir=internal/app/pismo/repository --name=IAccountRepo --filename=account_repo.go --output=internal/app/pismo/repository/mocks --outpkg=repository_mocks
	mockery --dir=internal/app/pismo/repository --name=IOperationTypeRepo --filename=operation_type_repo.go --output=internal/app/pismo/repository/mocks --outpkg=repository_mocks
	mockery --dir=internal/app/pismo/repository --name=ITransactionRepo --filename=transaction_repo.go --output=internal/app/pismo/repository/mocks --outpkg=repository_mocks

generate-mocks:
	# service mocks
	make service-mocks
	# repository mocks
	make repository-mocks