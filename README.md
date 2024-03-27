# pismo

# running the application 

## pre-requisite 
1. Should have the docker install , if not please follow : [link](https://docs.docker.com/desktop/install/mac-install/)

## steps to run application
1. Run the command `docker-compose build` to build the application 
2. Run the command `docker-compose up` to run the application in docker container
3. Once the above command ran successfuly it will ensure following things 
    1. The init scripts present in directory `scripts/init.sql` is executes and all the tables are created
    2. Postgres connection is established 
    3. application is running at port `http://0.0.0.0:8080`
    4. check the health using endpoint : `http://0.0.0.0:8080/pismo/api/v1/health`

# running test cases 

## pre-requisite
1. should have mockery install , if not please run command : `brew install mockery`

## steps to run test cases 
1. first generate the mocks , for this run command : `make generate-mocks` in root directory 
2. run the command `make tests` to run the test case

# postman collection 
1. Please check the `scripts/postman_collection.json` , which has the collection for apis