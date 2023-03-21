# DeepFlower
- Glory to the goats!

### Docker-compose
Проверить хосты перед запуском ->
- goback\config\config.yaml 
host: 0.0.0.0
port: 8787
- vue3front\package.json
"scripts": {
    "preview": "vite preview --host 0.0.0.0 --port 5173",
  },
```sh
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```
### Local testing
1.(config/config.yaml)
```sh
apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
docker run --name myPostgresDb -p 32768:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -e POSTGRES_DB=deepflower -d postgres
```
2.
```sh
apt search golang-go
apt install golang-go 
go run cmd/deepflower/main.go
```
3. front: 
Используется Vite (Hot Module Replacement(HMR))
```sh
apt update
apt install nodejs npm
nodejs --version
npm install
npm run dev
```