# Integração de Leads Grupo ZAP

Estamos disponibilizando neste repositório o manual de Integração de Leads do Grupo Zap. Este documento contém as informações básicas para os Softwares/CRM realizarem as integracões com nosso sistema afim de receber os leads gerados nos portais do GrupoZap (VivaReal e ZapImoveis).
Basicamente nosso sistema de integração de leads precisa de uma URL (Endpoint) para enviar as informações do lead para os softwares homologados, esta requisição será realizada via protocolo HTTP no verbo POST no Endpoint especificado passando um JSON como corpo do request. Abaixo exemplos de endpoints que poderão ser utilizados: 

Endpoint:

https://yourdomain.com.br/grupozap/lead/ID-DO-ANUNCIANTE

Ex:
https://crm-ou-imobiliaria.com.br/grupozap/lead/123456

OU

https://ID-DO-ANUNCIANTE.yourdomain.com.br/grupozap/lead

Ex:
https://imobiliariaxpto.crm.com.br/grupozap/lead

ID-DO-ANUNCIANTE é parte da URL que representa o identificador do anunciante no sistema que irá receber a requisição. É uma forma de identificação do cliente xpto no Software/CRM.

Caso não tenha ou não queira especificar o ID-DO-ANUNCIANTE o endpoint poderá ser implementado sem este identificador

Ex: 
https://yourdomain.com.br/grupozap/lead

Utilizando um único endpoint, o Software/CRM deverá identificar o cliente pelo `clientListingId` mais detalhes abaixo neste manual:

### Envio dos Leads
Os leads serão enviados via protocolo HTTP no verbo POST passando um JSON com as informações do lead no endpoint (URL) especificado e homologado por nosso time.

O processo de integração será acionado para cada lead individualmente, sempre que recebermos um contato do usuário. Nosso controle de sucesso do recebimento dos leads será feito através dos [códigos de status](https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html) do protocolo HTTP da seguinte forma:

2xx: indica que o cliente recebeu o lead com sucesso e que ações adicionais não são necessárias. Qualquer código de status da família 200 terá esse efeito.
3xx, 4xx ou 5xx: indica que houve falha no processamento do lead por parte do cliente.

Caso o endpoint do cliente retorne qualquer código de status que não os da família 200, haverá retentativa automática no envio do lead. O reenvio será tentado por 5 vezes em intervalos de 10 minutos e após isso ele será armazenado temporariamente por até 14 dias, podendo ser reenviado a pedido do cliente.

O controle do status de recebimento dos leads será feito exclusivamente através dos códigos de status do protocolo HTTP, qualquer informação enviada no corpo da resposta será totalmente ignorada.

Contrato JSON:
```
{
  "leadOrigin": "VivaReal",
  "timestamp": "2017-10-23T15:50:30.619Z",
  "originLeadId": "59ee0fc6e4b043e1b2a6d863",
  "originListingId": "87027856",
  "clientListingId": "a40171",
  "name": "Nome Consumidor",
  "email": "nome.consumidor@email.com",
  "ddd": "11",
  "phone": "999999999",
  "phoneNumber": "11999999999",
  "message": Olá, tenho interesse neste imóvel. Aguardo o contato. Obrigado.",
}
```

Onde:

- **leadOrigin**: Identificador do portal que o lead foi gerado (VivaReal, Zap);
- **timestamp**: Data e horário da criação do lead no formato ISO_LOCAL_DATE_TIME;
- **originLeadId**: Identificador do lead do GrupoZap;
- **originListingId**: Identificador do anúncio do GrupoZap;
- **clientListingId**: Identificador do anúncio para o anunciante (ExternalId);
- **name**: Nome do consumidor que gerou o lead;
- **email**: E-Mail do consumidor que gerou o lead;
- **ddd**: DDD do telefone do consumidor que gerou o lead Ex:`11`;
- **phone**: Telefone do consumidor que gerou o lead Ex:`999999999`;
- **phoneNumber**: [deprecado] Telefone do consumidor que gerou o lead (DDD + Phone) Ex:`11999999999`;
- **message**: Mensagem do consumidor que gerou o lead;

### Observações Importantes:

#### leadOrigin
Interprete este campo para fazer a distinção do portal onde o lead foi gerado, os valores são:

- `Zap`: Lead gerado no portal https://zapimoveis.com.br
- `VivaReal`: Lead gerado no portal https://vivareal.com.br

Ou seja, não será necessário criar 2 endpoints individualmente para cada portal, somente interprete o campo `leadOrigin` e faça a distribuição do lead a partir do portal específico.

#### clientListingId
O campo `clientListingId` é único identificador do anúncio/empreendimento conhecido pelo Software/CRM e GrupoZap, este identificador chega até nós via carga Feeds como codigo externo (externalId) ou customizado diretamente nos Software/CRM pelo cliente. Considere este identificador para realizar a devida associação do lead com o anúncio/empreendimento do cliente. Caso este identificador não estiver contído na requisição, deverá retornar o statusCode da família 400 (4xx) para ser analisado e reprocessado futuramente.

### Timeout
A requisição POST para o endpoint está configurada com timeout de 5 segundos, ou seja, qualquer requisição que demorar mais que 5 segundos será considerada ERRO sendo reenviadas de acordo com nossas regras de retentativas.

### Segurança
Para verificar a confiabilidade do emissor das requisições, será enviado no Header uma chave de segurança no padrão [Basic Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication). A chave a ser verificada será no seguinte formato:

Exemplo:
```
Authorization: Basic dml2YXJlYWw6NTk0RjgwM0IzODBBNDEzOTZFRDYzRENBMzk1MDM1NDI=
```
Para saber mais sobre como implementar a segurança [acesse a wiki](https://github.com/grupozap/crm-lead-integration/wiki/Como-Validar-Segurança-SECRET_KEY).

### Testes
Assim que as implementações forem devidamente realizadas e estiver pronto para iniciar os testes, entre em contato com: <p><a href="mailto:integracaoleads@grupozap.com">integracaoleads@grupozap.com</a></p>

### Arquivo de Importação de Anunciantes em Lote
Após homologação da Integração de Leads, enviar todos os clientes em um arquivo CSV contendo os seguintes campos: `DOCUMENT`,`ORIGIN`,`URL`,`SEND_LEAD_EMAIL`

- **DOCUMENT:** CPF ou CNPJ do cliente **sem máscara/formatação** [text];
- **ORIGIN:** Portal que será configurado a integração. Permitido valores: VIVAREAL ou ZAP [text];
- **URL:** Url de integração do cliente [text];
- **SEND_LEAD_EMAIL:** Flag para configurar se o lead também será enviado por e-mail [boolean];

#### Exemplo:
```
DOCUMENT,ORIGIN,URL,SEND_LEAD_EMAIL
56921485000100,VIVAREAL,https://crm.com.br/lead,false
56921485000100,ZAP,https://crm.com.br/lead,false
```
Enviar em anexo o arquivo csv para o e-mail: <p><a href="mailto:integracaoleads@grupozap.com">integracaoleads@grupozap.com</a></p>

### Dúvidas Sugestões ou Problemas
Caso tenha alguma dúvida, sugestão ou problemas durante a implementação da integração de Leads, abra um [Issue](https://github.com/grupozap/crm-lead-integration/issues) neste repositório que iremos responder assim que possível.

