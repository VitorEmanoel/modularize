# Modularize 

Modularize é um framework GoLang para desenvolver sistemas módulares

## Cliente
## Iniciar módulos
```shell
$: modularize start
```

### Iniciar modularize
#### Repositório Local
Iniciando um cliente modular local
```shell
$: modularize init 
```
Será criado um arquivo de configuração YAML para repositório local.
```yaml
modules: ./modules
extensions: ./extensions
```

#### Repositório Remoto
Existe a opção de repositórios remoto
```shell
$: modularize init --remote
```
Irá criar um arquivo de configuração para reposítorio remoto
```yaml
modules: 
  - name: TestModule
    version: v1.0.1
    url: https://repository.mobolife.com.br/modules/test-module
extensions:
  - name: TestExtension
    version: v1.0.1
    url: https://repository.mobolife.com.br/extension/test-extension
```
Ou
```yaml
modules:
  - name: TestModule
    version: v1.0.1
    path: ./modulesA/TestModule
extensions:
  - name: postgres
    version: v1.0.1
    path: ./extensionA/TestExtension
```
No segundo exemplo irá ser feito download diretamente do repositório declarado

### Baixar módulo
É possivel baixar um módulo utilizando o commando.
```shell
$: modularize get module repository.mobolife.com.br test-module 
```

E tambem baixar extension
```shell
$: modularize get extension repository.mobolife.com.br test-extension
```

Vale ressaltar que os módulos e extensões remotos não são baixados quando se executa o comando, são apenas validados, eles serão carregados em tempo de execução quando cliente inciado.

### Adicionar repositórios
É possivel adicionar repositorios para facilitar alguns comandos futuros
```shell
$: modularize add repository mobolife repository.mobolife.com.br
```

Com isso será adicionado a configuração o seguinte YAML
```yaml
repositories:
  - name: mobolife
    url: repository.mobolife.com.br
```
Isto irá facilitar quando nas proximas vezes que for baixar módulos
```shell
$: modularize get module mobolife test-module
```