
## Componentes

### 1. [flask-api/app.py](webhook/flask-api/app.py)

- API em Python (Flask) que simula eventos de pagamento.
- Ao acessar o endpoint `/payment`, um evento é gerado e publicado no canal `payments` do Redis.
- Utiliza variáveis de ambiente para configurar o endereço do Redis e o endpoint de destino do webhook.

### 2. [gateway-payments/](webhook/gateway-payments/)

- Serviço em Go que consome eventos do canal `payments` no Redis.
- Os eventos são processados em uma fila e enviados como webhooks para o endpoint especificado.
- Implementa retentativas automáticas com backoff exponencial em caso de falha no envio do webhook.

#### Principais arquivos:

- [`cmd/main.go`](webhook/gateway-payments/cmd/main.go): Inicializa o serviço, conecta ao Redis e gerencia a fila de webhooks.
- [`queue/worker.go`](webhook/gateway-payments/queue/worker.go): Worker responsável por processar e reenfileirar webhooks em caso de falha.
- [`redis/redis.go`](webhook/gateway-payments/redis/redis.go): Lida com a assinatura do canal Redis e parsing dos eventos.
- [`sender/webhook.go`](webhook/gateway-payments/sender/webhook.go): Realiza o envio HTTP do webhook.

## Como executar

### Pré-requisitos

- Docker (opcional)
- Redis rodando em `localhost:6379`
- Python 3.x (para rodar o Flask)
- Go 1.21+ (para rodar o serviço Go)

### Passos

1. **Inicie o Redis** (localmente ou via Docker).
2. **Rode a API Flask**:
    ```sh
    cd webhook/flask-api
    pip install flask python-dotenv redis
    python app.py
    ```
3. **Rode o serviço Go**:
    ```sh
    cd webhook/gateway-payments
    go run ./cmd/main.go
    ```
   Ou via Docker:
    ```sh
    docker build -t gateway-payments .
    docker run --env-file .env gateway-payments
    ```

4. **Teste**:
    - Faça uma requisição GET para `http://localhost:8000/payment`.
    - O serviço Go irá consumir o evento e tentar enviar o webhook para o endpoint definido em `WEBHOOK_ADDRESS` no [.env](http://_vscodecontentref_/6) do Flask.

## Variáveis de ambiente

- `REDIS_ADDRESS`: Endereço do Redis (ex: `localhost:6379`)
- `WEBHOOK_ADDRESS` (Flask): URL para onde o webhook será enviado

## Observações

- O sistema implementa retentativas automáticas com backoff exponencial para falhas no envio do webhook.
- O endpoint de destino pode ser testado usando serviços como [webhook.site](https://webhook.site/).

---

Desenvolvido para estudo de filas, webhooks, goroutines e integração entre Python e Go.