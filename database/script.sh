#DIR /database/ (*)
mongodump --db ecommerce -o ./

mongorestore --drop --db ecommerce ./ecommerce 
