# Étape 1: Build des fichiers statiques du Frontend
FROM node:20-alpine AS build

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
RUN npm run build-storybook

# Étape 2: Serveur de production Nginx
FROM nginx:alpine

# Copier la configuration Nginx customisée
# Le fichier est maintenant dans le même répertoire que le Dockerfile (frontend/)
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copier les fichiers statiques du build Storybook
COPY --from=build /app/storybook-static /usr/share/nginx/html