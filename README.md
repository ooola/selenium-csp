# Selenium Content-Security-Policy Test

This app tests whether BrowserStack (Selenium) supports the
Content-Security-Policy (CSP) header. Previous webdrivers had disabled CSP.

**tl;dr** It does support CSP in the latest Chrome driver

**Todo** check other webdriver stacks.

## Instructions

* Run `go run server.go` on an internet accessible ec2 instance
* pip install selenium somewhere
* Change command_executor to your BrowserStack info in `test.py`
* Change driver.get to your ec2 instance in `test.py`
* Run `./test.py`

### Example output

When the policy is in **report only** mode, the header value is
Content-Security-Policy-Report-Only, the second button click executes and a
warning is printed in the console log.

```
$ ./test.py        
initial title: CSP Selenium Test Page
title after 1st click: Set by script with nonce
title after 2nd click: Set by script without nonce
[{u'timestamp': 1484863777616, u'message':
u"http://ec2-XYZ.compute.amazonaws.com:8080/ - The Content Security Policy
'script-src http: https: 'self' 'unsafe-inline' 'strict-dynamic'
'nonce-0yNME4F15OSPevU+UdZXGg=='' was delivered in report-only mode, but does
not specify a 'report-uri'; the policy will have no effect. Please either add
a 'report-uri' directive, or deliver the policy via the
'Content-Security-Policy' header.", u'level': u'SEVERE'}, {u'timestamp':
1484863777626, u'message': u'http://ec2-XYZ.compute.amazonaws.com:8080/ 25
[Report Only] Refused to execute inline script because it violates the
following Content Security Policy directive: "script-src http: https: \'self\'
\'unsafe-inline\' \'strict-dynamic\' \'nonce-0yNME4F15OSPevU+UdZXGg==\'". Note
that \'unsafe-inline\' is ignored if either a hash or nonce value is present
in the source list.\n', u'level': u'SEVERE'}]
```

When the policy is enforced (change line 52 to line 53 in server.go) the second
button click handler never executes, and the title never changes after the
execution of the first script to change it.

```
$ ./test.py        
initial title: CSP Selenium Test Page
title after 1st click: Set by script with nonce
title after 2nd click: Set by script with nonce
[{u'timestamp': 1484863160231, u'message':
u'http://ec2-XYZ.compute.amazonaws.com:8080/ 25 Refused to execute inline
script because it violates the following Content Security Policy directive:
"script-src http: https: \'self\' \'unsafe-inline\' \'strict-dynamic\'
\'nonce-0yNME4F15OSPevU+UdZXGg==\'". Note that \'unsafe-inline\' is ignored if
either a hash or nonce value is present in the source list.\n', u'level':
u'SEVERE'}, {u'timestamp': 1484863160232, u'message':
u'http://ec2-XYZ.compute.amazonaws.com:8080/ 36:64 Uncaught ReferenceError:
changeTitle2 is not defined', u'level': u'SEVERE'}]
```
