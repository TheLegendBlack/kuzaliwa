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
ignite scaffold chain github.com/TheLegendBlack/kuzaliwa --address-prefix kuz
```
 Un dossier nommé 'kuzaliwa' se créé alors dans notre répertoire. On accède à ce dossier : `cd kuzaliwa`
 
 Une fois dans ce dossier on va taper successivement les commandes suivantes, afin le sous-répertoire bin de l'espace de travail à notre PATH (PS : Lorsqu'on ne le fait pas, le déploiement de la blockchain en locale ne marche pas.) :
 ```
$ export PATH=$PATH:$(go env GOPATH)/bin
$ export GOPATH=$(go env GOPATH)
```
Une fois cela fait, il nous reste qu'une seule chose à faire avant de déployer la blockchain en local, on va substituer le fichier config.yml du dossier 'kuzaliwa' par celui du github (on peut soit faire un copier-coller  et remplacer ou bien supprimer celui existant et le remplacer par celui du github).

## IV - Déploiement de la blockchain en local
 
Afin de déployer la blockchain en local, on va utiliser la commande `serve`, on va taper :
 ```
ignite chain serve
```
Une fois cette commande exécuter, la blockchain en développement est déployé en local, le terminal nous affiche les informations concernant les comptes lors de la 1ère exécution et nous affichent les liens suivants : 
```
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500
```
PS : Ils ne sont accessibles que lorsque la blockchain est déployée.
Une fois que le déploiement est exécuté, on peut plus modifier la blockchain, pour stopper ce processus et revenir à un état de modification, il faut taper `control + c` 

## V - Les commandes
 La principale commande est la commande `ignite`, elle précède la majorité des autres commandes et s'utilise de la façon suivante :
 ```
Usage:
  ignite [command]

Available Commands:
  scaffold    Scaffold a new blockchain, module, message, query, and more
  chain       Build, initialize and start a blockchain node or perform other actions on the blockchain
  generate    Generate clients, API docs from source code
  network     Launch a blockchain in production
  node        Make calls to a live blockchain node
  account     Commands for managing Ignite accounts
  relayer     Connect blockchains by using IBC protocol
  tools       Tools for advanced users
  docs        Show Ignite CLI docs
  version     Print the current build information
  help        Help about any command
  completion  Generate the autocompletion script for the specified shell
```
Afin de savoir à quoi sert une commande, il suffit de la faire succéder par `--help`:
```
Flags:
  -h, --help   help for ignite

Use "ignite [command] --help" for more information about a command.
```

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
