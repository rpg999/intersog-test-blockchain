# Test project

This test code is designed to create a simple Restful API back end for basic blockchain management application.
    
### Checkout sources

To get the latest stable version:
```
git clone https://github.com/rpg999/intersog-test-blockchain.git
cd intersog-test-blockchain
make init
```

### Build
```
make build
```

### Run Daemon
```
make run
```

### API Description
1. Create a blockchain 
    - URL: /api/chain/create
    - Accept method: POST
    - Request body:
    ```
    {
        "name" : "blockname"
    }
    ```
    - Response body:
    ```
    {
        "id": 3
    }
    ```
    
2. List all blockchains
    - URL: /api/chain/list
    - Accept method: GET
    - Response body:
    ```
    {
        "chains": [
            {
                "id": 1,
                "name": "Bitcoin",
                "blocks": [
                    {
                        "id": "81c302b7e5804e23ba50d11b83f6c44a74249870297d26113c97aa94c0307198",
                        "timestamp": 1532816367
                    },
                ]
            }
        ]
    }
    ```
3. Get a specific blockchain by id
    - URL: /api/chain/show/{id:[0-9]+}
        - Example: /api/chain/show/1
    - Accept method: GET
    - Response body:
    ```
    {
        "id": 1,
        "name": "Bitcoin",
        "blocks": [
            {
                "id": "81c302b7e5804e23ba50d11b83f6c44a74249870297d26113c97aa94c0307198",
                "timestamp": 1532816367
            },
        ]
    }
    ```
    
4. Add a block to a blockchain
    - URL: /api/chain/add-block/{id:[0-9]+}
    - Accept method: POST
    - Response body:
    ```
    {
        "id": 1,
        "name": "Bitcoin",
        "blocks": [
            {
                "id": "81c302b7e5804e23ba50d11b83f6c44a74249870297d26113c97aa94c0307198",
                "timestamp": 1532816367
            },
            {
                "id": "37ec7af7fec56b6d5addaa63f6816b472504c4595569f25206e6ef2126b0289c",
                "timestamp": 1532816401
            },
        ]
    }
    ```




