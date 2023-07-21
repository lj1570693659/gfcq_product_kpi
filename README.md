# gfcq_product_kpi

docker run -p 3307:3306 --name mysql8 -e MYSQL_ROOT_PASSWORD=123456 -d mysql:8

* docker build -t gfcq_product_kpi_images -f .\manifest\docker\Dockerfile .
* docker run -itd -p 8199:8199  --name gfcq_product_kpi  gfcq_product_kpi_images