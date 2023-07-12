# physics-problem-sharing-backend

## Overview
高専の物理問題集の解答をクラス内で共有できるサービスのバックエンド

クリーンアーキテクチャ勉強用でもあるから美しいコードは書けていないかも

## Requirement

### OS

- Mac OS Ventura 13.0(動作確認済み)

### Library

- Go
  - Gin
- Docker
- docker-compose

## Installation(local)

1. Clone this repository

```
git clone https://github.com/GoRuGoo/physics-problem-sharing-backend.git
```

2. Change directory

```
cd physics-problem-sharing-backend
```

3. Build docker image

```
docker-compose up 
```


## Usage(local)

1. Build & start container

```
docker-compose up
```

2.

```
docker-compose exec api go run main.go
```

## Author

- [Yuta Ito](https://github.com/GoRuGoo)
