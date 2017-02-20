#!/bin/bash

add-apt-repository -y ppa:qbittorrent-team/qbittorrent-stable
apt-get update

apt-get install -qy qbittorrent
