# parse-jwt-token
library to parse a JWT token.  
Do note that current method only reads token without validating it.

### steps to install
1. install from source package  
```bash
$ go get github.com/prembhaskal/parse-jwt-token
```
2. install from binary  
coming soon...


### how to use
The library reads from standard input and put the parsed jwt content on standard output.  
1. read from stdin.
```bash
$ echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" | parse-jwt-token
{
  "iat": 1516239022,
  "name": "John Doe",
  "sub": "1234567890"
}

```

2. read from variable.
```bash
$ export rawtoken="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
$ echo $rawtoken | parse-jwt-token
```

3. read from file
```bash
$ echo "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c" > /tmp/token.txt

$ cat /tmp/token.txt | parse-jwt-token

OR

$ parse-jwt-token < /tmp/token.txt

```
