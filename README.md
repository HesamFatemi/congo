### CONGO

an attempt to learn and create a basic container in Go.

this repo is based on this example:
https://www.infoq.com/articles/build-a-container-golang/

#### main concepts: 
```
    1- Namespaces: 
        namespaces in linux is the ability to manage the resources and processes, such that a process only sees a specific resources.
        in unix time sharing system namespace is the hostname.

    2- chroot
        changes the current root directory for the running processes.
        
    3- cgroups
         limits and isolates the resource usage of a collection of processes.

```
#### Usage:
```
go run main.go run /bin/bash (may need sudo permission)

```
    