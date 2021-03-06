Use frogger to dump all responses from a specific set of hosts.

## Example
1. Start frogger:

    frogger.exe -port 8082 -dump assets-cdn.github.com -dump www.google-analytics.com -dump *.googleusercontent.com

2. Configure your browser to use frogger as a proxy.

3. Navigate to any web page. All requests sent to a host that you want to dump is stored in the "dumps"-folder.

Note that HTTPS is not supported at the moment.

## CLI arguments

The following arguments can be passed to frogger:

* port Sets which port frogger listens on.
* nocache Controls wheter we allowed cached requests (if set to true, all "If-Modified-Since" headers are removed).
* dumpall If set to true, all responses from all hosts will be dumped.
* dumpheaders If set to true, all dumpfiles will contain the response headers at the bottom of the file.
* dump (multiple allowed) Sets which host to dump. If dumpall is true, this has no effect.