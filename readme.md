
# Jack's Spicy MS

A dead simple monitoring service, intended to replace things like Nagios.

## Build

Build in the normal way on Mac/Linux:

~~~
go build -o jacksspicyms cmd/web/*.go
~~~

Or on Windows:

~~~
go build -o jacksspicyms.exe cmd/web/.
~~~

Or for a particular platform:

~~~
env GOOS=linux GOARCH=amd64 go build -o jacksspicyms cmd/web/*.go
~~~

## Requirements

Jack's Spicy MS requires:
- Postgres 11 or later (db is set up as a repository, so other databases are possible)
- An account with [Pusher](https://pusher.com/), or a Pusher alternative 
(like [ipê](https://github.com/dimiro1/ipe))

## Run

First, make sure ipê is running (if you're using ipê):

On Mac/Linux
~~~
cd ipe
./ipe 
~~~

On Windows
~~~
cd ipe
ipe.exe
~~~

Run with flags:

~~~
./jacksspicyms \
-dbuser='[kiminonawa]' \
-pusherHost='localhost' \
-pusherPort='4001' \
-pusherKey='abc123' \
-pusherSecret='123abc' \
-pusherApp="1" \
-pusherSecure=false
~~~~

## All Flags

~~~~
Usage of ./jacksspicyms:
  -db string
        database name (default "jacksspicyms")
  -dbhost string
        database host (default "localhost")
  -dbport string
        database port (default "5432")
  -dbssl string
        database ssl setting (default "disable")
  -dbuser string
        database user
  -domain string
        domain name (e.g. example.com) (default "localhost")
  -identifier string
        unique identifier (default "jacksspicyms")
  -port string
        port to listen on (default ":4000")
  -production
        application is in production
  -pusherApp string
        pusher app id (default "9")
  -pusherHost string
        pusher host
  -pusherKey string
        pusher key
  -pusherPort string
        pusher port (default "443")
  -pusherSecret string
        pusher secret
   -pusherSecure
        pusher server uses SSL (true or false)
~~~~

## For the MOD peeps

Adding tables using soda:
- `soda generate fizz AddSomethingToSomethingTable` creates your /migrations/...up.fizz file where you soda to your sql
- `soda migrate` to migrate your ...up.fizz file to your (already running) database