# YakShop

### Quick Start

To run project, you can use make command `make local`, the server will be located at http://localhost:8080. You can also change the input data in `data.json` file

### Get Stock

`GET /yak-shop/stock/T`

This endpoint will return a view of your stock after T day

##### Response

```
{
    "milk" : 1104.48,
    "skins" : 3
}
```

### Get Herd

`GET /yak-shop/herd/T`

This endpoint will return a view of your herd after T day

##### Response

```
{
    "herd": [
        {
            "name": "Betty-1",
            "age": 0.63,
            "sex": "f"
        },
        {
            "name": "Betty-2",
            "age": 4.53,
            "sex": "f"
        },
        {
            "name": "Betty-3",
            "age": 10.01,
            "sex": "f"
        }
    ]
}
```
