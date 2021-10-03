# how to use
# sh init_video.sh video_cherry_blossom.mp4 1920x1080 5

# arguments
# $1 : src file path
# $2 : dst directory path
# $3 : resolution e.g 1920x1080

# make dst directory
if [ ! -e $2/hls ]; then
    mkdir -p $2/hls
fi
if [ ! -e $2/thumbnail ]; then
    mkdir -p $2/thumbnail
fi

# process video
ffmpeg -i $1 -profile:v baseline -level 3.0 -s $3 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls $2/hls/video.m3u8
ffmpeg -i $2/hls/video.m3u8 -s $3 -vframes 1 -f image2 $2/thumbnail/thumbnail.jpg
