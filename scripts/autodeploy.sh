#!/bin/bash

curl -s https://api.github.com/repos/mshogin/RandomTrader/releases/latest | grep "browser_download_url.*deb" | cut -d : -f 2,3 | tr -d \" | wget -qi -
ls -l randomtrader_* | tail -1 | awk '{print $9}' | xargs dpkg -i
