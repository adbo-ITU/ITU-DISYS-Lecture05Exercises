# Distributed Systems - Lecture 5 exercises

## Example output

Tested with 3 machines on a local network (1 localhost, 2 other machines):

```text
2021/10/12 09:10:55 Connecting to server 10.26.18.24:8080...
2021/10/12 09:10:56 Connecting to server 127.0.0.1:8080...
2021/10/12 09:10:56 Connecting to server 10.26.31.80:8080...
2021/10/12 09:10:57 (10.26.18.24:8080) Requesting Now...
2021/10/12 09:10:57 (10.26.31.80:8080) Requesting Now...
2021/10/12 09:10:57 (127.0.0.1:8080) Requesting Now...
2021/10/12 09:10:57 (127.0.0.1:8080) Received Now response:
  - Address: 127.0.0.1:8080
  - Start time: 2021-10-12 09:10:57.6660027 +0000 UTC m=+1.738120101
  - Round trip: 7.2957ms
  - Returned time: 2021-10-12 09:10:57.6713636 +0000 UTC
  - Time delta: 5.3609ms
  - Synchronized time (Cristian): 2021-10-12 09:10:57.67501145 +0000 UTC
2021/10/12 09:10:57 (10.26.31.80:8080) Received Now response:
  - Address: 10.26.31.80:8080
  - Start time: 2021-10-12 09:10:57.6656582 +0000 UTC m=+1.737775601
  - Round trip: 13.4643ms
  - Returned time: 2021-10-12 09:10:58.3259044 +0000 UTC
  - Time delta: 660.2462ms
  - Synchronized time (Cristian): 2021-10-12 09:10:58.33263655 +0000 UTC
2021/10/12 09:10:57 (10.26.18.24:8080) Received Now response:
  - Address: 10.26.18.24:8080
  - Start time: 2021-10-12 09:10:57.6656873 +0000 UTC m=+1.737804801
  - Round trip: 22.5289ms
  - Returned time: 2021-10-12 09:10:57.989517111 +0000 UTC
  - Time delta: 323.829811ms
  - Synchronized time (Cristian): 2021-10-12 09:10:58.000781561 +0000 UTC
```
