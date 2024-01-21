# Company Microservice

This microservice handles company operations through a REST API, providing functionalities such as create, update, delete, and get (one). It utilizes PostgreSQL as the database.

## Dependencies

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

## Run Locally for Testing

1. Use the following command to run the app locally:

```bash
docker-compose up -d
```

2. A test user is created by default. Use the following commands to test:

   - Authenticate and get a token:

   ```bash
   curl -vvv -X POST http://localhost:8000/login -H 'Content-Type: application/json' -d '{"Username":"testuser","Password":"12345"}'
   ```

   The above command should return a valid token.

   - Create a company:

   ```bash
   curl -vvv --compressed -X POST http://localhost:8000/companies -H 'Content-Type: application/json' -H "Authorization: Bearer {token}" -d '{"name":"XM","description":"An interesting description","employees":100,"registered":false,"companyType":3}'
   ```

   - Retrieve the company:

   ```bash
   curl -vvv http://localhost:8000/companies/{id} -H "Authorization: Bearer {token}"
   ```

   - Update the company:

   ```bash
   curl -vvv -X PATCH http://localhost:8000/companies/{id} 'Content-Type: application/json' -H "Authorization: Bearer {token}" -d '{"id":"a1f5e7ab-8b2a-4f48-bab5-9de29c2638a2","name":"New Name","description":"5","employees":60,"registered":false,"companyType":0}'
   ```

   - Delete the company:

   ```bash
   curl -vvv -X DELETE http://localhost:8000/companies/{id} -H "Authorization: Bearer {token}"
   ```

## Run in Production

To build a Docker image for production, use the following command:

```bash
make image
```

Make sure to provide the following environment variables:

- `DB_HOST`
- `DB_PORT`
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`
- `JWT_KEY`

If you want to set up a default user in the database, also provide:

- `DEFAULT_USER`
- `USER_PASS`

A configuration file with parameters for the http server is also provided.