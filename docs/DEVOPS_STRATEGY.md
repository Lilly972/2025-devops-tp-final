# Documentation DevOps – Christmas Gift List

## 1. Objectif du document

Ce document décrit la **stratégie DevOps complète** mise en place pour le projet *Christmas Gift List*, conformément aux exigences du **TP Final DevOps M1 Cyber**.

L’objectif est de permettre à :

* des développeurs,
* des membres non techniques (ex: investisseurs),
* et des enseignants

… de comprendre facilement comment le projet est versionné, testé, conteneurisé, déployé et maintenu, sans modifier le code applicatif

---

## 2. Contexte et besoins du projet

Le projet est une application **full-stack** permettant de gérer des idées de cadeaux de Noël.

Contraintes clés :

* Plusieurs personnes travaillent sur le projet
* Besoin de visualiser rapidement les nouvelles fonctionnalités
* Environnement partagé sans impacter les utilisateurs
* Code applicatif figé (interdiction de modification)

Une approche DevOps est donc nécessaire pour :

* automatiser les tests et déploiements
* standardiser les environnements
* faciliter la collaboration

---

## 3. Stack technique

### Frontend

* React 19 + TypeScript
* Vite
* Tailwind CSS
* Storybook (documentation UI)

### Backend

* Go (Chi router)
* PostgreSQL
* Migrations automatiques

### Tests

* Frontend : Vitest
* Backend : Go testing
* End-to-End : Playwright

### DevOps & Infrastructure

* GitHub (code source)
* GitHub Actions (CI/CD)
* Docker & Docker Compose
* Docker Hub (images)
* **Neon** (PostgreSQL managé)
* **Render** (déploiement backend)
* **Netlify** (hébergement frontend + reverse proxy)

---

## 4. Stratégie Git

### Repository

* Code hébergé sur **GitHub**

### Branches

* `main` : branche stable, prête pour la production
* `develop` : intégration des nouvelles fonctionnalités
* branches de feature : `feature/nom-feature`

### Workflow

1. Création d’une branche feature
2. Développement local
3. Pull Request vers `develop`
4. Tests automatiques via CI
5. Merge
6. Release automatique vers `main`

---

## 5. Conteneurisation avec Docker

### Objectifs

* Environnements reproductibles
* Lancement simple du projet
* Aucune dépendance locale requise

### Dockerfiles

#### Frontend

* Image basée sur Node
* Build du projet React
* Serveur Nginx pour le rendu statique
* Utilisation d’un **utilisateur non-root** pour la sécurité

#### Backend

* Image basée sur Go
* Compilation du binaire
* Exécution avec utilisateur non-root

### Images Docker

* Publiées automatiquement sur **Docker Hub** via CI/CD

---

## 6. Orchestration avec Docker Compose

Un fichier `docker-compose.yml` permet de lancer :

* frontend
* backend
* base de données PostgreSQL

### Avantages

* Démarrage en une seule commande
* Réseau interne Docker
* Variables d’environnement centralisées

Commande :

```bash
docker-compose up -d
```

---

## 7. Gestion des environnements

### Environnements supportés

* Local
* CI
* Production

### Variables d’environnement

#### Backend

* `DATABASE_URL`
* `PORT`

#### Frontend

* Gestion via reverse proxy (pas de modification du code)

Les secrets sont stockés dans **GitHub Secrets**.

---

## 8. Reverse Proxy et communication Front / Back

### Problème

Le frontend est une application statique et ne peut pas proxyfier `/api` en production.

### Solution mise en place

* Netlify est utilisé comme hébergeur frontend
* Mise en place de règles de redirection Netlify pour `/api/*`
* Les requêtes sont automatiquement transférées vers le backend hébergé sur Render



### Avantages

* Même origine pour le navigateur
* Aucun problème CORS
* Code frontend inchangé (exigence du TP)

(Détails complets dans `docs/reverse-proxy.md`)

---

## 9. CI/CD avec GitHub Actions

### Pipeline CI

Déclenché sur :

* push
* pull request

Étapes :

1. Installation des dépendances
2. Lint frontend et backend
3. Tests unitaires (Vitest + Go)
4. Tests E2E Playwright
5. Build des images Docker

### Pipeline CD

* Push automatique des images sur **Docker Hub**
* Déploiement du backend sur **Render**
* Déploiement du frontend sur **Netlify**
* Connexion automatique à la base PostgreSQL **Neon**
* Création automatique d’une **GitHub Release**

---

## 10. Documentation et Storybook

### Storybook

* Documentation visuelle des composants UI
* Accessible indépendamment du frontend
* Facilite la validation produit

### Documentation projet (`docs/`)

* `api.md` : documentation des endpoints
* `database.md` : schéma et migrations
* `reverse-proxy.md` : architecture réseau
* `devops.md` : stratégie DevOps (ce document)

---

## 11. Sécurité

Mesures mises en place :

* Utilisation d’utilisateurs non-root dans les conteneurs
* Isolation réseau Docker
* Secrets hors du code
* Aucune donnée sensible versionnée

---

## 12. Choix techniques et justification

| Élément          | Choix          | Justification                            |
| ---------------- | -------------- | ---------------------------------------- |
| CI/CD            | GitHub Actions | Intégration native GitHub                |
| Backend Hosting  | Render         | Déploiement simple des conteneurs Docker |
| Frontend Hosting | Netlify        | CDN performant + reverse proxy           |
| Base de données  | Neon           | PostgreSQL managé, sécurisé et scalable  |
| Conteneurs       | Docker         | Standard industriel                      |
| Orchestration    | Docker Compose | Simplicité et reproductibilité           |
| Registry         | Docker Hub     | Public et simple                         |



Ces évolutions sont envisageables sans remise en cause de l’architecture actuelle.


## 14. Conclusion

Cette stratégie DevOps permet :

* un déploiement fiable et automatisé
* une collaboration efficace
* une application facilement démontrable

