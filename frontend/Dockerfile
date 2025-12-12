# --- Étape 1 : Construction (Builder) ---
# Utilisation de Node 18+ 
FROM node:18-alpine AS builder

WORKDIR /app

# Installation des dépendances
COPY package*.json ./
RUN npm install

# Copie du code et construction de l'application
COPY . .
RUN npm run build

# --- Étape 2 : Serveur Web (Nginx) ---
FROM nginx:alpine

# Copie des fichiers construits (dossier dist/) vers le dossier public de Nginx
COPY --from=builder /app/dist /usr/share/nginx/html

# Le port par défaut de Nginx est 80
EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]