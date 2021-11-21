# Scalability
Notes on the Harvard CS75 Lecture 9 - https://www.youtube.com/watch?v=-W9F__D3oY4&ab_channel=JorgeScott

What features should you expect, minimally, from a web hosting platform?

## VPS x Shared Webhost

## Vertical Scaling

When you're running out of resources(CPU, memory, disk, etc.), just get more of it. You're limited to the physical machine.

## Horizontal Scaling
Instead of using one more expensive machine, why not use a bunch of cheaper machines to share the load?
But then how to handle a request if we have multiple servers? Using a load balancer. The servers can have private IP addresses and just the load balancer needs to have a public IP, it can decide which server will receive the request based on load, on the type of request(html files, images, etc.) but it can also do it in the server order, once a request hits the load balancer it returns the IP from server 1, then from server 2 and so on.(Round robin)

## Caching
Once you visit a website, your response will be cached on the way back either by your browser, operating system, or CDN. So when you click a link inside this website your request does not need to go all the way to the server going through the DNS and the load balancer.

## Sessions
Sessions are stored in the servers. So if you have multiple servers to distribute the load you will end up losing the session created in a first request when you hit the servers again because you will probably not hit the same server.
To solve this you should store sessions in a specific server to share the session state with all the other servers. But then you lose redundancy if this server dies. You can use RAID in the session server to make sure it will keeps the sessions if something goes wrong. Anyway, this is not ideal because you only have one server, if it dies you will have a HUGE downtime if something not related to its disks happens to it. This is where shared storage and replication comes in.
How to achieve Sticky sessions(even if you visist a website multiple times you're gonna visist the exact same backend server)? We could store a random Id number in a cookie and send it to the user, once it hits the load balancer it will be able to distingish to which server that user's sessions is saved.

## Load Balancers
Examples
- Software
  - ELB (aws)
  - HAProxy
  - LVS
- Hardware(Overpriced)
  - Barracuda
  - Cisco
  - Citrix