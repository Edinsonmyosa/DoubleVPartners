sudo docker build . -t prueba
sudo docker run -d --net="host" --restart=always  -e TZ=America/Bogota  --name prueba prueba