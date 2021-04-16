#DIR /database/ (*)
mongodump --db ecommerce -o ./

mongorestore --db ecommerce ./ecommerce
