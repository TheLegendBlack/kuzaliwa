## I - Installer GO
On commence par installer go :
https://go.dev/dl/

## II - Installer IGNITE CLI
Ensuite, on va installer Ignite CLI.
Personnellement, je me suis créé un dossier 'Blockchain' sur mon bureau et c'est dans ce dossier que je l'ai installé et qu'ensuite j'ai créé la blockchain.
On va ouvrir le terminal, acccéder au répertoire choisi pour l'installation et taper :
```
curl https://get.ignite.com/cli | bash
```
Une fois l'installation terminée, on tape :
```
sudo mv ignite /usr/local/bin/
```
## III - Création de la blockchain en local
 PS : Je me suis rendu compte que lorsque l'on créé une blockchain en utilisant le lien du github, certaines informations sont réinitialisées, donc j'ai adapté la procédure pour que vous ayez les mêmes informations que moi
 
 Pour créer la blockchain en local, on va taper la commande suivante au sein du répertoire dans lequel on a installé Ignite :
 ```
ignite scaffold chain kuzaliwa --address-prefix kuz
```
 
`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/Kuzaliwa@latest! | sudo bash
```
`username/Kuzaliwa` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)
