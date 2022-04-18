# trackingapp

Essa aplicação tem como objetivo receber os acessos feitos à aplicação e armazenar as informações do usuário, caso seja o seu primeiro acesso.

## Iniciand a aplicação

Para iniciar a aplicação, certifique-se de que está no diretório de nível mais alto do projeto, no mesmo nível do arquivo
docker-compose.yml. Em seguide rode o comando:
`sudo docker-compose up`
Esse comando iniciará a aplicação principal, um banco mongodb, um proxy nginx e uma rede de conexão entre os dois. A aplicação é
acessível somente através da porta 80.

A aplicação conta também com um arquivo `.env`, o qual aqui não está no `.gitignore`. Nesse arquivo `.env` estão
contidas as credenciais do banco de dados. Se houver necessidade de alterá-las, também é necessário alterar as variáveis
utilizadas na inicializão do banco mongo diretamente no arquivo docker-compose.yml

## Utilizando a aplicação

A aplicação conta com apenas um endpoint, com os métodos Get e Post.
Localmente
Endpoint: `http://localhost:80/access`
Remotamente (esse é um ip público pode será alterando quando a máquina for reiniciada. Caso não esteja acessível, favor contactar-me)
Endpoint: `http://ec2-15-228-14-183.sa-east-1.compute.amazonaws.com:80/access`
POST: retorna um json com os dados do usuário em questão. Caso o usuário já exista no banco de dados, suas informações serão buscadas no mesmo,
caso contrário, suas informações serão salvas. Em ambos os casos o payload de retorno contém apenas as infomações do usuário referentes ao seu User-Agent,
seu "clientId", que é utilizado para identificá-lo unicamente e seu endereço de IP. As demais informações referentes a um acesso são descartadas.

Exemplo de payload a ser enviado:

{
"ip":"192.92.168.2",
"ua":"Mozilla/5.0",
"utm":"utm_medium=facebook",
"url":"https://google.com",
"sessionId": "12345",
"clientId":253
}

Exemplo de resposta:

{
"clientId": 253,
"ip": "192.92.168.2",
"ua": "Mozilla/5.0"
}

Caso você queira atestar a capacidade da aplicação em lidar com diversas requisições, há um arquivo chamado "testApplication.py" dentro do folder "testApplication.
Ele vai mandar 500 requisições para a aplicação. Caso ele não funcione, é possível que seja porque o ip público da aplicação foi alterado, nesse caso, entre em contato comigo para eu informar o novo. Para rodá-lo, utilize o comando: "python3 testApplication" a partir do folder "testApplication". É necessário que você tenha
os pacotes presentes dentro do requirements.txt instalados para rodar a rotinha.

## Verificações feitas pela aplicação

1 - É verificado se o IP recebido na requisição é um ip válido. Caso não seja, o cliente será informado do mesmo.
2 - É verificado o user-agent no header da requisição. Se o user-agent em questão não corresponder a notebook, a um mobile ou a um table, ou caso ela corresponda a algum bot, a requisição será bloqueada e o cliente informado que um bot foi detectado.

## Testes

Foram criados testes para apenas 2 funções, referentes às rotinas encontradas no folder "utils". Para rodá-los, acesse o folder `shared/utils/` e rode o comando
go test. As funções testadas são a função "BlockBot" e "IsIpv4".

## Observações

A aplicação não foi capaz de atender a 1000 requisições simultâneas, o máximo que foi capaz de atender foi 500. Acredito que isso ocorreu pois a máquina utilizada armazenar a aplicação, o banco mongo e o nginx é uma t2.micro na Amazon, com apenas 1Gb de Ram e 2 CPU's e que uma máquina com mais memória RAM disponível e mais CPU's, a aplicação seria capaz de atender a todas as requisições. Foi utilizada uma t2.micro pois a aws me disponibilizou 750h de uso gratuíto desse tipo de máquina.

A aplicação responde em no máximo 50ms, dependendo se foi necessário armazenar as informações de um usuário no banco, ou apenas acessá-las.

A máquina utilizada foi alocada em São Paulo, para minimizar o tempo de latência de rede da aplicação.

Tentei inicialmente rodar a aplicação no Heroku, mas tive uma série de dificuldade impostas pela maneira como o Heroku exige que o desenvolvedor trabalhe.
1 - Não possui servidores na América do Sul, os mais próximos são nos Estados Unidos
2 - Não consegui rodar um container com mongodb sozinho dentro de uma das máquinas do Heroku, então teria que utilizar um banco auto-gerenciado deles, o que tornaria o projeto muito caro.

Não utilizei github actions para montar a aplicação. Para fazer o deploy simplesmente fiz o upload para uma máquina da Amazon e rodei os containers a partir de lá.

## Algumas considerações

Talvez seja por falta de conhecimento meu sobre o que seria uma "Api de Tracking", mas não ficou claro para mim que deveria ser feito com as informações do acesso, então simplesmente as ignorei. Como estava mais claro que o usuário deveria ser identificado se já tivesse acessado a aplicação anteriormente, ficou mais claro o que fazer com as informações referentes ao usuário.

Como não foi imposta nenhuma condição sobre como as informações do usuário chegariam, tomei liberdade de enviá-las como um json.

A aplicação devolve status 200 mesmo caso haja problemas de validação das informações. Um próximo passo seria devolver o status adequado.

Outro passo importante seria colocar mais testes unitários.

Não me importei com o uso de um gitflow durante a criação da aplicação.

Quando a aplicação fica muito tempo inativa, ela demora um pouco mais para responder, pouco mais de 100ms. As seguintes requisições são mais rápidas.

Tentei montar a aplicação utilizando banco de dados postgres, mas acredito que não é possível atender ao critério dos 100ms com esse banco de dados.
