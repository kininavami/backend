package constants

const RestApiPrefix = "/api/v1/"

//Health Check URLs
const HealthCheckUrl = "/health"


//User Operations URLs
const CreateUser = RestApiPrefix + "user"
const GetAllUser = CreateUser
const GetUserByUsername = CreateUser + "/{username:[a-zA-Z0-9_-]+}"
const DeleteUserByUsername = GetUserByUsername