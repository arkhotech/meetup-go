# meetup-go
Código ejemplo demo meetup

La carpeta goworkspace estaba en C:\
y se tenía configurada la siguiente variable de entorno:
GOPATH=C:\goworkspace\

Para medir el tiempo de ejecución  desde cygwin,
se ejecutaron los siguientes comandos:

time ./noconcurrency.exe

time GOMAXPROCS=1 ./concurrency.exe

time GOMAXPROCS=4 ./concurrency.exe
