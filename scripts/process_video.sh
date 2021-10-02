# how to use
# sh init_video.sh video_cherry_blossom.mp4 1920x1080 5

# arguments
# $1 : tmp video dir path
# $2 : resolution e.g 1920x1080
# $3 : video id

# make video directory with id
if [ ! -e assets/media/$3 ]; then
    mkdir assets/media/$3
fi
if [ ! -e assets/media/$3/hls ]; then
    mkdir assets/media/$3/hls
fi
if [ ! -e assets/media/$3/thumbnail ]; then
    mkdir assets/media/$3/thumbnail
fi

# make video
ffmpeg -i $1/video.mp4 -profile:v baseline -level 3.0 -s $2 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls assets/media/$3/hls/video.m3u8
ffmpeg -i assets/media/$3/hls/video.m3u8 -s $2 -vframes 1 -f image2 assets/media/$3/thumbnail/thumbnail.jpg
