# [Agregados | DDD do jeito certo | Parte 06](https://www.youtube.com/watch?v=1AEOcQWQR2o)

Após mapear Entidades e Objetos de valor, é preciso mapear seus relacionamentos.

- Relações de composição.
- Relações de referência.

Exemplo:
Uma Entidade **Pedido** possui uma relação de composição com uma entidade Item de **Pedido**, pois um **Pedido** possui diversos Items.  
Já com uma Entidade **Cliente** existe uma relação de referência, pois o **Cliente** não faz parte do **Pedido**, mas está relacionado a ele.

Ao identificar essas relações, é possível começar a observar como será a feita a persistência dessas entidades.  
No exemplo de **Pedidos/Item De Pedido**, a transação responsável por persistir um **Pedido** é a mesma que persiste seus itens. Logo podemos considerar que eles fazem parte do mesmo **Agregado**.
Já o **Cliente** será persistido em uma outra transação, logo fará parte de outro agregado.

Logo, em DDD **Agregados** servem para identificar os limites transacionais. Apesar disso, não existe menção ao **Agregado** no código. **Não** criar uma classe como por exemplo `ClientAggregate`. Eles servirão para ajudar na criação da camada de `Repository`.

Em bancos NoSQL, normalmente, agregados são traduzidos para as coleções que criaremos. Por exemplo, o agregado de Pedidos dará o nome a coleção Pedido, que persistirá os dados das Entidades Pedido e Item do Pedido(em um array), do Objeto de valor Endereço, e uma referência ao ID do Cliente.