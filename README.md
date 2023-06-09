# **Projet-Forum :**

Ce projet a été réalisé par [Valentin Lamine](https://github.com/valentinlamine/), [Mattéo Vocanson](https://github.com/matteoVcs), [Noa Gambey](https://github.com/NoaYnov) et [Adrien Gendron](https://github.com/AdrienGend) dans le cadre d'un devoir scolaire visant à créer un forum en utilisant le langage de programmation Go et SQL pour le backend et HTML/CSS/JavaScript pour le frontend.

# Présentation du projet :

Pour notre projet forum nous avons décidé de faire le parfait mixe entre deux célèbres forums, Reddit et StackOverflow. Nous récupérons les fonctionnalités et l'interface phare de Reddit avec un contenu lui qui est plus orienté vers StackOverflow. Notre forum est donc basé sur l'entraide et les communautés de reddit sur du contenu aussi bien développeur qu'annexe.

## Fonctionnalités du site :

* Algorithme de recommandation pour afficher une liste de topic populaire
* Possibilité d'afficher les topics suivis par l'utilisateur connecté (système de suivi de topic)
* Système de connexion et de création de compte
* Système de tag servant à être indexé dans le menu sur la gauche pour chercher par catégorie
* Possibilité d'afficher un topic ainsi que tous ces messages
* Possibilité d'interagir avec les messages (upvote, downvote, signalement)
* Possibilité de supprimer les topics / message pour les personnes ayant la permission
* Système de permission
* Politique d'utilisation et de vie privée
* Panel admin qui permet de :
  * Ajouter un rôle à un utilisateur
  * Réinitialiser les rôles d'un utilisateur
  * Créer un rôle
  * Supprimer un rôle
  * Afficher la liste des messages signalés
* Possibilité de créer des messages
* Possibilité de créer des topics

# **Lancement du projet :**

Pour accéder au projet, vous devez d'abord avoir WAMP d'installé sur votre machine ainsi que de l'avoir lancé. Vous devez ensuite importer [la base de données forum](data/forum.sql) dans phpMyAdmin.

Une fois le forum importé et WAMP lancé :

1. Accédez au dossier [src](https://github.com/valentinlamine/Projet-Forum/tree/main/src).
2. Exécutez la commande `go run main.go`

```bash
cd src
go run main.go
```

# **Organisation des dossiers :**

## **Dossier docs**

Contient tout ce qui à un lien de près ou de loin avec la documentation du projet ainsi que la maquette du site

## **Dossier src**

Le dossier src contient tous les fichiers sources nécessaires au fonctionnement du projet :

* Fichier [main.go](https://github.com/valentinlamine/Projet-Forum/blob/main/src/main.go) : fichier de lancement du projet
* Dossier [CSS](https://github.com/valentinlamine/Projet-Forum/tree/main/src/CSS) : contient tous les fichiers de style pour le frontend du site.
* Dossier [JS](https://github.com/valentinlamine/Projet-Forum/tree/main/src/JS) : contient tous les scripts JavaScript pour le frontend du site.
* Dossier [Utilities](https://github.com/valentinlamine/Projet-Forum/tree/main/src/Utilities) : contient l'intégralité du golang nécessaire au site (tout le code du backend).
* Dossier [assets](https://github.com/valentinlamine/Projet-Forum/tree/main/src/assets) : contient tous les fichiers externes nécessaires au site.
* Dossier [template](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template) qui contient tous les templates utilisés dans le site, il est lui même divisé en sous dossier :
  * [base](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base) qui contient la base des templates qui seront affichés à l'utilisateur lui même divisé en deux catégories :
    * [connected](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base/connected) qui contient les templates pour les utilisateurs connectés.
    * [disconnected](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base/disconnected) qui contient les templates pour les utilisateurs non connectés.
  * Dossier [componants](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/componants) qui contient des fragments HTML réutilisables tels que des messages, des en-têtes, etc.

## **Autres fichiers**

* [go.mod](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.mod) et [go.sum](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.sum) sont des fichiers générés automatiquement par Go pour gérer les dépendances du projet.
* README.md est ce fichier, qui fournit des informations sur le projet et comment l'utiliser.

# **Répartition des tâches :**

### Noa Gambey :

* Transition SQL-Go
* Création des structures Go
* Formatage des données
* Création de toutes les fonctions Go pour chaque tables de la BDD

### Mattéo Vocanson :

* Création des nouveaux tags pdt la création d'un topic
* Système login v1 (à étais refait)
* Le routage (mux)
* Le go dans l'html
* Suppression des messages du bot quand un premier message est entré dans le forum
* Passer la data du go a l'html
* Système d'envoi de message
* La sécurité du login (regex)
* Pannel d'administration (ajouter rôle a utilisateur, réinitialiser rôle utilisateur, crée un rôle, supprimer un rôle, affichage des messages report)

### Adrien Gendron :

* Frontend (HTML,CSS,JS)
* Responsive
* Ajustement CSS
* README

### Valentin Lamine :

* Maquette du site
* Création du script de la BDD
* Modification de la structure de la BDD
* Système de login et register (v2 avec fetch)
* Système de stockage de la session de connexion
* Système fetch et routage :
  * Upvote / Downvote
  * Report
  * Follow topic
* Système implémenté à 100 % (template, fetch, routage, traitement go, lien SQL, CSS)
  * Création de topic
  * Création de message
  * Delete message
  * Delete topic
* Fonction de formatage des dates
* Algorithme de tri par popularité des topics
* CSS panel admin
* Correction Responsive, README
* Aide / Debug global (HTML, CSS, JS, GO, SQL)
