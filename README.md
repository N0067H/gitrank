# GBSW GitRank
경북소프트웨어마이스터고등학교 학생 개발단의 깃허브 활동 순위를 보여줍니다.

## Usage
이 과정을 위해 멤버 읽기 권한을 가진 조직 액세스 토큰이 필요합니다.\
https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens\
\
\
먼저 이 저장소를 클론하고 프로젝트 루트로 이동합니다.
```bash
git clong https://github.com/gbswhs/gbsw-gitrank
cd gbsw-gitrank
```
\
그다음, 환경 변수로 액세스 토큰과 조직명을 설정합니다.
```bash
export GITHUB_TOKEN={{GITHUB_TOKEN}}
export ORGANIZATION_NAME={{ORGANIZATION_NAME}}
```
\
마지막으로, API 서버를 실행합니다.
```bash
go run api/main.go
```
