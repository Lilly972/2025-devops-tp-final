# Étape 1: Build des fichiers statiques du Frontend (Utilise Node 20 comme requis par Storybook)
FROM node:20-alpine AS build

WORKDIR /app
COPY package.json package-lock.json ./
RUN npm install
COPY . .
# Commande Storybook pour créer les fichiers statiques (le --url-path est géré via package.json)
RUN npm run build-storybook

# Étape 2: Serveur de production Nginx
FROM nginx:alpine

# Copier la configuration Nginx customisée
# Cela remplace le fichier de configuration par défaut de Nginx
# Le chemin "../nginx/nginx.conf" est correct si le Dockerfile est dans 'frontend'
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copier les fichiers statiques du build Storybook (le contenu de storybook-static)
COPY --from=build /app/storybook-static /usr/share/nginx/html

# Le conteneur Nginx écoute le port 80 par défaut