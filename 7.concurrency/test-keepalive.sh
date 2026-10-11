#!/usr/bin/env bash
set -euo pipefail

host=${1:-127.0.0.1}
port=${2:-7777}

request() {
  printf 'GET /hello HTTP/1.1\r\nHost: %s:%s\r\nConnection: keep-alive\r\n\r\n' "$host" "$port"
}

printf '1. 첫 요청 → TCP new connection\n' >&2
{
  request
  sleep 2
  printf '\n2. 2초 뒤 같은 연결로 요청 → TCP reused\n' >&2
  request
  printf '\n3. 6초 대기 → 서버가 5초 뒤 TCP idle timeout으로 연결 종료\n' >&2
  sleep 6
} | nc "$host" "$port"

printf '\n4. 새 nc 연결로 요청 → TCP new connection\n' >&2
request | nc "$host" "$port"
printf '\n완료. 서버 로그에서 연결 주소와 재사용 여부를 확인하세요.\n' >&2
