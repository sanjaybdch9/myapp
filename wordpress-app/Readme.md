# Nextjs application

1. Create Docker Secrets:
Create a secret for the MySQL root password and another for the WordPress database password. Open a terminal or command prompt and execute the following commands:

echo "your_mysql_root_password" | docker secret create mysql_root_password -
echo "wordpress_password" | docker secret create wordpress_db_password -

Replace your_mysql_root_password and wordpress_password with your actual passwords

2. Deploy and Run the Stack:
With the docker-compose.yml file, deploy and run your Docker stack using below command

docker stack deploy -c docker-compose.yml my-wordpress

This command will deploy the WordPress application stack using Docker secrets for sensitive data.


To access the application from browser use localhost: http://localhost:8081

