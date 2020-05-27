[English](https://github.com/xingliuhua/easyserver/blob/master/README.cn.md)
# easyserver

A light but powerful simulation server desktop program.


## Background
As a front-end developer or app developer, we often encounter that the server is not developed and deployed, but we want to see the effect of network request in advance. To simulate the data returned by the server, we can use the packet capturing tool

For example: Fiddler, Charles, etc., the general process is to intercept the request and the response returned by the server.

There are many advantages of the bag grabbing tool on the market:
* Our computers are basically installed.
* Easy to use, less bugs.

But we also found some shortcomings:
* It takes up a lot of memory to turn on.
* Some tools need to change the response information every time they request it, which is troublesome.
* Some tools can set the data to be returned automatically, but the data is saved by themselves. After a long time, they forget where to put it, which is easy to lose.
* Tools are all for request paths. If you want to simulate the data returned by different parameters, you need to constantly change them.
* Although the same group can also work with a bag grabbing tool, it's more troublesome to think of other people's computer changes.


## Feature
* It can set the return information for the request and automatically return the information.
* The return information is automatically saved to the disk, and the next time automatic loading is started.
* Only when the request method + interface address + all parameters are identical can they be considered as the same request.
* It also supports GUI operation and file configuration, and is easy to use.
* Support multi person online configuration and use.
* Small and light, less memory.
* Mac window linux.

## Install
1. Download the zip package of the corresponding platform.

[mac](https://github.com/xingliuhua/easyserver/blob/master/easyserver_mac_v1.0.tar.gz)

2. Decompress it without any dependency.

## Decompress it without any dependency.

1. Command line start
Run after decompression
```tex
./easyserver
```
2. Set port and start
<img src="https://github.com/xingliuhua/easyserver/blob/master/easyserver_pic_run.png"  >

We see that a dialog box will appear. We can change the port to be monitored.

For example, we want to simulate: http://test : 8888 / login interface, we can input 8888.
Click the run button (you can see the button status change after clicking the mouse to leave)

3. Browser open localhost:8888/easyserver/index

<img src="https://github.com/xingliuhua/easyserver/blob/master/easyserver_pic_index.png">
Front end or app request http://test : 8888 / login interface, refresh the webpage, we can see the request record,
Click the modify icon on the right to enter the add page, modify the return information submission, so we set a request return information. Later, requesting the login interface again returns the desired information.

4. Click the config link on the index page to see the list of previously configured requests, which supports modification and deletion.

5. If you don't want to use it, you can click Close directly. If you want to temporarily close it, please click stop.

## optimization
In the process of use, we found that there are still many areas to be optimized, which will be optimized one by one later.
* After the app requests, the home page will not automatically refresh the request record.
* Ugly page.
* The same request needs to be configured only for parameter change, which is too cumbersome, both advantages and disadvantages.

### Maintainers

[@xingliuhua](https://github.com/xingliuhua).

### Contributing

Feel free to dive in! [Open an issue] or submit PRs.
