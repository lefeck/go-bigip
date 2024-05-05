when HTTP_REQUEST {
    set req_uri [string tolower [HTTP::uri]]
    if { [string match "/bigip*" $req_uri] } {
        HTTP::redirect http://bigip.org[HTTP::uri]
    } else {
        HTTP::redirect http://bigip.com[HTTP::uri]
    }
}