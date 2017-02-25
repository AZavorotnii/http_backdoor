# http_backdoor
HTTP interface to execite shell commands remotely (very unsecure tool).

#Examples

Read text file:

  curl http://localhost:9090/cat?arg=/home/azavorotnii/text.txt

Write text file:

  curl -d 'This is example test data' http://localhost:9090/tee?arg=/home/azavorotnii/text.txt
