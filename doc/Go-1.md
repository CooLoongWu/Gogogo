# Go基础

## 1、安装【Windows下】
  建议下载.msi文件，安装时默认安装在C:\Go\文件夹下（go1.10版本）。
  安装完毕后会自动在系统环境变量的Path中写入C:\Go\bin，并帮我们建立了一个GOROOT的变量，值为C:\Go\。
  
  然后我们打开黑窗口（cmd，命令行）输入：```go version```，出现相应的go的版本则说明安装配置成功。
  ![image](img/1_version.png)
## 2、GOROOT，GOPATH
  我开发工具使用的是GoLand，开发前还需要进行下简单配置，打开GoLand在File->Settings->Go中可以看到我们需要进行GOROOT和GOPATH的配置：  
  
  GOROOT：就是go的安装路径了，C:\Go，如下图：
  ![image](img/1_goroot.png)  
  
  GOPATH：这个暂时我们设置为我们go工程的目录，比如我们的go工程为D:\Works_Go\Gogogo，那么如下图：
  ![image](img/1_gopath.png)  