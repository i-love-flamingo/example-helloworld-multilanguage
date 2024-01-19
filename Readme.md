## Overview
This is a demo of a minimal multi-language Flamingo application. It can be used to kickstart a new project.

If you are not familiar with Flamingo framework, consult the following links to get a basic understanding of multi language setup.

- [Flamingo Bootstrap](https://docs.flamingo.me/2.%20Flamingo%20Core/1.%20Flamingo%20Basics/7.%20Flamingo%20Bootstrap.html)
- [Localization](https://github.com/i-love-flamingo/flamingo/tree/master/core/locale)        
- [Prefix Router](https://github.com/i-love-flamingo/flamingo/tree/master/framework/prefixrouter)  

### Prerequisites
- Go >= 1.21.4
- nodejs >= 18
- Make (build tool)

### Setup
```shell
make all
```

### Build frontend
Must be triggered after making changes in the frontend code. It resides in the top level `frontend` directory.

```shell
make frontend
```

### Prepare Translations

```shell
make translations
```

### Compile and start the application

```shell
make run
```
The application server listens on port `3210`. Use http://localhost:3210 in your browser to access the app.