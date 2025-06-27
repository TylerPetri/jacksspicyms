#!/bin/zsh

# This is the bare minimum to run in development. For full list of flags,
# run ./jacksspicyms -help

go build -o jacksspicyms cmd/web/*.go && ./jacksspicyms \
-dbuser='tylerpetri' \
-pusherHost='localhost' \
-pusherKey='abc123' \
-pusherSecret='123abc' \
-pusherApp="1" \
-pusherPort="4001" \
-pusherSecure=false