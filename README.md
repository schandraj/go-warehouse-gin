# Go Warehouse Gin

This repository is RESTful API writing using Go and Gin Framework. Also implemented Clean Code Architecture.

## How to Run this Project

### 1. Clone Project
```
$ git clone https://github.com/schandraj/go-warehouse-gin.git
$ cd go-warehouse-gin 
```

### 2. Pull
```
$ git pull origin master
```

### 3. Copy and Replace .env with your credentials
```
$ cp .env.example .env
```

### 4. Running Program
```
$ go run ./cmd/main.go
```

### 4a. Running Unit Test
```
$ go test ./... -v
```

## API lists

### Index

- [Users](#users)
- [Products](#products)
- [Locations](#locations)
- [Orders](#orders)

#### Users
|                 API                 | Description               | Auth | HTTPS | Roles |
|:-----------------------------------:|---------------------------|:----:| :---: |:-----:|
| [Register](localhost:5689/register) | Register an user          |  No  |  Yes  |  No   |
|    [Login](localhost:5689/login)    | User Login                |  No  |  Yes  |  No   |
|  [Get Me](localhost:5689/users/me)  | Get User Detail           | JWT  |  Yes  |  No   |
|  [All User](localhost:5689/users)   | Get All Users in database | JWT  |  Yes  | Admin |

#### Products
|                      API                      | Description               | Auth | HTTPS | Roles |
|:---------------------------------------------:|---------------------------|:----:| :---: |:-----:|
|  [Add New Product](localhost:5689/products)   | Add new Product           | JWT  |  Yes  | Admin |
|      [Get All](localhost:5689/products)       | Get All Products          | JWT  |  Yes  |  No   |
|   [Get By ID](localhost:5689/products/:id)    | Get Product Detail By ID  | JWT  |  Yes  |  No   |
|  [Put Product](localhost:5689/products/:id)   | Update Product By ID      | JWT  |  Yes  | Admin |
| [Delete Product](localhost:5689/products/:id) | Delete Product By ID      | JWT  |  Yes  | Admin |

#### Locations
|                     API                      | Description                | Auth | HTTPS | Roles |
|:--------------------------------------------:|----------------------------|:----:| :---: |:-----:|
|  [Add Locations](localhost:5689/locations)   | Add new Warehouse Location | JWT  |  Yes  | Admin |
| [Get All Location](localhost:5689/locations) | Get All Warehouse Location | JWT  |  Yes  | Admin |

#### Orders
|                     API                      | Description                               | Auth | HTTPS | Roles |
|:--------------------------------------------:|-------------------------------------------|:----:| :---: |:-----:|
| [Add New Order](localhost:5689/orders/:type) | Add new Order with type (receive or ship) | JWT  |  Yes  | Staff |
|       [Get All](localhost:5689/orders)       | Get All Orders                            | JWT  |  Yes  |  No   |
|    [Get By ID](localhost:5689/orders/:id)    | Get Order Detail By ID                    | JWT  |  Yes  |  No   |
