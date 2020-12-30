# email-messages #

Endpoints:

* POST /api/message
* POST /api/send
* GET /api/messages/{emailValue}

Example requests provided in request.http file.

Required ENV's:

    PORT=8080
    DB_CLUSTER1=YourClusterAddress
    DB_USERNAME=cassandra
    DB_PASSWORD=cassandra
    DB_KEYSPACE=em
    EMAIL_HOST=smtp.gmail.com
    EMAIL_PORT=587
    EMAIL_LOGIN=YourEmailLogin
    EMAIL_PASSWORD=YourEmailPassword

Author: Klaudia Dulemba
