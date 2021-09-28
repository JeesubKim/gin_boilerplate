#start mongodb docker
> $ docker-compose up -d


#install robomongo
https://robomongo.org/

1. download file and extract using below commands
`sudo tar -xvzf robo3t-1.2.1-linux-x86_64-3e50a65.tar.gz -C /opt`
`cd /opt`
`sudo mv robo3t-1.2.1-linux-x86_64-3e50a65/ robo3t/``

2. symbolic link
`sudo ln -s /opt/robo3t/bin/robo3t /usr/local/bin/robo3t`


3. generate `~/.local/share/applications/robo3t.desktop` file and put below
```
[Desktop Entry]
Encoding=UTF-8
Name=Robo 3T
Exec=robo3t
Icon=/opt/robo3t/robo3t-logo-256x256.png
Terminal=false
Type=Application
Categories=Development;
```


4. install below
`$ sudo desktop-file-install robo3t.desktop`
