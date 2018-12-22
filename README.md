# DNS-server
To implement a DNS server to understand its working
<br>
<br>
First I will implement what I understand currently of how a DNS server works: 
* The records are simple -> Just the host to IP mapping
* The records are stores in a MySQL database
* The query is sent to the root DNS server, which parses the IP being received and forwards it to subsequent servers, finally reaching the authoritative servers.
* The authoritative server talks back to the root via the path taken and the result of the query is taken by the client.
<br>
<br>
Testing: <br>
Following are the way I will test my DNS server<br>
* Single query.
* Multiple clients multiple query
* Large number of clients concurrently
