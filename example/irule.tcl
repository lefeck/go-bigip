when HTTP_REQUEST {
    set req_uri [string tolower [HTTP::uri]]
    if { [string match "/example*" $req_uri] } {
        HTTP::redirect http://example.org[HTTP::uri]
    } else {
        HTTP::redirect http://example.com[HTTP::uri]
    }
}