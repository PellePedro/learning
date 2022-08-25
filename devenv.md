### Design Environment with VScode

This tutorial will show a technical previw of integrating skyramp into the service development workflow.

We identify the requirements on a design environment for microservices as:

- Mocking open API interfaces (rest/grps/thrift) with Skyramp MockWorker.
- Development with all tools in docker containers provides cositent and immutable tooling with the dev organization. 
- Development and debugging tools at runtime (breakpoints, single stepping)


## Containerized Development Tools (devcontainers).
This tutorial will explore recent trends in development tools, which is to support code development inside containers e.g:
<BR/>[AmbassadorTelepresence](https://www.getambassador.io/products/telepresence/)
<BR/>[Gitlab Codespaces](https://github.com/features/codespaces)

The moitvation to support VSCode and devcontainers:
- Only a container runtime (Docker) and a coding/debugging UI (VScode) is required to be installed on local machine. 
- Standardized development environment where tools (e.g. nodejs, python, dotnet, go, java) are packaded in immutable container images.
- Supports Development and Debugging on a production like environmane


## Host Requirements
- VScode with extension for remote containers
- Docker
- Access to skyramp worker container image

## Demo Scenario 


### GCP Microservices (Recommendation Service)




## Docker compose

```
services:
  # The service beeing developed using VScdode remote extension and developer container suport
  dev:
    build:
      context: ./.devcontainer
      dockerfile: Dockerfile
    #image: docker.io/pellepedro/recommendationservice:v0.0.1
    volumes:
      - .:/workspace:cached
      - ~/.ssh:/root/.ssh:z
    entrypoint: [ "sh", "-c", "while sleep 3600000; do :; done" ]

  # The product-catalog-service is mocked by the Skyramp worker to  
  product-catalog-service:
    image: 296613639307.dkr.ecr.us-west-2.amazonaws.com/dev/worker:5c20658
    volumes:
      - ../project/thrift:/skyramp/project
    entrypoint: [ "/worker", "-config", "/skyramp/project/containers/product-catalog-service.yaml" ]
```


### Productcatalog Service - Container desription

```
containers:
  - name: product-catalog-service
    image:
      repository: docker.io/pellepedro/productcatalogservice
      tag: v0.0.2
    ports:
      - product-catalog-service-port50000

ports:
  - name: product-catalog-service-port50000
    endpoints:
    - ProductCatalogService
    secure: true
    thrift-options:
      protocol: "binary"
      buffered: true
      framed: false
    protocol: thrift
    port: 50000
endpoints:
  - name: ProductCatalogService
    defined:
      name: ProductCatalogService
      path: "thrift/demo.thrift"
    methods:
      - name: GetProducts
        output: GetProductsResponse
      - name: ListProducts
        output: ListProductsResponse
      - name: SearchProducts
        output: SearchProductsResponse
signatures:
  - name: GetProductsResponse
    fields:
    - name: result
      type:  struct
      default: |
            {
                "id": "OLJCESPC7Z",
                "name": "Sunglasses",
                "description": "Add a modern touch to your outfits with these sleek aviator sunglasses.",
                "picture": "/static/img/products/sunglasses.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 19,
                    "nanos": 990000000
                },
                "categories": [
                    "accessories"
                ]
            },
  - name: ListProductsResponse
    fields:
    - name: result
      type:  struct
      default: |
          [
            {
                "id": "OLJCESPC7Z",
                "name": "Sunglasses",
                "description": "Add a modern touch to your outfits with these sleek aviator sunglasses.",
                "picture": "/static/img/products/sunglasses.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 19,
                    "nanos": 990000000
                },
                "categories": [
                    "accessories"
                ]
            },
            {
                "id": "66VCHSJNUP",
                "name": "Tank Top",
                "description": "Perfectly cropped cotton tank, with a scooped neckline.",
                "picture": "/static/img/products/tank-top.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 18,
                    "nanos": 990000000
                },
                "categories": [
                    "clothing",
                    "tops"
                ]
            },
            {
                "id": "1YMWWN1N4O",
                "name": "Watch",
                "description": "This gold-tone stainless steel watch will work with most of your outfits.",
                "picture": "/static/img/products/watch.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 109,
                    "nanos": 990000000
                },
                "categories": [
                    "accessories"
                ]
            },
            {
                "id": "L9ECAV7KIM",
                "name": "Loafers",
                "description": "A neat addition to your summer wardrobe.",
                "picture": "/static/img/products/loafers.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 89,
                    "nanos": 990000000
                },
                "categories": [
                    "footwear"
                ]
            },
            {
                "id": "2ZYFJ3GM2N",
                "name": "Hairdryer",
                "description": "This lightweight hairdryer has 3 heat and speed settings. It's perfect for travel.",
                "picture": "/static/img/products/hairdryer.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 24,
                    "nanos": 990000000
                },
                "categories": [
                    "hair",
                    "beauty"
                ]
            },
            {
                "id": "0PUK6V6EV0",
                "name": "Candle Holder",
                "description": "This small but intricate candle holder is an excellent gift.",
                "picture": "/static/img/products/candle-holder.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 18,
                    "nanos": 990000000
                },
                "categories": [
                    "decor",
                    "home"
                ]
            },
            {
                "id": "LS4PSXUNUM",
                "name": "Salt \u0026 Pepper Shakers",
                "description": "Add some flavor to your kitchen.",
                "picture": "/static/img/products/salt-and-pepper-shakers.jpg",
                "price_usd": {
                    "currency_code": "USD",
                    "units": 18,
                    "nanos": 490000000
                },
                "categories": [
                    "kitchen"
                ]
            },
            {
                "id": "9siqt8tojo",
                "name": "bamboo glass jar",
                "description": "this bamboo glass jar can hold 57 oz (1.7 l) and is perfect for any kitchen.",
                "picture": "/static/img/products/bamboo-glass-jar.jpg",
                "price_usd": {
                    "currency_code": "usd",
                    "units": 5,
                    "nanos": 490000000
                },
                "categories": [
                    "kitchen"
                ]
            },
            {
                "id": "6e92zmyyfz",
                "name": "mug",
                "description": "a simple mug with a mustard interior.",
                "picture": "/static/img/products/mug.jpg",
                "price_usd": {
                    "currency_code": "usd",
                    "units": 8,
                    "nanos": 990000000
                },
                "categories": [
                    "kitchen"
                ]
            }
          ]
  - name: SearchProductsResponse
    fields:
    - name: result
      type:  struct
      default: |
        [
            {
                "id": "9siqt8tojo",
                "name": "bamboo glass jar",
                "description": "this bamboo glass jar can hold 57 oz (1.7 l) and is perfect for any kitchen.",
                "picture": "/static/img/products/bamboo-glass-jar.jpg",
                "price_usd": {
                    "currency_code": "usd",
                    "units": 5,
                    "nanos": 490000000
                },
                "categories": [
                    "kitchen"
                ]
            },
            {
                "id": "6e92zmyyfz",
                "name": "mug",
                "description": "a simple mug with a mustard interior.",
                "picture": "/static/img/products/mug.jpg",
                "price_usd": {
                    "currency_code": "usd",
                    "units": 8,
                    "nanos": 990000000
                },
                "categories": [
                    "kitchen"
                ]
            }
        ]

```
