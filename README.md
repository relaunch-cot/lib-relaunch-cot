# ReLaunch Core Library

## Sobre o Projeto

Biblioteca compartilhada desenvolvida em Go para centralizar e reutilizar código comum entre múltiplos microserviços. O objetivo principal é promover consistência, reduzir duplicação de código e facilitar a manutenção em arquiteturas distribuídas.

## Principais Recursos

### Funcionalidades de Autenticação e Usuários
- Gerenciamento completo de usuários (cadastro, login, logout)
- Redefinição e recuperação de senha via email
- Personalização de perfil e configurações
- Exclusão de conta

### Sistema de Comunicação
- Gerenciamento de chats entre usuários
- Envio e recuperação de mensagens
- Sistema de notificações em tempo real

### Gestão de Projetos
- Criação e busca de projetos
- Vinculação de freelancers a projetos
- Listagem de projetos disponíveis

### Recursos Adicionais
- Exportação de relatórios em PDF
- Sistema de notificações para eventos da plataforma

## Arquitetura e Padrões de Projeto

A biblioteca implementa diversos padrões de design para garantir qualidade e manutenibilidade:

**Padrões GoF:**
- **Singleton** - Instâncias únicas de recursos compartilhados
- **Adapter** - Integração com diferentes interfaces externas
- **Facade** - Simplificação de operações complexas
- **Strategy** - Algoritmos intercambiáveis
- **Factory** - Criação de objetos padronizada
- **Iterator** - Navegação em coleções

**Padrões Arquiteturais:**
- **Repository** - Abstração da camada de dados
- **Dependency Injection** - Desacoplamento e testabilidade

## Tecnologias

- **Go** - Linguagem de programação
- **Protocol Buffers** - Serialização de dados
- Arquitetura preparada para microserviços

## Objetivo

Manter uma base de código consistente e testada que possa ser facilmente integrada em diferentes serviços, acelerando o desenvolvimento e garantindo padrões de qualidade em toda a plataforma.