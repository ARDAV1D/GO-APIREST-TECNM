pipeline {
    agent any 

    tools {
        go "go"
    }

    parameters {
        string(name: 'API_CONTAINER_NAME', defaultValue: 'go_apirest-tecnm-restapi-1', description: 'Nombre del contenedor api de docker.')
        string(name: 'DB_CONTAINER_NAME', defaultValue: 'go_apirest-tecnm-some-postgres-1', description: 'Nombre del contenedor bd de docker.')
        string(name: 'API_IMAGE_NAME', defaultValue: 'apitecnm', description: 'Nombre de la imagen api docker.')
        string(name: 'DB_IMAGE_NAME', defaultValue: 'postgres', description: 'Nombre de la imagen bd docker.')
        string(name: 'API_IMAGE_TAG', defaultValue: 'v1', description: 'Tag de la imagen de la apirest.')
        string(name: 'DB_IMAGE_TAG', defaultValue: 'latest', description: 'Tag de la imagen de bd.')
        string(name: 'API_PORT', defaultValue: '8081', description: 'Puerto que usa el contenedor.')
    }

    stages {
        stage('Clonar repositorio') {
            steps {
                git branch: 'develop', url: 'https://github.com/ARDAV1D/GO-APIREST-TECNM.git'
            }
        }

        stage('Construir') {
            steps {
                sh 'go build -o go_apirest-tecnm'
            }
        }

        stage('Pruebas') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Deploy Containers') { 
            steps {
                script {
                    // Construir la imagen de la API
                    sh "docker build -t ${params.API_IMAGE_NAME}:${params.API_IMAGE_TAG} ."
                    
                    // Detener y eliminar cualquier contenedor existente de la API
                    sh "docker stop ${params.API_CONTAINER_NAME} || true"
                    sh "docker rm ${params.API_CONTAINER_NAME} || true"
                    
                    // Ejecutar el contenedor de la API
                    sh "docker run -d --name ${params.API_CONTAINER_NAME} -p ${params.API_PORT}:8081 ${params.API_IMAGE_NAME}:${params.API_IMAGE_TAG}"
                    
                    // Extraer la imagen de la base de datos (si es necesario)
                    sh "docker
