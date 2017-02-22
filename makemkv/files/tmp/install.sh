#!/bin/bash

apt-get update -qq
apt-get install -qy build-essential pkg-config libc6-dev libssl-dev libexpat1-dev libavcodec-dev libgl1-mesa-dev libqt4-dev wget less

mkdir -p /tmp/sources
wget -O /tmp/sources/makemkv-bin-1.10.4.tar.gz http://www.makemkv.com/download/makemkv-bin-1.10.4.tar.gz
wget -O /tmp/sources/makemkv-oss-1.10.4.tar.gz http://www.makemkv.com/download/makemkv-oss-1.10.4.tar.gz
wget -O /tmp/sources/ffmpeg-2.8.tar.bz2 https://ffmpeg.org/releases/ffmpeg-2.8.tar.bz2
cd /tmp/sources/
tar xvzf /tmp/sources/makemkv-bin-1.10.4.tar.gz
tar xvzf /tmp/sources/makemkv-oss-1.10.4.tar.gz
tar xvjf /tmp/sources/ffmpeg-2.8.tar.bz2

#FFmpeg
cd /tmp/sources/ffmpeg-2.8
./configure --prefix=/tmp/ffmpeg --enable-static --disable-shared --enable-pic --disable-yasm
make install -j4

#Makemkv-oss
cd /tmp/sources/makemkv-oss-1.10.4
PKG_CONFIG_PATH=/tmp/ffmpeg/lib/pkgconfig ./configure
make
make install

#Makemkv-bin
cd /tmp/sources/makemkv-bin-1.10.4
/bin/echo -e 'yes' | make install

sed -i.bak '/[default_rdp_layouts]/ a rdp_layout_no=0x00000414' /etc/xrdp/xrdp_keyboard.ini
sed -i.bak '/[default_layouts_map]/ a rdp_layout_no=no' /etc/xrdp/xrdp_keyboard.ini

# cp /tmp/install/keymaps/*.ini /etc/xrdp/
# cp /tmp/install/startapp.sh /startapp.sh
# chmod +x /startapp.sh
# cp /tmp/install/firstrun.sh /etc/my_init.d/firstrun.sh
# chmod +x /etc/my_init.d/firstrun.sh

apt-get remove -qy build-essential pkg-config libc6-dev libssl-dev libexpat1-dev libavcodec-dev libgl1-mesa-dev libqt4-dev
apt-get autoremove -y
apt-get clean -y
rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/* /var/cache/*
