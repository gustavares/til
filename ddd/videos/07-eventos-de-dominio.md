# [Eventos de Domínio | DDD do jeito certo | Parte 07](https://www.youtube.com/watch?v=_By3QRBMHSo)

Eventos de domínio podem ser pensados sob duas perspectivas, a do espaço do problema e do espaço da solução.

No **espaço do problema**, um evento é um registro de algo significativo. Por exemplo em um hospital quando ocorre a entrada de um paciente ou a alta de um paciente.

No **espaço da solução**, um evento é a ocorrência de uma mudança de estado, quando precisa ocorrer uma mudança ou alteração em uma entidade, por exemplo o salário de um Funcionário. Então o ideal é termos métodos que tornem explicitas essas mudanças para podermos lançar eventos correspondentes. No caso do funcionário, um método para aumentar o salário.

Eventos podem ser representados através de classes com nomes relevantes, por exemplo `EmployeeAddressChangedEvent`. Essa classe deve possuir uma propriedade para indicar quando o evento aconteceu e as propriedades relevantes. No caso de uma mudança de endereço, o novo endereço e talvez também o endereço antigo para comparação.

**DDD não é descritivo quanto a como endereçar a instância da classe de evento**, porém é ideal utilizar um mecânismo de mensageria. Como esse mecânismo funciona extrapola o contexto do DDD, pois trata de uma complexidade de solução técnica.

Algo para ficar atento, é que o contexto delimitado que produziu o evento não precisa necessariamente consumir a mensagem enviada, afinal isso provavelmente poderia ser resolvido dentro da mesma operação que produziu a mensagem.

