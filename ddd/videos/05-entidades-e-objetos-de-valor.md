# [Entidades e Objetos de Valor | DDD do jeito certo | Parte 05](https://www.youtube.com/watch?v=6Pjod34OCnE)

Entidades aparecem quando o especialista usa substantivos. Exemplo:  
Registrar um funcionário. Funcionário é uma entidade ou objeto de valor.

## Entidade
- São substantivos que podem ser representadas através de um ID.  
- Esse Id vai ser utilizado para comparar entidades.  
- Geralmente são mutáveis. Por exemplo, o endereço de um funcionário pode ser alterado.
- Entidades podem aparecer em bounded contexts diferentes, com podem ser representadas de maneiras diferentes, com atributos diferentes, mas sempre tendo o ID como ponto em comum.

## Objeto de valor 

- Sempre vive associado a uma Entidade. Por exemplo uma propriedade Endereço de um funcionário.
- Diferentemente de uma Entidade, a comparação com outros Objetos de valor deve ser feita através de todas as propriedades do mesmo ao invés de um ID.
- Propriedades de Objetos de valor são imutáveis.