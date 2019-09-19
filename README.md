is-even
=======

A microservice to check whether a number is even.

Dependencies
------------

You will need to run [is-odd](https://github.com/q3k/is-odd).

Running
-------

    $ go get github.com/q3k/is-even
    $ go generate github.com/q3k/is-even/...
    $ go build github.com/q3k/is-even

    $ ./is-even -odd 127.0.0.1:2137

    $ grpcurl -plaintext -format text -d 'number: 2138' 127.0.0.1:2138 iseven.IsEven.IsEven
    result: RESULT_EVEN
    $ grpcurl -plaintext -format text -d 'number: 2137' 127.0.0.1:2138 iseven.IsEven.IsEven
    result: RESULT_NON_EVEN

License
-------

Copyright © 2019 Serge Bazanski <serge@bazanski.pl>
This work is free. You can redistribute it and/or modify it under the
terms of the Do What The Fuck You Want To Public License, Version 2,
as published by Sam Hocevar. See the LICENSE file for more details.
