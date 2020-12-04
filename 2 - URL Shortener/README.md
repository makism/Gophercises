
# v2
* Pass the mapping YAML file from the command line using flags.

# v1
* for `http.HandlerFunc`, I used lots of DuckDuckGo-ing and the docs https://golang.org/pkg/net/http/.
* The members in the struct must start with uppercase to be exported.
* The ref-variable passed to Unmarshall, must an array! (d`oh!).
