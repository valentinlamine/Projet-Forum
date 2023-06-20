# **Projet-Forum :** 

Ce projet a été réalisé par Valentin Lamine, Mattéo Vocanson, Noa Gambey et Adrien Gendron dans le cadre d'un devoir scolaire visant à créer un forum en utilisant le langage de programmation Go et SQL pour le backend et HTML/CSS/JavaScript pour le frontend.

# **Utilisation du projet :**

Pour accéder au projet, vous devez d'abord avoir WAMP installé et lancé sur votre machine. Vous devez ensuite importer la base de données "forum" dans phpMyAdmin. Le fichier forum.sql se trouve dans le dossier [docs](https://github.com/valentinlamine/Projet-Forum/tree/main/docs).

Une fois le forum importé et WAMP lancé :

1-Accédez au dossier [src](https://github.com/valentinlamine/Projet-Forum/tree/main/src).  
2-Exécutez la commande `go run main.go.`

```
cd src
go run main.go
```

# **Organisation des dossiers :**
-Le projet est organisé en plusieurs dossiers :

## **dossier .idea**
-Contient les fichiers de configuration pour l'IDE.

## **dossier docs**
-Contient divers documents relatifs au projet, y compris forum.sql qui est utilisé pour configurer la base de données du forum.

## **dossier src**
-Le dossier [src](https://github.com/valentinlamine/Projet-Forum/tree/main/src) contient les fichiers sources du projet, y compris :

-[main.go](https://github.com/valentinlamine/Projet-Forum/blob/main/src/main.go) qui est le point d'entrée principal du projet.  
-Dossier [CSS](https://github.com/valentinlamine/Projet-Forum/tree/main/src/CSS) qui contient tous les fichiers de style pour le frontend du site.  
-Dossier [JS](https://github.com/valentinlamine/Projet-Forum/tree/main/src/JS) qui contient tous les scripts JavaScript pour le frontend du site.  
-Dossier [Utilities](https://github.com/valentinlamine/Projet-Forum/tree/main/src/Utilities) qui contient le Golang organiser pour le projet.  
-Dossier [assets](https://github.com/valentinlamine/Projet-Forum/tree/main/src/assets) qui contient des fichiers multimédias tels que des images.  
-Dossier [template](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template) qui contient les différents modèles de pages du site.   
Ce dossier est lui-même divisé en deux sous-dossiers :  
-[base](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base) qui contient les templates de base:  
-[connected](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base/connected) qui contient les templates pour les utilisateurs connectés.  
-[disconnected](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/base/disconnected) qui contient les templates pour les utilisateurs non connectés.  
-Dossier [componants](https://github.com/valentinlamine/Projet-Forum/tree/main/src/template/componants) qui contient des fragments HTML réutilisables tels que des messages, des en-têtes, etc.  

## **Autres fichiers :**

-[go.mod](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.mod) et [go.sum](https://github.com/valentinlamine/Projet-Forum/blob/main/src/go.sum) sont des fichiers générés automatiquement par Go pour gérer les dépendances du projet.  
-README.md est ce fichier, qui fournit des informations sur le projet et comment l'utiliser.  
-forum.sql est le fichier qui permet de créer la base de données pour le forum.  

# **Répartition des tâches :**

### Noa Gambey : 
-Transition SQL-Go  
-Creation des structs Go  
-Formatage des données  
-Creation de toutes les fonctions Go pour chaque tables de la BDD  

### Mattéo Vocanson : 
-Creation des nouveaux tags pdt la création d'un topic  
-Systeme login v1 (a étais refait)  
-Le routage (mux)  
-Le go dans l'html  
-Suppression des messages du bots quand un premier message est entré dans le forum  
-Passer la data du go a l'html  
-Systeme d'envoi de message  
-La securité du login (regex) 
-Pannel Admin (ajouter role a utilisateur, reinitialiser role utilisateur, crée un role, supprimer un role, affichage des messages report) 

### Adrien Gendron :
-Frontend (HTML,CSS,JS)  
-Responsive   
-Ajustement CSS   
-README  
-Documentation  

### Valentin Lamine :
-Figma  
-Login/register (v2 avec fetch)  
-Stockage de l'id dans le localstorage  
-Fonction following/popular  
-Algorithme de tri des topics  
-Css pannel admin  
-Debug global  