# Devops-Assignment-Inc42

My docker hub id : doc468991


## Go application

Image tag for Go-app : doc468991/ankush-go-app:v2-jenkins

Go application is exposed on port 8000 So to spin up the docker container use below command:

docker run -d -p 8000:8000 doc468991/ankush-go-app:v2-jenkins

To access the application from browser use localhost: http://localhost:8000

## Nextjs application

Image tag for nextjs-app : doc468991/ankush-nextjs-app:v2-jenkins

Nextjs application is exposed on port 3000 So to spin up the docker container use below command:

docker run -d -p 3000:3000 doc468991/ankush-nextjs-app:v2-jenkins

To access the application from browser use localhost: http://localhost:3000


### jenkins pipeline

we are using Jenkins pipeline script for building and deploying applications. This pipeline assumes that you have Jenkins installed and configured with the necessary plugins, such as the Docker Pipeline plugin , git plugin etc.

while pushing the docker image to docker hub using jenkins make sure to configure your docker hub credentials in jenkins. 
And make sure to replace the 'dockerhub' (shown in below script) with your own dockerhub_credentials which you have configured in jenkins.

// Push the Docker image to Docker Hub
script {
    docker.withRegistry('https://registry.hub.docker.com', 'dockerhub') {
        docker.image(DOCKER_IMAGE + ":" + DOCKER_TAG).push()
    }
}


# docker compose 

To Build all the application simultaneously run the docker-compose.yaml file present in docker-compose directory.

Use these commands to build the applications using docker-compose.yaml

docker compose up -d

To stop or remove all the builds usning docker-compose.yaml use below command:

docker compose down -d
