# --- Étape 1 : Construction (Builder) ---
# On utilise l'image Go officielle version 1.23 comme demandé [cite: 206]
FROM golang:1.23-alpine AS builder

WORKDIR /app

# On copie les fichiers de gestion des dépendances
COPY go.mod go.sum ./
RUN go mod download

# On copie tout le code source
COPY . .

# On compile l'application. Le point d'entrée est cmd/server/main.go [cite: 168]
RUN go build -o main cmd/server/main.go

# --- Étape 2 : Exécution (Runner) ---
FROM alpine:latest

WORKDIR /app

# Création d'un utilisateur non-root pour la sécurité 
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# On copie l'exécutable compilé depuis l'étape précédente
COPY --from=builder /app/main .
# On copie aussi les migrations et scripts nécessaires au démarrage [cite: 453]
COPY --from=builder /app/migrations ./migrations
COPY --from=builder /app/scripts ./scripts

# On définit les permissions et on change d'utilisateur
RUN chown -R appuser:appgroup /app
USER appuser

# Le serveur écoute sur le port 8080 par défaut [cite: 178]
EXPOSE 8080

CMD ["./main"]