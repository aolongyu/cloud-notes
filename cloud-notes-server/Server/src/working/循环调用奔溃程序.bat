@echo off 

rem 定义需监控程序的进程名和程序路径，可根据需要进行修改

set AppName=server.exe

set AppPath=

title 进程监控

cls

echo.

echo server.exe Starting

echo.

rem 定义循环体

:startjc

   rem 从进程列表中查找指定进程

   rem  下面语句也可写成 qprocess %AppName% >nul （经验发布后补充）

   qprocess|findstr /i %AppName% >nul

   rem 变量errorlevel的值等于0表示查找到进程，否则没有查找到进程

   if %errorlevel%==0 (

         echo ^>%date:~0,10% %time:~0,8% server.exe running……

    )else (

           echo ^>%date:~0,10% %time:~0,8% don't find the server.exe

           echo ^>%date:~0,10% %time:~0,8% reset server.exe

           start %AppPath%%AppName% 2>nul && echo ^>%date:~0,10% %time:~0,8% reset success

   )

   rem 用ping命令来实现延时运行

   for /l %%i in (1,1,10) do ping -n 1 -w 1000 168.20.0.1>nul

   goto startjc

echo on