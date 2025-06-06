# GBSW GitRank
경북소프트웨어마이스터고등학교 학생 개발단의 깃허브 활동 순위를 보여줍니다.

## Usage
먼저 이 저장소를 클론합니다.
```bash
git clong https://github.com/gbswhs/gbsw-gitrank
```
\
다음으로, 멤버 읽기 권한을 가진 액세스 토큰을 발급받아야 합니다.
https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens\
\
마지막으로 ldflags -X 옵션에 액세스 토큰을 넣고 실행합니다.
```bash
go run -ldflags="-X main.apiKey=<GITHUB_ACCESS_TOKEN>" cmd/main.go
```
