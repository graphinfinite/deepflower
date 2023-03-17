# DeepFlower
Task manager

### Docker-compose
- TODO
### Local testing
1.(config/config.yaml)
```sh
docker run --name myPostgresDb -p 32768:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -e POSTGRES_DB=deepflower -d postgres
```
2.
```sh
go run cmd/deepflower/main.go
```

3. front: 
```sh
npm run dev
install node.js and npm
npm install create-vue@3.6.1
npm install vite
npm install vue-router@4
npm i the-new-css-reset
npm i --save axios
npm i @vueuse/core
```