# product-service-mongodb
test code simple microservice, you can see how i write code in go, I implement authorization to add/edit/delete product, so need login from auth-service-mongodb using username:password that has been registered from user-service-mongodb don't need to login to get products
this service function as same as product-service, but use mongodb as the databases and make some improved code
UPDATE: I'm adding caching feature for example of slow query(getting data) using Redis.