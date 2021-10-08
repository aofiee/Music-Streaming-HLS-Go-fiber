ffmpeg -y -i sender.mp3 -c:a aac -b:a 128k -map 0:a -muxdelay 0 -f segment -sc_threshold 0 -segment_time 7 -segment_list "playlist.m3u8" -segment_format mpegts "file%d.m4a"
