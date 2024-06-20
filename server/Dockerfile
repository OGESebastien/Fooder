# Utiliser l'image officielle de Golang comme image de base
FROM golang:latest

# Définir le répertoire de travail à l'intérieur du conteneur
WORKDIR /app

# Copier les fichiers du projet dans le répertoire de travail
COPY . .

# Créer un fichier go.mod pour initialiser le module Go
RUN go mod tidy

# Installer les dépendances
RUN go install ./...

# Construire l'application
RUN go build -o main .

# Exposer le port sur lequel l'application sera disponible
EXPOSE 8080

# Commande pour exécuter l'application
CMD ["./main"]
