#!/bin/bash
DOCKERVER=$(<../VERSION)

docker save earthcube/chronon:$DOCKERVER | bzip2 | pv |  ssh -i /home/fils/.ssh/id_rsa root@geodex.org 'bunzip2 | docker load'

