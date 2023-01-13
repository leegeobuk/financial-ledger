# household-ledger
Simple RESTful API for keeping household ledger.

# Prerequisite
- git
- Go 1.18+
- Makefile
- Docker

# Clone project
## HTTPs
```bash
git clone https://github.com/leegeobuk/household-ledger.git
```
## SSH
```bash
git clone git@github.com:leegeobuk/household-ledger.git
```

# Test
## Unit test
```bash
make test
```

## coverage.html
```bash
make testcover
```

# Build

## Binary
```bash
make build os=<os> arch=<arch> name=<name>
```
## Linux
```bash
make build os=linux arch=amd64 name=ledger-linux-amd64
```

## MacOS
```bash
make build os=darwin arch=amd64 name=ledger-darwin-amd64
```

## Windows
```bash
make build os=windows arch=amd64 name=ledger-windows-amd64.exe
```

After the ```make build``` command, then executable binary will be created with name you entered in ```make build```.  
Full list of supported OS and architecture: https://github.com/golang/go/blob/master/src/go/build/syslist.go

## Docker
```bash
make buildimage profile=<profile>
```

### dev
```bash
make buildimage profile=dev
```

### stg
```bash
make buildimage profile=stg
```

### prd
```bash
make buildimage profile=prd
```

# Run
## Executable
if name = ledger-linux-amd64 then run
```bash
./atm-linux-amd64
```

## Go run command
```bash
go run main.go
```

## Docker
```bash
make runimage profile=<profile>
```

### dev
```bash
make runimage profile=dev
```

### stg
```bash
make runimage profile=stg
```

### prd
```bash
make runimage profile=prd
```

# 구현
우선 프로덕션 레벨에서 협업한다는 가정하에 프로젝트 환경에 맞춰 브랜치를 dev, stg, main으로 나누었습니다.  
그 다음엔 프로젝트 전체 틀을 잡기 위해서 go.mod 추가부터 도커 이미지 빌드 등 편의 기능을 담은 Makefile을 작성했습니다.  
여러 사람이 함께 협업을 하는 경우 또는 급하게 개발하는 경우 기본적인 코드 습관들을 놓치는 경우가 종종 있어서  
이런 실수들을 잡기 위해 GitHub Actions로 ```push``` 또는 ```pull request``` 생성시에 lint 검사를 하도록 추가했습니다.  

기능은
1. 가계부 목록 조회
2. 가계부 단건 조회
3. 가계부 추가
4. 회원가입
5. 로그인
6. 접근 제한 처리  
까지 구현했습니다.  

DML 파일들은 db/migrations에 위치해 있습니다.  
가계부 관련 기능은 최소한의 기능만 구현하여서 아쉬운 면이 있습니다.  
회원가입의 경우는 비밀번호를 해싱한 후에 db에 저장하도록 했습니다.
로그인 기능은 로그인 시 access_token과 refresh_token을 반환합니다.  
그리고 접근 제한 처리를 할 때 access_token 유무와 만료여부를 확인하여 처리했습니다.
아쉽게도 시간 관계상, 만료가 된 access_token을 갱신하기 위해 유효한 refresh_token을 사용하여 새로운 access_token을 재발급 받는 기능을 구현하지 못해서 아쉽습니다.  

만약 더 고도화를 한다고 하면 접근 권한 세분화, OAuth 연동 기능을 추가하고 싶습니다.  
아쉬움이 많이 남지만 또 저의 부족한 부분을 확인하는 의미 있는 시간이었습니다.  
