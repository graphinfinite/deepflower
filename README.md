# DeepFlower
Glory to the goats!

### Docker-compose
- TODO
```sh
docker-compose build
docker-compose up
```
### Local testing
1.(config/config.yaml)
```sh
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
docker run --name myPostgresDb -p 32768:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgrespw -e POSTGRES_DB=deepflower -d postgres
```
2.
```sh
sudo apt search golang-go
sudo apt install golang-go 
go run cmd/deepflower/main.go
```

3. front: 
```sh
sudo apt update
sudo apt install nodejs npm
nodejs --version
sudo npm install
npm run dev
```